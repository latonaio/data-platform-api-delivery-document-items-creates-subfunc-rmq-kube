package api_processing_data_formatter

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	BusinessPartner struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"business_partner"`
	APISchema     string   `json:"api_schema"`
	Accepter      []string `json:"accepter"`
	MaterialCode  string   `json:"material_code"`
	Plant         string   `json:"plant/supplier"`
	Stock         string   `json:"stock"`
	DocumentType  string   `json:"document_type"`
	DocumentNo    string   `json:"document_no"`
	PlannedDate   string   `json:"planned_date"`
	ValidatedDate string   `json:"validated_date"`
	Deleted       bool     `json:"deleted"`
}

type SDC struct {
	MetaData                  *MetaData                  `json:"MetaData"`
	ProcessType               *ProcessType               `json:"ProcessType"`
	OrderID                   *[]OrderID                 `json:"OrderID"`
	OrderItemScheduleLine     *[]OrderItemScheduleLine   `json:"OrderItemScheduleLine"`
	CalculateDeliveryDocument *CalculateDeliveryDocument `json:"CalculateDeliveryDocument"`
	DeliveryDocumentItem      *[]DeliveryDocumentItem    `json:"DeliveryDocumentItem"`
	OrdersItem                *[]OrdersItem              `json:"OrdersItem"`
}

// Initializer
type MetaData struct {
	BusinessPartnerID *int   `json:"business_partner"`
	ServiceLabel      string `json:"service_label"`
}

type ProcessType struct {
	BulkProcess       bool `json:"BulkProcess"`
	IndividualProcess bool `json:"IndividualProcess"`
}

type OrderIDKey struct {
	OrderID                           *int   `json:"OrderID"`
	OrderItem                         *int   `json:"OrderItem"`
	IssuingPlantPartnerFunction       string `json:"IssuingPlantPartnerFunction"`
	IssuingPlantBusinessPartner       []*int `json:"IssuingPlantBusinessPartner"`
	IssuingPlantBusinessPartnerFrom   *int   `json:"IssuingPlantBusinessPartnerFrom"`
	IssuingPlantBusinessPartnerTo     *int   `json:"IssuingPlantBusinessPartnerTo"`
	ReceivingPlantPartnerFunction     string `json:"ReceivingPlantPartnerFunction"`
	ReceivingPlantBusinessPartner     []*int `json:"ReceivingPlantBusinessPartner"`
	ReceivingPlantBusinessPartnerFrom *int   `json:"ReceivingPlantBusinessPartnerFrom"`
	ReceivingPlantBusinessPartnerTo   *int   `json:"ReceivingPlantBusinessPartnerTo"`
	IssuingPlant                      []*int `json:"IssuingPlant"`
	IssuingPlantFrom                  *int   `json:"IssuingPlantFrom"`
	IssuingPlantTo                    *int   `json:"IssuingPlantTo"`
	ReceivingPlant                    []*int `json:"ReceivingPlant"`
	ReceivingPlantFrom                int    `json:"ReceivingPlantFrom"`
	ReceivingPlantTo                  int    `json:"ReceivingPlantTo"`
	ItemCompleteDeliveryIsDefined     bool   `json:"ItemCompleteDeliveryIsDefined"`
	ItemDeliveryStatus                string `json:"ItemDeliveryStatus"`
	ItemDeliveryBlockStatus           bool   `json:"ItemDeliveryBlockStatus"`
}

type OrderID struct {
	OrderID                       int     `json:"OrderID"`
	OrderItem                     int     `json:"OrderItem"`
	IssuingPlantPartnerFunction   *string `json:"IssuingPlantPartnerFunction"`
	IssuingPlantBusinessPartner   *int    `json:"IssuingPlantBusinessPartner"`
	ReceivingPlantPartnerFunction *string `json:"ReceivingPlantPartnerFunction"`
	ReceivingPlantBusinessPartner *int    `json:"ReceivingPlantBusinessPartner"`
	IssuingPlant                  *int    `json:"IssuingPlant"`
	ReceivingPlant                *int    `json:"ReceivingPlant"`
	ItemCompleteDeliveryIsDefined *bool   `json:"ItemCompleteDeliveryIsDefined"`
	ItemDeliveryStatus            *string `json:"ItemDeliveryStatus"`
	ItemDeliveryBlockStatus       *bool   `json:"ItemDeliveryBlockStatus"`
}

type OrderItemScheduleLineKey struct {
	OrderID                             []int    `json:"OrderID"`
	OrderItem                           []int    `json:"OrderItem"`
	ItemScheduleLineDeliveryBlockStatus bool     `json:"ItemScheduleLineDeliveryBlockStatus"`
	ConfirmedDeliveryDate               []string `json:"ConfirmedDeliveryDate"`
	ConfirmedDeliveryDateFrom           string   `json:"ConfirmedDeliveryDateFrom"`
	ConfirmedDeliveryDateTo             string   `json:"ConfirmedDeliveryDateTo"`
}

type OrderItemScheduleLine struct {
	OrderID                             int      `json:"OrderID"`
	OrderItem                           int      `json:"OrderItem"`
	ItemScheduleLineDeliveryBlockStatus *bool    `json:"ItemScheduleLineDeliveryBlockStatus"`
	ConfirmedDeliveryDate               *string  `json:"ConfirmedDeliveryDate"`
	OrderItemScheduleLine               int      `json:"OrderItemScheduleLine"`
	ScheduleLineOrderQuantity           *float32 `json:"ScheduleLineOrderQuantity"`
	OpenConfdDelivQtyInOrdQtyUnit       *float32 `json:"OpenConfdDelivQtyInOrdQtyUnit"`
}

// Header
type CalculateDeliveryDocumentKey struct {
	ServiceLabel             string `json:"service_label"`
	FieldNameWithNumberRange string `json:"FieldNameWithNumberRange"`
}

type CalculateDeliveryDocumentQueryGets struct {
	ServiceLabel                 string `json:"service_label"`
	FieldNameWithNumberRange     string `json:"FieldNameWithNumberRange"`
	DeliveryDocumentLatestNumber *int   `json:"DeliveryDocumentLatestNumber"`
}

type CalculateDeliveryDocument struct {
	DeliveryDocumentLatestNumber *int `json:"DeliveryDocumentLatestNumber"`
	DeliveryDocument             int  `json:"DeliveryDocument"`
}

// Item
type DeliveryDocumentItem struct {
	DeliveryDocumentItemNumber int `json:"DeliveryDocumentItemNumber"`
}

type OrdersItem struct {
	OrderID                                      int      `json:"OrderID"`
	OrderItem                                    int      `json:"OrderItem"`
	OrderItemCategory                            string   `json:"OrderItemCategory"`
	OrderItemText                                string   `json:"OrderItemText"`
	OrderItemTextByBuyer                         string   `json:"OrderItemTextByBuyer"`
	OrderItemTextBySeller                        string   `json:"OrderItemTextBySeller"`
	Product                                      string   `json:"Product"`
	ProductStandardID                            string   `json:"ProductStandardID"`
	BaseUnit                                     string   `json:"BaseUnit"`
	StockConfirmationPartnerFunction             *string  `json:"StockConfirmationPartnerFunction"`
	StockConfirmationBusinessPartner             *int     `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant                       *string  `json:"StockConfirmationPlant"`
	StockConfirmationPlantBatch                  *string  `json:"StockConfirmationPlantBatch"`
	StockConfirmationPlantBatchValidityStartDate *string  `json:"StockConfirmationPlantBatchValidityStartDate"`
	StockConfirmationPlantBatchValidityEndDate   *string  `json:"StockConfirmationPlantBatchValidityEndDate"`
	OrderQuantityInBaseUnit                      float32  `json:"OrderQuantityInBaseUnit"`
	OrderQuantityInIssuingUnit                   float32  `json:"OrderQuantityInIssuingUnit"`
	OrderQuantityInReceivingUnit                 float32  `json:"OrderQuantityInReceivingUnit"`
	OrderIssuingUnit                             string   `json:"OrderIssuingUnit"`
	OrderReceivingUnit                           string   `json:"OrderReceivingUnit"`
	StockConfirmationPolicy                      *string  `json:"StockConfirmationPolicy"`
	StockConfirmationStatus                      *string  `json:"StockConfirmationStatus"`
	ItemWeightUnit                               *string  `json:"ItemWeightUnit"`
	ItemGrossWeight                              *float32 `json:"ItemGrossWeight"`
	ItemNetWeight                                *float32 `json:"ItemNetWeight"`
	ProductionPlantPartnerFunction               *string  `json:"ProductionPlantPartnerFunction"`
	ProductionPlantBusinessPartner               *int     `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                              *string  `json:"ProductionPlant"`
	ProductionPlantTimeZone                      *string  `json:"ProductionPlantTimeZone"`
	ProductionPlantStorageLocation               *string  `json:"ProductionPlantStorageLocation"`
	IssuingPlantPartnerFunction                  *string  `json:"IssuingPlantPartnerFunction"`
	IssuingPlantBusinessPartner                  *int     `json:"IssuingPlantBusinessPartner"`
	IssuingPlant                                 *string  `json:"IssuingPlant"`
	IssuingPlantTimeZone                         *string  `json:"IssuingPlantTimeZone"`
	IssuingPlantStorageLocation                  *string  `json:"IssuingPlantStorageLocation"`
	ReceivingPlantPartnerFunction                *string  `json:"ReceivingPlantPartnerFunction"`
	ReceivingPlantBusinessPartner                *int     `json:"ReceivingPlantBusinessPartner"`
	ReceivingPlant                               *string  `json:"ReceivingPlant"`
	ReceivingPlantTimeZone                       *string  `json:"ReceivingPlantTimeZone"`
	ReceivingPlantStorageLocation                *string  `json:"ReceivingPlantStorageLocation"`
	ProductIsBatchManagedInProductionPlant       *bool    `json:"ProductIsBatchManagedInProductionPlant"`
	ProductIsBatchManagedInIssuingPlant          *bool    `json:"ProductIsBatchManagedInIssuingPlant"`
	ProductIsBatchManagedInReceivingPlant        *bool    `json:"ProductIsBatchManagedInReceivingPlant"`
	BatchMgmtPolicyInProductionPlant             *string  `json:"BatchMgmtPolicyInProductionPlant"`
	BatchMgmtPolicyInIssuingPlant                *string  `json:"BatchMgmtPolicyInIssuingPlant"`
	BatchMgmtPolicyInReceivingPlant              *string  `json:"BatchMgmtPolicyInReceivingPlant"`
	ProductionPlantBatch                         *string  `json:"ProductionPlantBatch"`
	IssuingPlantBatch                            *string  `json:"IssuingPlantBatch"`
	ReceivingPlantBatch                          *string  `json:"ReceivingPlantBatch"`
	ProductionPlantBatchValidityStartDate        *string  `json:"ProductionPlantBatchValidityStartDate"`
	ProductionPlantBatchValidityEndDate          *string  `json:"ProductionPlantBatchValidityEndDate"`
	IssuingPlantBatchValidityStartDate           *string  `json:"IssuingPlantBatchValidityStartDate"`
	IssuingPlantBatchValidityEndDate             *string  `json:"IssuingPlantBatchValidityEndDate"`
	ReceivingPlantBatchValidityStartDate         *string  `json:"ReceivingPlantBatchValidityStartDate"`
	ReceivingPlantBatchValidityEndDate           *string  `json:"ReceivingPlantBatchValidityEndDate"`
	Incoterms                                    *string  `json:"Incoterms"`
	BPTaxClassification                          string   `json:"BPTaxClassification"`
	ProductTaxClassification                     string   `json:"ProductTaxClassification"`
	BPAccountAssignmentGroup                     string   `json:"BPAccountAssignmentGroup"`
	ProductAccountAssignmentGroup                string   `json:"ProductAccountAssignmentGroup"`
	PaymentTerms                                 string   `json:"PaymentTerms"`
	PaymentMethod                                string   `json:"PaymentMethod"`
	Project                                      *string  `json:"Project"`
	TaxCode                                      *string  `json:"TaxCode"`
	TaxRate                                      *float32 `json:"TaxRate"`
	CountryOfOrigin                              *string  `json:"CountryOfOrigin"`
}
