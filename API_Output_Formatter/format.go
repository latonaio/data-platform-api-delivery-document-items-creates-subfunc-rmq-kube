package dpfm_api_output_formatter

import (
	api_input_reader "data-platform-api-delivery-document-items-creates-subfunc/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-delivery-document-items-creates-subfunc/API_Processing_Data_Formatter"
	"encoding/json"
)

func ConvertToItem(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]Item, error) {
	itemOrdersItem := psdc.ItemOrdersItem
	calculateDeliveryDocument := psdc.CalculateDeliveryDocument
	deliveryDocumentItem := psdc.DeliveryDocumentItem
	items := make([]Item, 0, len(*itemOrdersItem))

	for i, v := range *itemOrdersItem {
		item := Item{}
		inputItem := sdc.DeliveryDocument.DeliveryDocumentItem[0]
		inputData, err := json.Marshal(inputItem)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(inputData, &item)
		if err != nil {
			return nil, err
		}

		data, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(data, &item)
		if err != nil {
			return nil, err
		}

		item.DeliveryDocument = calculateDeliveryDocument.DeliveryDocumentLatestNumber
		item.DeliveryDocumentItem = (*deliveryDocumentItem)[i].DeliveryDocumentItemNumber
		items = append(items, item)
	}

	return &items, nil
}
