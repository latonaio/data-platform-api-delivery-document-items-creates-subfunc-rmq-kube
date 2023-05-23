package subfunction

import (
	"context"
	api_input_reader "data-platform-api-delivery-document-items-creates-subfunc/API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-delivery-document-items-creates-subfunc/API_Output_Formatter"
	api_processing_data_formatter "data-platform-api-delivery-document-items-creates-subfunc/API_Processing_Data_Formatter"
	"strings"

	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	database "github.com/latonaio/golang-mysql-network-connector"
	"golang.org/x/xerrors"
)

type SubFunction struct {
	ctx context.Context
	db  *database.Mysql
	l   *logger.Logger
}

func NewSubFunction(ctx context.Context, db *database.Mysql, l *logger.Logger) *SubFunction {
	return &SubFunction{
		ctx: ctx,
		db:  db,
		l:   l,
	}
}

func (f *SubFunction) MetaData(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) *api_processing_data_formatter.MetaData {
	metaData := psdc.ConvertToMetaData(sdc)

	return metaData
}

func (f *SubFunction) ProcessType(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) *api_processing_data_formatter.ProcessType {
	processType := psdc.ConvertToProcessType()

	processType.BulkProcess = true
	// processType.IndividualProcess = true

	return processType
}

func (f *SubFunction) OrderIDInBulkProcess(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]api_processing_data_formatter.OrderID, error) {
	var data *[]api_processing_data_formatter.OrderID
	var err error

	// data, err = f.OrderIDByArraySpec(sdc, psdc)
	// if err != nil {
	// 	return nil, err
	// }

	data, err = f.OrderIDByRangeSpec(sdc, psdc)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (f *SubFunction) OrderIDByArraySpec(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]api_processing_data_formatter.OrderID, error) {
	var args []interface{}

	issuingPlantBusinessPartner := sdc.DeliveryDocumentInputParameters.IssuingPlantBusinessPartner
	receivingPlantBusinessPartner := sdc.DeliveryDocumentInputParameters.ReceivingPlantBusinessPartner
	issuingPlant := sdc.DeliveryDocumentInputParameters.IssuingPlant
	receivingPlant := sdc.DeliveryDocumentInputParameters.ReceivingPlant

	dataKey := psdc.ConvertToOrderIDKey()

	for i := range *issuingPlantBusinessPartner {
		dataKey.IssuingPlantBusinessPartner = append(dataKey.IssuingPlantBusinessPartner, (*issuingPlantBusinessPartner)[i])
	}
	for i := range *receivingPlantBusinessPartner {
		dataKey.ReceivingPlantBusinessPartner = append(dataKey.ReceivingPlantBusinessPartner, (*receivingPlantBusinessPartner)[i])
	}
	for i := range *issuingPlant {
		dataKey.IssuingPlant = append(dataKey.IssuingPlant, (*issuingPlant)[i])
	}
	for i := range *receivingPlant {
		dataKey.ReceivingPlant = append(dataKey.ReceivingPlant, (*receivingPlant)[i])
	}

	repeat1 := strings.Repeat("?,", len(dataKey.IssuingPlantBusinessPartner)-1) + "?"
	for _, v := range dataKey.IssuingPlantBusinessPartner {
		args = append(args, v)
	}
	repeat2 := strings.Repeat("?,", len(dataKey.ReceivingPlantBusinessPartner)-1) + "?"
	for _, v := range dataKey.ReceivingPlantBusinessPartner {
		args = append(args, v)
	}
	repeat3 := strings.Repeat("?,", len(dataKey.IssuingPlant)-1) + "?"
	for _, v := range dataKey.IssuingPlant {
		args = append(args, v)
	}
	repeat4 := strings.Repeat("?,", len(dataKey.ReceivingPlant)-1) + "?"
	for _, v := range dataKey.ReceivingPlant {
		args = append(args, v)
	}

	args = append(args, dataKey.IssuingPlantPartnerFunction, dataKey.ReceivingPlantPartnerFunction, dataKey.ItemCompleteDeliveryIsDefined, dataKey.ItemDeliveryBlockStatus, dataKey.ItemDeliveryStatus)

	var count *int
	err := f.db.QueryRow(
		`SELECT COUNT(*)
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_item_data
		WHERE IssuingPlantBusinessPartner IN ( `+repeat1+` )
		AND ReceivingPlantBusinessPartner IN ( `+repeat2+` )
		AND IssuingPlant IN ( `+repeat3+` )
		AND ReceivingPlant IN ( `+repeat4+` )
		AND (IssuingPlantPartnerFunction, ReceivingPlantPartnerFunction, ItemCompleteDeliveryIsDefined, ItemDeliveryBlockStatus) = (?, ?, ?, ?)
		AND ItemDeliveryStatus <> ?;`, args...,
	).Scan(&count)
	if err != nil {
		return nil, err
	}
	if *count == 0 || *count > 1000 {
		return nil, xerrors.Errorf("OrderID, OrderItemの検索結果がゼロ件または1,000件超です。")
	}

	rows, err := f.db.Query(
		`SELECT OrderID, OrderItem, IssuingPlantBusinessPartner, ReceivingPlantBusinessPartner, 
		IssuingPlant, ReceivingPlant, IssuingPlantPartnerFunction, ReceivingPlantPartnerFunction,
		ItemCompleteDeliveryIsDefined, ItemDeliveryStatus, ItemDeliveryBlockStatus
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_item_data
		WHERE IssuingPlantBusinessPartner IN ( `+repeat1+` )
		AND ReceivingPlantBusinessPartner IN ( `+repeat2+` )
		AND IssuingPlant IN ( `+repeat3+` )
		AND ReceivingPlant IN ( `+repeat4+` )
		AND (IssuingPlantPartnerFunction, ReceivingPlantPartnerFunction, ItemCompleteDeliveryIsDefined, ItemDeliveryBlockStatus) = (?, ?, ?, ?)
		AND ItemDeliveryStatus <> ?;`, args...,
	)
	if err != nil {
		return nil, err
	}

	data, err := psdc.ConvertToOrderIDByArraySpec(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) OrderIDByRangeSpec(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]api_processing_data_formatter.OrderID, error) {
	dataKey := psdc.ConvertToOrderIDKey()

	dataKey.IssuingPlantBusinessPartnerFrom = sdc.DeliveryDocumentInputParameters.IssuingPlantBusinessPartnerFrom
	dataKey.IssuingPlantBusinessPartnerTo = sdc.DeliveryDocumentInputParameters.IssuingPlantBusinessPartnerTo
	dataKey.ReceivingPlantBusinessPartnerFrom = sdc.DeliveryDocumentInputParameters.ReceivingPlantBusinessPartnerFrom
	dataKey.ReceivingPlantBusinessPartnerTo = sdc.DeliveryDocumentInputParameters.ReceivingPlantBusinessPartnerTo
	dataKey.IssuingPlantFrom = sdc.DeliveryDocumentInputParameters.IssuingPlantFrom
	dataKey.IssuingPlantTo = sdc.DeliveryDocumentInputParameters.IssuingPlantTo
	dataKey.ReceivingPlantFrom = sdc.DeliveryDocumentInputParameters.ReceivingPlantFrom
	dataKey.ReceivingPlantTo = sdc.DeliveryDocumentInputParameters.ReceivingPlantTo

	var count *int
	err := f.db.QueryRow(
		`SELECT COUNT(*)
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_item_data
		WHERE IssuingPlantBusinessPartner BETWEEN ? AND ?
		AND ReceivingPlantBusinessPartner BETWEEN ? AND ?
		AND IssuingPlant BETWEEN ? AND ?
		AND ReceivingPlant BETWEEN ? AND ?
		AND (IssuingPlantPartnerFunction, ReceivingPlantPartnerFunction, ItemCompleteDeliveryIsDefined, ItemDeliveryBlockStatus) = (?, ?, ?, ?)
		AND ItemDeliveryStatus <> ?;`, dataKey.IssuingPlantBusinessPartnerFrom, dataKey.IssuingPlantBusinessPartnerTo, dataKey.ReceivingPlantBusinessPartnerFrom, dataKey.ReceivingPlantBusinessPartnerTo, dataKey.IssuingPlantFrom, dataKey.IssuingPlantTo, dataKey.ReceivingPlantFrom, dataKey.ReceivingPlantTo, dataKey.IssuingPlantPartnerFunction, dataKey.ReceivingPlantPartnerFunction, dataKey.ItemCompleteDeliveryIsDefined, dataKey.ItemDeliveryBlockStatus, dataKey.ItemDeliveryStatus,
	).Scan(&count)
	if err != nil {
		return nil, err
	}
	if *count == 0 || *count > 1000 {
		return nil, xerrors.Errorf("OrderID, OrderItemの検索結果がゼロ件または1,000件超です。")
	}

	rows, err := f.db.Query(
		`SELECT OrderID, OrderItem, IssuingPlantBusinessPartner, ReceivingPlantBusinessPartner, 
		IssuingPlant, ReceivingPlant, IssuingPlantPartnerFunction, ReceivingPlantPartnerFunction,
		ItemCompleteDeliveryIsDefined, ItemDeliveryStatus, ItemDeliveryBlockStatus
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_item_data
		WHERE IssuingPlantBusinessPartner BETWEEN ? AND ?
		AND ReceivingPlantBusinessPartner BETWEEN ? AND ?
		AND IssuingPlant BETWEEN ? AND ?
		AND ReceivingPlant BETWEEN ? AND ?
		AND (IssuingPlantPartnerFunction, ReceivingPlantPartnerFunction, ItemCompleteDeliveryIsDefined, ItemDeliveryBlockStatus) = (?, ?, ?, ?)
		AND ItemDeliveryStatus <> ?;`, dataKey.IssuingPlantBusinessPartnerFrom, dataKey.IssuingPlantBusinessPartnerTo, dataKey.ReceivingPlantBusinessPartnerFrom, dataKey.ReceivingPlantBusinessPartnerTo, dataKey.IssuingPlantFrom, dataKey.IssuingPlantTo, dataKey.ReceivingPlantFrom, dataKey.ReceivingPlantTo, dataKey.IssuingPlantPartnerFunction, dataKey.ReceivingPlantPartnerFunction, dataKey.ItemCompleteDeliveryIsDefined, dataKey.ItemDeliveryBlockStatus, dataKey.ItemDeliveryStatus,
	)
	if err != nil {
		return nil, err
	}

	data, err := psdc.ConvertToOrderIDByRangeSpec(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) OrderIDInIndividualProcess(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]api_processing_data_formatter.OrderID, error) {
	dataKey := psdc.ConvertToOrderIDInIndividualProcessKey()

	dataKey.OrderID = sdc.DeliveryDocumentInputParameters.ReferenceDocument
	dataKey.OrderItem = sdc.DeliveryDocumentInputParameters.ReferenceDocumentItem

	rows, err := f.db.Query(
		`SELECT OrderID, OrderItem, ItemCompleteDeliveryIsDefined, ItemDeliveryStatus, ItemDeliveryBlockStatus
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_item_data
		WHERE (OrderID, OrderItem, ItemCompleteDeliveryIsDefined, ItemDeliveryBlockStatus) = (?, ?, ?, ?)
		AND ItemDeliveryStatus <> ?;`, dataKey.OrderID, dataKey.OrderItem, dataKey.ItemCompleteDeliveryIsDefined, dataKey.ItemDeliveryBlockStatus, dataKey.ItemDeliveryStatus,
	)
	if err != nil {
		return nil, err
	}

	data, err := psdc.ConvertToOrderIDInIndividualProcess(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) CreateSdc(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
	osdc *dpfm_api_output_formatter.SDC,
) error {
	var err error
	var e error

	wg := sync.WaitGroup{}
	wg.Add(2)

	psdc.MetaData = f.MetaData(sdc, psdc)
	psdc.ProcessType = f.ProcessType(sdc, psdc)

	processType := psdc.ProcessType

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		if processType.BulkProcess {
			// I-1. OrderItemの絞り込み
			psdc.OrderID, e = f.OrderIDInBulkProcess(sdc, psdc)
			if e != nil {
				err = e
				return
			}

			// I-3. オーダー明細納入日程行の絞り込みと取得
			psdc.OrderItemScheduleLine, e = f.OrderItemScheduleLine(sdc, psdc)
			if e != nil {
				err = e
				return
			}
		} else if processType.IndividualProcess {
			// II-1-1. OrderIDが未入出荷であり、かつ、OrderIDに入出荷伝票未登録残がある、明細の取得
			psdc.OrderID, e = f.OrderIDInIndividualProcess(sdc, psdc)
			if e != nil {
				err = e
				return
			}

			// II-1-3. オーダー明細納入日程行の絞り込みと取得
			psdc.OrderItemScheduleLine, e = f.OrderItemScheduleLineIndividualRegistration(sdc, psdc)
			if e != nil {
				err = e
				return
			}
		}

		// 1-5-0. DeliveryDocumentItem
		psdc.DeliveryDocumentItem = f.DeliveryDocumentItem(sdc, psdc)

		// 1-5. オーダー参照レコード・値の取得（オーダー明細）
		psdc.OrdersItem, e = f.OrdersItem(sdc, psdc)
		if e != nil {
			err = e
			return
		}
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 1-7. DeliveryDocument
		psdc.CalculateDeliveryDocument, e = f.CalculateDeliveryDocument(sdc, psdc)
		if e != nil {
			err = e
			return
		}
	}(&wg)

	wg.Wait()
	if err != nil {
		return err
	}

	f.l.Info(psdc)

	osdc, err = f.SetValue(sdc, psdc, osdc)
	if err != nil {
		return err
	}

	return nil
}
