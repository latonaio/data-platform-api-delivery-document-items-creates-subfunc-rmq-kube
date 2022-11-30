package subfunction

import (
	api_input_reader "data-platform-api-delivery-document-items-creates-subfunc/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-delivery-document-items-creates-subfunc/API_Processing_Data_Formatter"
	"strings"
)

func (f *SubFunction) DeliveryDocumentItem(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]api_processing_data_formatter.DeliveryDocumentItem, error) {

	data, err := psdc.ConvertToDeliveryDocumentItem(sdc)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) OrdersItem(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]api_processing_data_formatter.ItemOrdersItem, error) {
	var args []interface{}

	orderID := psdc.OrderID
	dataKey, err := psdc.ConvertToItemOrdersItemKey(sdc, orderID)
	if err != nil {
		return nil, err
	}

	repeat1 := strings.Repeat("(?, ?),", len(*dataKey)-1) + "(?, ?)"
	for _, tag := range *dataKey {
		args = append(args, tag.OrderID, tag.ItemCompleteDeliveryIsDefined)
	}

	repeat2 := strings.Repeat("?,", len(*dataKey)-1) + "?"
	for _, tag := range *dataKey {
		args = append(args, tag.StockConfirmationStatus)
	}

	rows, err := f.db.Query(
		`SELECT OrderID,ItemCompleteDeliveryIsDefined, OrderItem, OrderItemCategory, OrderItemText, Product, ProductStandardID, 
		BaseUnit, OrderQuantityInBaseUnit, StockConfirmationPartnerFunction, StockConfirmationBusinessPartner, StockConfirmationPlant,
		StockConfirmationPolicy, StockConfirmationStatus, ItemGrossWeight, ItemNetWeight, ItemWeightUnit,
		ProductGroup, ProductionPlantPartnerFunction, ProductionPlantBusinessPartner, ProductionPlant,
		ProductionPlantTimeZone, ProductionPlantStorageLocation, IssuingPlantPartnerFunction, IssuingPlantBusinessPartner,
		IssuingPlant, IssuingPlantTimeZone, IssuingPlantStorageLocation, ReceivingPlantPartnerFunction,
		ReceivingPlantBusinessPartner, ReceivingPlant, ReceivingPlantTimeZone, ReceivingPlantStorageLocation,
		ProductIsBatchManagedInProductionPlant, ProductIsBatchManagedInIssuingPlant, ProductIsBatchManagedInReceivingPlant,
		BatchMgmtPolicyInProductionPlant, BatchMgmtPolicyInIssuingPlant, BatchMgmtPolicyInReceivingPlant,
		ProductionPlantBatch, IssuingPlantBatch, ReceivingPlantBatch, ProductionPlantBatchValidityStartDate,
		ProductionPlantBatchValidityEndDate, IssuingPlantBatchValidityStartDate, IssuingPlantBatchValidityEndDate,
		ReceivingPlantBatchValidityStartDate, ReceivingPlantBatchValidityEndDate, Incoterms, BPTaxClassification,
		ProductTaxClassification, BPAccountAssignmentGroup, ProductAccountAssignmentGroup, PaymentTerms, PaymentMethod,
		Project, TaxCode, TaxRate, CountryOfOrigin
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_item_data
		WHERE (OrderID, ItemCompleteDeliveryIsDefined) IN ( `+repeat1+` )
		AND StockConfirmationStatus NOT IN ( `+repeat2+` );`, args...,
	)
	if err != nil {
		return nil, err
	}

	data, err := psdc.ConvertToItemOrdersItem(sdc, rows)
	if err != nil {
		return nil, err
	}

	return data, err
}
