package api_processing_data_formatter

import (
	api_input_reader "data-platform-api-delivery-document-items-creates-subfunc/API_Input_Reader"
	"data-platform-api-delivery-document-items-creates-subfunc/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func getBoolPtr(b bool) *bool {
	return &b
}

// initializer
func (psdc *SDC) ConvertToMetaData(sdc *api_input_reader.SDC) (*MetaData, error) {
	pm := &requests.MetaData{
		BusinessPartnerID: sdc.BusinessPartnerID,
		ServiceLabel:      sdc.ServiceLabel,
	}
	data := pm

	metaData := MetaData{
		BusinessPartnerID: data.BusinessPartnerID,
		ServiceLabel:      data.ServiceLabel,
	}

	return &metaData, nil
}

func (psdc *SDC) ConvertToOrderIDKey(sdc *api_input_reader.SDC) (*OrderIDKey, error) {
	pm := &requests.OrderIDKey{
		ReferenceDocument:               sdc.DeliveryDocument.ReferenceDocument,
		HeaderCompleteDeliveryIsDefined: getBoolPtr(false),
		OverallDeliveryStatus:           "CL",
	}
	data := pm

	orderIDKey := OrderIDKey{
		ReferenceDocument:               data.ReferenceDocument,
		HeaderCompleteDeliveryIsDefined: data.HeaderCompleteDeliveryIsDefined,
		OverallDeliveryStatus:           data.OverallDeliveryStatus,
	}

	return &orderIDKey, nil
}

func (psdc *SDC) ConvertToOrderID(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*[]OrderID, error) {
	var orderID []OrderID
	pm := &requests.OrderID{
		ReferenceDocument:               nil,
		OrderID:                         nil,
		HeaderCompleteDeliveryIsDefined: nil,
		OverallDeliveryStatus:           "",
	}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.OrderID,
			&pm.HeaderCompleteDeliveryIsDefined,
			&pm.OverallDeliveryStatus,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		pm.ReferenceDocument = pm.OrderID

		data := pm
		orderID = append(orderID, OrderID{
			ReferenceDocument:               data.ReferenceDocument,
			OrderID:                         data.OrderID,
			HeaderCompleteDeliveryIsDefined: data.HeaderCompleteDeliveryIsDefined,
			OverallDeliveryStatus:           data.OverallDeliveryStatus,
		})
	}

	return &orderID, nil
}

// DeliveryDocument
func (psdc *SDC) ConvertToDeliveryDocumentItem(sdc *api_input_reader.SDC) (*[]DeliveryDocumentItem, error) {
	var deliveryDocumentItem []DeliveryDocumentItem

	for i := range sdc.DeliveryDocument.DeliveryDocumentItem {
		pm := &requests.DeliveryDocumentItem{}

		pm.DeliveryDocumentItemNumber = i + 1

		data := pm
		deliveryDocumentItem = append(deliveryDocumentItem, DeliveryDocumentItem{
			DeliveryDocumentItemNumber: data.DeliveryDocumentItemNumber,
		})
	}

	return &deliveryDocumentItem, nil
}

func (psdc *SDC) ConvertToCalculateDeliveryDocumentKey() (*CalculateDeliveryDocumentKey, error) {
	pm := &requests.CalculateDeliveryDocumentKey{
		ServiceLabel:             "",
		FieldNameWithNumberRange: "DeliveryDocument",
	}
	data := pm

	calculateDeliveryDocumentKey := CalculateDeliveryDocumentKey{
		ServiceLabel:             data.ServiceLabel,
		FieldNameWithNumberRange: data.FieldNameWithNumberRange,
	}

	return &calculateDeliveryDocumentKey, nil
}

func (psdc *SDC) ConvertToCalculateDeliveryDocumentQueryGets(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*CalculateDeliveryDocumentQueryGets, error) {
	pm := &requests.CalculateDeliveryDocumentQueryGets{
		ServiceLabel:                 "",
		FieldNameWithNumberRange:     "",
		DeliveryDocumentLatestNumber: nil,
	}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.ServiceLabel,
			&pm.FieldNameWithNumberRange,
			&pm.DeliveryDocumentLatestNumber,
		)
		if err != nil {
			return nil, err
		}
	}
	data := pm

	calculateDeliveryDocumentQueryGets := CalculateDeliveryDocumentQueryGets{
		ServiceLabel:                 data.ServiceLabel,
		FieldNameWithNumberRange:     data.FieldNameWithNumberRange,
		DeliveryDocumentLatestNumber: data.DeliveryDocumentLatestNumber,
	}

	return &calculateDeliveryDocumentQueryGets, nil
}

func (psdc *SDC) ConvertToCalculateDeliveryDocument(
	deliveryDocumentLatestNumber *int,
) (*CalculateDeliveryDocument, error) {
	pm := &requests.CalculateDeliveryDocument{
		DeliveryDocumentLatestNumber: nil,
		DeliveryDocument:             nil,
	}

	pm.DeliveryDocumentLatestNumber = deliveryDocumentLatestNumber
	data := pm

	calculateDeliveryDocument := CalculateDeliveryDocument{
		DeliveryDocumentLatestNumber: data.DeliveryDocumentLatestNumber,
		DeliveryDocument:             data.DeliveryDocument,
	}

	return &calculateDeliveryDocument, nil
}

// Item
func (psdc *SDC) ConvertToItemOrdersItemKey(sdc *api_input_reader.SDC, orderID *[]OrderID) (*[]ItemOrdersItemKey, error) {
	var itemOrdersItemKey []ItemOrdersItemKey
	var pm []requests.ItemOrdersItemKey
	for _, v := range *orderID {
		pm = append(pm, requests.ItemOrdersItemKey{
			OrderID:                       v.OrderID,
			ItemCompleteDeliveryIsDefined: getBoolPtr(false),
			StockConfirmationStatus:       "NP",
		})
	}

	data := pm

	for _, v := range data {
		itemOrdersItemKey = append(itemOrdersItemKey, ItemOrdersItemKey{
			OrderID:                       v.OrderID,
			ItemCompleteDeliveryIsDefined: v.ItemCompleteDeliveryIsDefined,
			StockConfirmationStatus:       v.StockConfirmationStatus,
		})
	}

	return &itemOrdersItemKey, nil
}

func (psdc *SDC) ConvertToItemOrdersItem(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*[]ItemOrdersItem, error) {
	var itemOrdersItem []ItemOrdersItem

	pm := &requests.ItemOrdersItem{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.OrderID,
			&pm.ItemCompleteDeliveryIsDefined,
			&pm.OrderItem,
			&pm.OrderItemCategory,
			&pm.OrderItemText,
			&pm.Product,
			&pm.ProductStandardID,
			&pm.BaseUnit,
			&pm.OrderQuantityInBaseUnit,
			&pm.StockConfirmationPartnerFunction,
			&pm.StockConfirmationBusinessPartner,
			&pm.StockConfirmationPlant,
			&pm.StockConfirmationPolicy,
			&pm.StockConfirmationStatus,
			&pm.ItemGrossWeight,
			&pm.ItemNetWeight,
			&pm.ItemWeightUnit,
			&pm.ProductGroup,
			&pm.ProductionPlantPartnerFunction,
			&pm.ProductionPlantBusinessPartner,
			&pm.ProductionPlant,
			&pm.ProductionPlantTimeZone,
			&pm.ProductionPlantStorageLocation,
			&pm.IssuingPlantPartnerFunction,
			&pm.IssuingPlantBusinessPartner,
			&pm.IssuingPlant,
			&pm.IssuingPlantTimeZone,
			&pm.IssuingPlantStorageLocation,
			&pm.ReceivingPlantPartnerFunction,
			&pm.ReceivingPlantBusinessPartner,
			&pm.ReceivingPlant,
			&pm.ReceivingPlantTimeZone,
			&pm.ReceivingPlantStorageLocation,
			&pm.ProductIsBatchManagedInProductionPlant,
			&pm.ProductIsBatchManagedInIssuingPlant,
			&pm.ProductIsBatchManagedInReceivingPlant,
			&pm.BatchMgmtPolicyInProductionPlant,
			&pm.BatchMgmtPolicyInIssuingPlant,
			&pm.BatchMgmtPolicyInReceivingPlant,
			&pm.ProductionPlantBatch,
			&pm.IssuingPlantBatch,
			&pm.ReceivingPlantBatch,
			&pm.ProductionPlantBatchValidityStartDate,
			&pm.ProductionPlantBatchValidityEndDate,
			&pm.IssuingPlantBatchValidityStartDate,
			&pm.IssuingPlantBatchValidityEndDate,
			&pm.ReceivingPlantBatchValidityStartDate,
			&pm.ReceivingPlantBatchValidityEndDate,
			&pm.Incoterms,
			&pm.BPTaxClassification,
			&pm.ProductTaxClassification,
			&pm.BPAccountAssignmentGroup,
			&pm.ProductAccountAssignmentGroup,
			&pm.PaymentTerms,
			&pm.PaymentMethod,
			&pm.Project,
			&pm.TaxCode,
			&pm.TaxRate,
			&pm.CountryOfOrigin,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		itemOrdersItem = append(itemOrdersItem, ItemOrdersItem{
			DeliveryDocument:                       data.DeliveryDocument,
			OrderID:                                data.OrderID,
			ItemCompleteDeliveryIsDefined:          data.ItemCompleteDeliveryIsDefined,
			OrderItem:                              data.OrderItem,
			OrderItemCategory:                      data.OrderItemCategory,
			OrderItemText:                          data.OrderItemText,
			Product:                                data.Product,
			ProductStandardID:                      data.ProductStandardID,
			BaseUnit:                               data.BaseUnit,
			OrderQuantityInBaseUnit:                data.OrderQuantityInBaseUnit,
			StockConfirmationPartnerFunction:       data.StockConfirmationPartnerFunction,
			StockConfirmationBusinessPartner:       data.StockConfirmationBusinessPartner,
			StockConfirmationPlant:                 data.StockConfirmationPlant,
			StockConfirmationPolicy:                data.StockConfirmationPolicy,
			StockConfirmationStatus:                data.StockConfirmationStatus,
			ItemGrossWeight:                        data.ItemGrossWeight,
			ItemNetWeight:                          data.ItemNetWeight,
			ItemWeightUnit:                         data.ItemWeightUnit,
			ProductGroup:                           data.ProductGroup,
			ProductionPlantPartnerFunction:         data.ProductionPlantPartnerFunction,
			ProductionPlantBusinessPartner:         data.ProductionPlantBusinessPartner,
			ProductionPlant:                        data.ProductionPlant,
			ProductionPlantTimeZone:                data.ProductionPlantTimeZone,
			ProductionPlantStorageLocation:         data.ProductionPlantStorageLocation,
			IssuingPlantPartnerFunction:            data.IssuingPlantPartnerFunction,
			IssuingPlantBusinessPartner:            data.IssuingPlantBusinessPartner,
			IssuingPlant:                           data.IssuingPlant,
			IssuingPlantTimeZone:                   data.IssuingPlantTimeZone,
			IssuingPlantStorageLocation:            data.IssuingPlantStorageLocation,
			ReceivingPlantPartnerFunction:          data.ReceivingPlantPartnerFunction,
			ReceivingPlantBusinessPartner:          data.ReceivingPlantBusinessPartner,
			ReceivingPlant:                         data.ReceivingPlant,
			ReceivingPlantTimeZone:                 data.ReceivingPlantTimeZone,
			ReceivingPlantStorageLocation:          data.ReceivingPlantStorageLocation,
			ProductIsBatchManagedInProductionPlant: data.ProductIsBatchManagedInProductionPlant,
			ProductIsBatchManagedInIssuingPlant:    data.ProductIsBatchManagedInIssuingPlant,
			ProductIsBatchManagedInReceivingPlant:  data.ProductIsBatchManagedInReceivingPlant,
			BatchMgmtPolicyInProductionPlant:       data.BatchMgmtPolicyInProductionPlant,
			BatchMgmtPolicyInIssuingPlant:          data.BatchMgmtPolicyInIssuingPlant,
			BatchMgmtPolicyInReceivingPlant:        data.BatchMgmtPolicyInReceivingPlant,
			ProductionPlantBatch:                   data.ProductionPlantBatch,
			IssuingPlantBatch:                      data.IssuingPlantBatch,
			ReceivingPlantBatch:                    data.ReceivingPlantBatch,
			ProductionPlantBatchValidityStartDate:  data.ProductionPlantBatchValidityStartDate,
			ProductionPlantBatchValidityEndDate:    data.ProductionPlantBatchValidityEndDate,
			IssuingPlantBatchValidityStartDate:     data.IssuingPlantBatchValidityStartDate,
			IssuingPlantBatchValidityEndDate:       data.IssuingPlantBatchValidityEndDate,
			ReceivingPlantBatchValidityStartDate:   data.ReceivingPlantBatchValidityStartDate,
			ReceivingPlantBatchValidityEndDate:     data.ReceivingPlantBatchValidityEndDate,
			Incoterms:                              data.Incoterms,
			BPTaxClassification:                    data.BPTaxClassification,
			ProductTaxClassification:               data.ProductTaxClassification,
			BPAccountAssignmentGroup:               data.BPAccountAssignmentGroup,
			ProductAccountAssignmentGroup:          data.ProductAccountAssignmentGroup,
			PaymentTerms:                           data.PaymentTerms,
			PaymentMethod:                          data.PaymentMethod,
			Project:                                data.Project,
			TaxCode:                                data.TaxCode,
			TaxRate:                                data.TaxRate,
			CountryOfOrigin:                        data.CountryOfOrigin,
		})
	}

	return &itemOrdersItem, nil
}
