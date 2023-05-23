package subfunction

import (
	api_input_reader "data-platform-api-delivery-document-items-creates-subfunc/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-delivery-document-items-creates-subfunc/API_Processing_Data_Formatter"
	"strings"
)

func (f *SubFunction) OrderItemScheduleLine(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]api_processing_data_formatter.OrderItemScheduleLine, error) {
	var args []interface{}

	orderID := psdc.OrderID
	inputParameters := sdc.DeliveryDocumentInputParameters

	dataKey := psdc.ConvertToOrderItemScheduleLineKey()

	for i := range *orderID {
		dataKey.OrderID = append(dataKey.OrderID, (*orderID)[i].OrderID)
		dataKey.OrderItem = append(dataKey.OrderItem, (*orderID)[i].OrderItem)
	}

	dataKey.ConfirmedDeliveryDateFrom = inputParameters.ConfirmedDeliveryDateFrom
	dataKey.ConfirmedDeliveryDateTo = inputParameters.ConfirmedDeliveryDateTo

	repeat := strings.Repeat("(?,?),", len(dataKey.OrderID)-1) + "(?,?)"
	for i := range dataKey.OrderID {
		args = append(args, dataKey.OrderID[i], dataKey.OrderItem[i])
	}

	args = append(args, dataKey.ItemScheduleLineDeliveryBlockStatus, dataKey.ConfirmedDeliveryDateFrom, dataKey.ConfirmedDeliveryDateTo)

	rows, err := f.db.Query(
		`SELECT OrderID, OrderItem, ItemScheduleLineDeliveryBlockStatus, ConfirmedDeliveryDate, ScheduleLine, OrderQuantityInBaseUnit, OpenConfirmedQuantityInBaseUnit
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_item_schedule_line_data
		WHERE (OrderID, OrderItem) IN ( `+repeat+` )
		AND ItemScheduleLineDeliveryBlockStatus = ?
		AND ConfirmedDeliveryDate BETWEEN ? AND ?;`, args...,
	)
	if err != nil {
		return nil, err
	}

	// // ConfirmedDeliveryDateをInputParameterの配列を使う場合
	// for i := range *inputParameters.ConfirmedDeliveryDate {
	// 	dataKey.ConfirmedDeliveryDate = append(dataKey.ConfirmedDeliveryDate, (*inputParameters.ConfirmedDeliveryDate)[i])
	// }

	// repeat1 := strings.Repeat("(?,?),", len(dataKey.OrderID)-1) + "(?,?)"
	// for i := range dataKey.OrderID {
	// 	args = append(args, dataKey.OrderID[i], dataKey.OrderItem[i])
	// }

	// args = append(args, dataKey.ItemScheduleLineDeliveryBlockStatus)

	// repeat2 := strings.Repeat("?,", len(dataKey.ConfirmedDeliveryDate)-1) + "?"
	// for i := range dataKey.ConfirmedDeliveryDate {
	// 	args = append(args, dataKey.ConfirmedDeliveryDate[i])
	// }

	// rows, err := f.db.Query(
	// 	`SELECT OrderID, OrderItem, ItemScheduleLineDeliveryBlockStatus, ConfirmedDeliveryDate, ScheduleLine, OrderQuantityInBaseUnit, OpenConfirmedQuantityInBaseUnit
	// 	FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_item_schedule_line_data
	// 	WHERE (OrderID, OrderItem) IN ( `+repeat1+` )
	// 	AND ItemScheduleLineDeliveryBlockStatus = ?
	// 	AND ConfirmedDeliveryDate IN ( `+repeat2+` );`, args...,
	// )
	// if err != nil {
	// 	return nil, err
	// }

	data, err := psdc.ConvertToOrderItemScheduleLine(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) OrderItemScheduleLineIndividualRegistration(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]api_processing_data_formatter.OrderItemScheduleLine, error) {
	var args []interface{}

	orderID := psdc.OrderID

	dataKey := psdc.ConvertToOrderItemScheduleLineKey()

	for i := range *orderID {
		dataKey.OrderID = append(dataKey.OrderID, (*orderID)[i].OrderID)
		dataKey.OrderItem = append(dataKey.OrderItem, (*orderID)[i].OrderItem)
	}

	repeat := strings.Repeat("(?,?),", len(dataKey.OrderID)-1) + "(?,?)"
	for i := range dataKey.OrderID {
		args = append(args, dataKey.OrderID[i], dataKey.OrderItem[i])
	}

	args = append(args, dataKey.ItemScheduleLineDeliveryBlockStatus)

	rows, err := f.db.Query(
		`SELECT OrderID, OrderItem, ItemScheduleLineDeliveryBlockStatus, ConfirmedDeliveryDate, ScheduleLine, OrderQuantityInBaseUnit, OpenConfirmedQuantityInBaseUnit
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_item_schedule_line_data
		WHERE (OrderID, OrderItem) IN ( `+repeat+` )
		AND ItemScheduleLineDeliveryBlockStatus = ?;`, args...,
	)
	if err != nil {
		return nil, err
	}

	data, err := psdc.ConvertToOrderItemScheduleLine(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) DeliveryDocumentItem(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) *[]api_processing_data_formatter.DeliveryDocumentItem {
	var length int
	processType := psdc.ProcessType

	if processType.BulkProcess {
		length = len(*psdc.OrderItemScheduleLine)
	} else if processType.IndividualProcess {
		length = len(sdc.DeliveryDocument.DeliveryDocumentItem)
	}

	data := psdc.ConvertToDeliveryDocumentItem(length)

	return data
}

func (f *SubFunction) OrdersItem(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]api_processing_data_formatter.OrdersItem, error) {
	var args []interface{}

	orderID := psdc.OrderID

	repeat := strings.Repeat("(?, ?),", len(*orderID)-1) + "(?, ?)"
	for _, v := range *orderID {
		args = append(args, v.OrderID, v.OrderItem)
	}

	rows, err := f.db.Query(
		`SELECT OrderID, OrderItem, OrderItemCategory, OrderItemText, OrderItemTextByBuyer, OrderItemTextBySeller,
		Product, ProductStandardID, BaseUnit, StockConfirmationPartnerFunction, StockConfirmationBusinessPartner,
		StockConfirmationPlant, StockConfirmationPlantBatch, StockConfirmationPlantBatchValidityStartDate,
		StockConfirmationPlantBatchValidityEndDate, OrderQuantityInBaseUnit, OrderQuantityInIssuingUnit,
		OrderQuantityInReceivingUnit, OrderIssuingUnit, OrderReceivingUnit, StockConfirmationPolicy,
		StockConfirmationStatus, ItemWeightUnit, ItemGrossWeight, ItemNetWeight, ProductionPlantPartnerFunction,
		ProductionPlantBusinessPartner, ProductionPlant, ProductionPlantTimeZone, ProductionPlantStorageLocation,
		IssuingPlantPartnerFunction, IssuingPlantBusinessPartner, IssuingPlant, IssuingPlantTimeZone,
		IssuingPlantStorageLocation, ReceivingPlantPartnerFunction, ReceivingPlantBusinessPartner, ReceivingPlant,
		ReceivingPlantTimeZone, ReceivingPlantStorageLocation, ProductIsBatchManagedInProductionPlant,
		ProductIsBatchManagedInIssuingPlant, ProductIsBatchManagedInReceivingPlant, BatchMgmtPolicyInProductionPlant,
		BatchMgmtPolicyInIssuingPlant, BatchMgmtPolicyInReceivingPlant, ProductionPlantBatch, IssuingPlantBatch,
		ReceivingPlantBatch, ProductionPlantBatchValidityStartDate, ProductionPlantBatchValidityEndDate,
		IssuingPlantBatchValidityStartDate, IssuingPlantBatchValidityEndDate, ReceivingPlantBatchValidityStartDate,
		ReceivingPlantBatchValidityEndDate, Incoterms, BPTaxClassification, ProductTaxClassification,
		BPAccountAssignmentGroup, ProductAccountAssignmentGroup, PaymentTerms, PaymentMethod, Project, TaxCode, TaxRate, CountryOfOrigin
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_item_data
		WHERE (OrderID, OrderItem) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		return nil, err
	}

	data, err := psdc.ConvertToOrdersItem(sdc, rows)
	if err != nil {
		return nil, err
	}

	return data, err
}
