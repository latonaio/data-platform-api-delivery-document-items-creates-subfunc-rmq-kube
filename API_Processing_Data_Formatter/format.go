package api_processing_data_formatter

import (
	api_input_reader "data-platform-api-delivery-document-items-creates-subfunc/API_Input_Reader"
	"data-platform-api-delivery-document-items-creates-subfunc/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

// Initializer
func (psdc *SDC) ConvertToMetaData(sdc *api_input_reader.SDC) *MetaData {
	pm := &requests.MetaData{
		BusinessPartnerID: sdc.BusinessPartnerID,
		ServiceLabel:      sdc.ServiceLabel,
	}
	data := pm

	metaData := MetaData{
		BusinessPartnerID: data.BusinessPartnerID,
		ServiceLabel:      data.ServiceLabel,
	}

	return &metaData
}

func (psdc *SDC) ConvertToProcessType() *ProcessType {
	pm := &requests.ProcessType{}
	data := pm

	processType := ProcessType{
		BulkProcess:       data.BulkProcess,
		IndividualProcess: data.IndividualProcess,
	}

	return &processType
}

func (psdc *SDC) ConvertToOrderIDKey() *OrderIDKey {
	pm := &requests.OrderIDKey{
		IssuingPlantPartnerFunction:   "DELIVERFROM",
		ReceivingPlantPartnerFunction: "DELIVERTO",
		ItemCompleteDeliveryIsDefined: false,
		ItemDeliveryStatus:            "CL",
		ItemDeliveryBlockStatus:       false,
	}

	data := pm
	orderIDKey := OrderIDKey{
		OrderID:                           data.OrderID,
		OrderItem:                         data.OrderItem,
		IssuingPlantPartnerFunction:       data.IssuingPlantPartnerFunction,
		IssuingPlantBusinessPartner:       data.IssuingPlantBusinessPartner,
		IssuingPlantBusinessPartnerFrom:   data.IssuingPlantBusinessPartnerFrom,
		IssuingPlantBusinessPartnerTo:     data.IssuingPlantBusinessPartnerTo,
		ReceivingPlantPartnerFunction:     data.ReceivingPlantPartnerFunction,
		ReceivingPlantBusinessPartner:     data.ReceivingPlantBusinessPartner,
		ReceivingPlantBusinessPartnerFrom: data.ReceivingPlantBusinessPartnerFrom,
		ReceivingPlantBusinessPartnerTo:   data.ReceivingPlantBusinessPartnerTo,
		IssuingPlant:                      data.IssuingPlant,
		IssuingPlantFrom:                  data.IssuingPlantFrom,
		IssuingPlantTo:                    data.IssuingPlantTo,
		ReceivingPlant:                    data.ReceivingPlant,
		ReceivingPlantFrom:                data.ReceivingPlantFrom,
		ReceivingPlantTo:                  data.ReceivingPlantTo,
		ItemCompleteDeliveryIsDefined:     data.ItemCompleteDeliveryIsDefined,
		ItemDeliveryStatus:                data.ItemDeliveryStatus,
		ItemDeliveryBlockStatus:           data.ItemDeliveryBlockStatus,
	}

	return &orderIDKey
}

func (psdc *SDC) ConvertToOrderIDByArraySpec(rows *sql.Rows) (*[]OrderID, error) {
	var orderID []OrderID

	for i := 0; true; i++ {
		pm := &requests.OrderID{}

		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_orders_item_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.OrderID,
			&pm.OrderItem,
			&pm.IssuingPlantBusinessPartner,
			&pm.ReceivingPlantBusinessPartner,
			&pm.IssuingPlant,
			&pm.ReceivingPlant,
			&pm.IssuingPlantPartnerFunction,
			&pm.ReceivingPlantPartnerFunction,
			&pm.ItemCompleteDeliveryIsDefined,
			&pm.ItemDeliveryStatus,
			&pm.ItemDeliveryBlockStatus,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		orderID = append(orderID, OrderID{
			OrderID:                       data.OrderID,
			OrderItem:                     data.OrderItem,
			IssuingPlantPartnerFunction:   data.IssuingPlantPartnerFunction,
			IssuingPlantBusinessPartner:   data.IssuingPlantBusinessPartner,
			ReceivingPlantPartnerFunction: data.ReceivingPlantPartnerFunction,
			ReceivingPlantBusinessPartner: data.ReceivingPlantBusinessPartner,
			IssuingPlant:                  data.IssuingPlant,
			ReceivingPlant:                data.ReceivingPlant,
			ItemCompleteDeliveryIsDefined: data.ItemCompleteDeliveryIsDefined,
			ItemDeliveryStatus:            data.ItemDeliveryStatus,
			ItemDeliveryBlockStatus:       data.ItemDeliveryBlockStatus,
		})
	}

	return &orderID, nil
}

func (psdc *SDC) ConvertToOrderIDByRangeSpec(rows *sql.Rows) (*[]OrderID, error) {
	var orderID []OrderID

	for i := 0; true; i++ {
		pm := &requests.OrderID{}

		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_orders_item_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.OrderID,
			&pm.OrderItem,
			&pm.IssuingPlantBusinessPartner,
			&pm.ReceivingPlantBusinessPartner,
			&pm.IssuingPlant,
			&pm.ReceivingPlant,
			&pm.IssuingPlantPartnerFunction,
			&pm.ReceivingPlantPartnerFunction,
			&pm.ItemCompleteDeliveryIsDefined,
			&pm.ItemDeliveryStatus,
			&pm.ItemDeliveryBlockStatus,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		orderID = append(orderID, OrderID{
			OrderID:                       data.OrderID,
			OrderItem:                     data.OrderItem,
			IssuingPlantPartnerFunction:   data.IssuingPlantPartnerFunction,
			IssuingPlantBusinessPartner:   data.IssuingPlantBusinessPartner,
			ReceivingPlantPartnerFunction: data.ReceivingPlantPartnerFunction,
			ReceivingPlantBusinessPartner: data.ReceivingPlantBusinessPartner,
			IssuingPlant:                  data.IssuingPlant,
			ReceivingPlant:                data.ReceivingPlant,
			ItemCompleteDeliveryIsDefined: data.ItemCompleteDeliveryIsDefined,
			ItemDeliveryStatus:            data.ItemDeliveryStatus,
			ItemDeliveryBlockStatus:       data.ItemDeliveryBlockStatus,
		})
	}

	return &orderID, nil
}

func (psdc *SDC) ConvertToOrderIDInIndividualProcessKey() *OrderIDKey {
	pm := &requests.OrderIDKey{
		ItemCompleteDeliveryIsDefined: false,
		ItemDeliveryStatus:            "CL",
		ItemDeliveryBlockStatus:       false,
	}

	data := pm
	orderIDKey := OrderIDKey{
		OrderID:                           data.OrderID,
		OrderItem:                         data.OrderItem,
		IssuingPlantPartnerFunction:       data.IssuingPlantPartnerFunction,
		IssuingPlantBusinessPartner:       data.IssuingPlantBusinessPartner,
		IssuingPlantBusinessPartnerFrom:   data.IssuingPlantBusinessPartnerFrom,
		IssuingPlantBusinessPartnerTo:     data.IssuingPlantBusinessPartnerTo,
		ReceivingPlantPartnerFunction:     data.ReceivingPlantPartnerFunction,
		ReceivingPlantBusinessPartner:     data.ReceivingPlantBusinessPartner,
		ReceivingPlantBusinessPartnerFrom: data.ReceivingPlantBusinessPartnerFrom,
		ReceivingPlantBusinessPartnerTo:   data.ReceivingPlantBusinessPartnerTo,
		IssuingPlant:                      data.IssuingPlant,
		IssuingPlantFrom:                  data.IssuingPlantFrom,
		IssuingPlantTo:                    data.IssuingPlantTo,
		ReceivingPlant:                    data.ReceivingPlant,
		ReceivingPlantFrom:                data.ReceivingPlantFrom,
		ReceivingPlantTo:                  data.ReceivingPlantTo,
		ItemCompleteDeliveryIsDefined:     data.ItemCompleteDeliveryIsDefined,
		ItemDeliveryStatus:                data.ItemDeliveryStatus,
		ItemDeliveryBlockStatus:           data.ItemDeliveryBlockStatus,
	}

	return &orderIDKey
}

func (psdc *SDC) ConvertToOrderIDInIndividualProcess(rows *sql.Rows) (*[]OrderID, error) {
	var orderID []OrderID

	for i := 0; true; i++ {
		pm := &requests.OrderID{}

		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_orders_item_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.OrderID,
			&pm.OrderItem,
			&pm.ItemCompleteDeliveryIsDefined,
			&pm.ItemDeliveryStatus,
			&pm.ItemDeliveryBlockStatus,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		orderID = append(orderID, OrderID{
			OrderID:                       data.OrderID,
			OrderItem:                     data.OrderItem,
			IssuingPlantPartnerFunction:   data.IssuingPlantPartnerFunction,
			IssuingPlantBusinessPartner:   data.IssuingPlantBusinessPartner,
			ReceivingPlantPartnerFunction: data.ReceivingPlantPartnerFunction,
			ReceivingPlantBusinessPartner: data.ReceivingPlantBusinessPartner,
			IssuingPlant:                  data.IssuingPlant,
			ReceivingPlant:                data.ReceivingPlant,
			ItemCompleteDeliveryIsDefined: data.ItemCompleteDeliveryIsDefined,
			ItemDeliveryStatus:            data.ItemDeliveryStatus,
			ItemDeliveryBlockStatus:       data.ItemDeliveryBlockStatus,
		})
	}

	return &orderID, nil
}

func (psdc *SDC) ConvertToOrderItemScheduleLineKey() *OrderItemScheduleLineKey {
	pm := &requests.OrderItemScheduleLineKey{
		ItemScheduleLineDeliveryBlockStatus: false,
	}

	data := pm
	orderItemScheduleLineKey := OrderItemScheduleLineKey{
		OrderID:                             data.OrderID,
		OrderItem:                           data.OrderItem,
		ItemScheduleLineDeliveryBlockStatus: data.ItemScheduleLineDeliveryBlockStatus,
		ConfirmedDeliveryDate:               data.ConfirmedDeliveryDate,
		ConfirmedDeliveryDateFrom:           data.ConfirmedDeliveryDateFrom,
		ConfirmedDeliveryDateTo:             data.ConfirmedDeliveryDateTo,
	}

	return &orderItemScheduleLineKey
}

func (psdc *SDC) ConvertToOrderItemScheduleLine(rows *sql.Rows) (*[]OrderItemScheduleLine, error) {
	var orderItemScheduleLine []OrderItemScheduleLine

	for i := 0; true; i++ {
		pm := &requests.OrderItemScheduleLine{}
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_orders_item_schedule_line_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.OrderID,
			&pm.OrderItem,
			&pm.ItemScheduleLineDeliveryBlockStatus,
			&pm.ConfirmedDeliveryDate,
			&pm.OrderItemScheduleLine,
			&pm.ScheduleLineOrderQuantity,
			&pm.OpenConfdDelivQtyInOrdQtyUnit,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		orderItemScheduleLine = append(orderItemScheduleLine, OrderItemScheduleLine{
			OrderID:                             data.OrderID,
			OrderItem:                           data.OrderItem,
			ItemScheduleLineDeliveryBlockStatus: data.ItemScheduleLineDeliveryBlockStatus,
			ConfirmedDeliveryDate:               data.ConfirmedDeliveryDate,
			OrderItemScheduleLine:               data.OrderItemScheduleLine,
			ScheduleLineOrderQuantity:           data.ScheduleLineOrderQuantity,
			OpenConfdDelivQtyInOrdQtyUnit:       data.OpenConfdDelivQtyInOrdQtyUnit,
		})
	}

	return &orderItemScheduleLine, nil
}

// Header
func (psdc *SDC) ConvertToCalculateDeliveryDocumentKey() *CalculateDeliveryDocumentKey {
	pm := &requests.CalculateDeliveryDocumentKey{
		FieldNameWithNumberRange: "DeliveryDocument",
	}

	data := pm
	calculateDeliveryDocumentKey := CalculateDeliveryDocumentKey{
		ServiceLabel:             data.ServiceLabel,
		FieldNameWithNumberRange: data.FieldNameWithNumberRange,
	}

	return &calculateDeliveryDocumentKey
}

func (psdc *SDC) ConvertToCalculateDeliveryDocumentQueryGets(rows *sql.Rows) (*CalculateDeliveryDocumentQueryGets, error) {
	pm := &requests.CalculateDeliveryDocumentQueryGets{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_number_range_latest_number_data'テーブルに対象のレコードが存在しません。")
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

func (psdc *SDC) ConvertToCalculateDeliveryDocument(deliveryDocumentLatestNumber *int, deliveryDocument int) *CalculateDeliveryDocument {
	pm := &requests.CalculateDeliveryDocument{}

	pm.DeliveryDocumentLatestNumber = deliveryDocumentLatestNumber
	pm.DeliveryDocument = deliveryDocument

	data := pm
	calculateDeliveryDocument := CalculateDeliveryDocument{
		DeliveryDocumentLatestNumber: data.DeliveryDocumentLatestNumber,
		DeliveryDocument:             data.DeliveryDocument,
	}

	return &calculateDeliveryDocument
}

// Item
func (psdc *SDC) ConvertToDeliveryDocumentItem(length int) *[]DeliveryDocumentItem {
	var deliveryDocumentItem []DeliveryDocumentItem

	for i := 0; i < length; i++ {
		pm := &requests.DeliveryDocumentItem{}

		pm.DeliveryDocumentItemNumber = i + 1

		data := pm
		deliveryDocumentItem = append(deliveryDocumentItem, DeliveryDocumentItem{
			DeliveryDocumentItemNumber: data.DeliveryDocumentItemNumber,
		})
	}

	return &deliveryDocumentItem
}

func (psdc *SDC) ConvertToOrdersItem(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*[]OrdersItem, error) {
	var ordersItem []OrdersItem

	for i := 0; true; i++ {
		pm := &requests.OrdersItem{}
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_orders_item_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.OrderID,
			&pm.OrderItem,
			&pm.OrderItemCategory,
			&pm.OrderItemText,
			&pm.OrderItemTextByBuyer,
			&pm.OrderItemTextBySeller,
			&pm.Product,
			&pm.ProductStandardID,
			&pm.BaseUnit,
			&pm.StockConfirmationPartnerFunction,
			&pm.StockConfirmationBusinessPartner,
			&pm.StockConfirmationPlant,
			&pm.StockConfirmationPlantBatch,
			&pm.StockConfirmationPlantBatchValidityStartDate,
			&pm.StockConfirmationPlantBatchValidityEndDate,
			&pm.OrderQuantityInBaseUnit,
			&pm.OrderQuantityInIssuingUnit,
			&pm.OrderQuantityInReceivingUnit,
			&pm.OrderIssuingUnit,
			&pm.OrderReceivingUnit,
			&pm.StockConfirmationPolicy,
			&pm.StockConfirmationStatus,
			&pm.ItemWeightUnit,
			&pm.ItemGrossWeight,
			&pm.ItemNetWeight,
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
		ordersItem = append(ordersItem, OrdersItem{
			OrderID:                          data.OrderID,
			OrderItem:                        data.OrderItem,
			OrderItemCategory:                data.OrderItemCategory,
			OrderItemText:                    data.OrderItemText,
			OrderItemTextByBuyer:             data.OrderItemTextByBuyer,
			OrderItemTextBySeller:            data.OrderItemTextBySeller,
			Product:                          data.Product,
			ProductStandardID:                data.ProductStandardID,
			BaseUnit:                         data.BaseUnit,
			StockConfirmationPartnerFunction: data.StockConfirmationPartnerFunction,
			StockConfirmationBusinessPartner: data.StockConfirmationBusinessPartner,
			StockConfirmationPlant:           data.StockConfirmationPlant,
			StockConfirmationPlantBatch:      data.StockConfirmationPlantBatch,
			StockConfirmationPlantBatchValidityStartDate: data.StockConfirmationPlantBatchValidityStartDate,
			StockConfirmationPlantBatchValidityEndDate:   data.StockConfirmationPlantBatchValidityEndDate,
			OrderQuantityInBaseUnit:                      data.OrderQuantityInBaseUnit,
			OrderQuantityInIssuingUnit:                   data.OrderQuantityInIssuingUnit,
			OrderQuantityInReceivingUnit:                 data.OrderQuantityInReceivingUnit,
			OrderIssuingUnit:                             data.OrderIssuingUnit,
			OrderReceivingUnit:                           data.OrderReceivingUnit,
			StockConfirmationPolicy:                      data.StockConfirmationPolicy,
			StockConfirmationStatus:                      data.StockConfirmationStatus,
			ItemWeightUnit:                               data.ItemWeightUnit,
			ItemGrossWeight:                              data.ItemGrossWeight,
			ItemNetWeight:                                data.ItemNetWeight,
			ProductionPlantPartnerFunction:               data.ProductionPlantPartnerFunction,
			ProductionPlantBusinessPartner:               data.ProductionPlantBusinessPartner,
			ProductionPlant:                              data.ProductionPlant,
			ProductionPlantTimeZone:                      data.ProductionPlantTimeZone,
			ProductionPlantStorageLocation:               data.ProductionPlantStorageLocation,
			IssuingPlantPartnerFunction:                  data.IssuingPlantPartnerFunction,
			IssuingPlantBusinessPartner:                  data.IssuingPlantBusinessPartner,
			IssuingPlant:                                 data.IssuingPlant,
			IssuingPlantTimeZone:                         data.IssuingPlantTimeZone,
			IssuingPlantStorageLocation:                  data.IssuingPlantStorageLocation,
			ReceivingPlantPartnerFunction:                data.ReceivingPlantPartnerFunction,
			ReceivingPlantBusinessPartner:                data.ReceivingPlantBusinessPartner,
			ReceivingPlant:                               data.ReceivingPlant,
			ReceivingPlantTimeZone:                       data.ReceivingPlantTimeZone,
			ReceivingPlantStorageLocation:                data.ReceivingPlantStorageLocation,
			ProductIsBatchManagedInProductionPlant:       data.ProductIsBatchManagedInProductionPlant,
			ProductIsBatchManagedInIssuingPlant:          data.ProductIsBatchManagedInIssuingPlant,
			ProductIsBatchManagedInReceivingPlant:        data.ProductIsBatchManagedInReceivingPlant,
			BatchMgmtPolicyInProductionPlant:             data.BatchMgmtPolicyInProductionPlant,
			BatchMgmtPolicyInIssuingPlant:                data.BatchMgmtPolicyInIssuingPlant,
			BatchMgmtPolicyInReceivingPlant:              data.BatchMgmtPolicyInReceivingPlant,
			ProductionPlantBatch:                         data.ProductionPlantBatch,
			IssuingPlantBatch:                            data.IssuingPlantBatch,
			ReceivingPlantBatch:                          data.ReceivingPlantBatch,
			ProductionPlantBatchValidityStartDate:        data.ProductionPlantBatchValidityStartDate,
			ProductionPlantBatchValidityEndDate:          data.ProductionPlantBatchValidityEndDate,
			IssuingPlantBatchValidityStartDate:           data.IssuingPlantBatchValidityStartDate,
			IssuingPlantBatchValidityEndDate:             data.IssuingPlantBatchValidityEndDate,
			ReceivingPlantBatchValidityStartDate:         data.ReceivingPlantBatchValidityStartDate,
			ReceivingPlantBatchValidityEndDate:           data.ReceivingPlantBatchValidityEndDate,
			Incoterms:                                    data.Incoterms,
			BPTaxClassification:                          data.BPTaxClassification,
			ProductTaxClassification:                     data.ProductTaxClassification,
			BPAccountAssignmentGroup:                     data.BPAccountAssignmentGroup,
			ProductAccountAssignmentGroup:                data.ProductAccountAssignmentGroup,
			PaymentTerms:                                 data.PaymentTerms,
			PaymentMethod:                                data.PaymentMethod,
			Project:                                      data.Project,
			TaxCode:                                      data.TaxCode,
			TaxRate:                                      data.TaxRate,
			CountryOfOrigin:                              data.CountryOfOrigin,
		})
	}

	return &ordersItem, nil
}
