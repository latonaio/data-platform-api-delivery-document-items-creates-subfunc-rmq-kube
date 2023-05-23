package subfunction

import (
	api_input_reader "data-platform-api-delivery-document-items-creates-subfunc/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-delivery-document-items-creates-subfunc/API_Processing_Data_Formatter"

	"golang.org/x/xerrors"
)

func (f *SubFunction) CalculateDeliveryDocument(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*api_processing_data_formatter.CalculateDeliveryDocument, error) {
	metaData := psdc.MetaData
	dataKey := psdc.ConvertToCalculateDeliveryDocumentKey()

	dataKey.ServiceLabel = metaData.ServiceLabel

	rows, err := f.db.Query(
		`SELECT ServiceLabel, FieldNameWithNumberRange, LatestNumber
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_number_range_latest_number_data
		WHERE (ServiceLabel, FieldNameWithNumberRange) = (?, ?);`, dataKey.ServiceLabel, dataKey.FieldNameWithNumberRange,
	)
	if err != nil {
		return nil, err
	}

	dataQueryGets, err := psdc.ConvertToCalculateDeliveryDocumentQueryGets(rows)
	if err != nil {
		return nil, err
	}

	if dataQueryGets.DeliveryDocumentLatestNumber == nil {
		return nil, xerrors.Errorf("'data_platform_number_range_latest_number_data'テーブルのLatestNumberがNULLです。")
	}

	deliveryDocumentLatestNumber := dataQueryGets.DeliveryDocumentLatestNumber
	deliveryDocument := *dataQueryGets.DeliveryDocumentLatestNumber + 1

	data := psdc.ConvertToCalculateDeliveryDocument(deliveryDocumentLatestNumber, deliveryDocument)

	return data, err
}
