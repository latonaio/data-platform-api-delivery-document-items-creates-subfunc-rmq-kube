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
	OrderID                   *[]OrderID                 `json:"OrderID`
	CalculateDeliveryDocument *CalculateDeliveryDocument `json:"CalculateDeliveryDocument`
	ItemOrdersItem            *[]ItemOrdersItem          `json"ItemOrdersItem"`
	DeliveryDocumentItem      *[]DeliveryDocumentItem    `json:"DeliveryDocumentItem"`
}

type MetaData struct {
	BusinessPartnerID *int   `json:"business_partner"`
	ServiceLabel      string `json:"service_label"`
}

type OrderIDKey struct {
	ReferenceDocument               *int   `json:"ReferenceDocument"`
	HeaderCompleteDeliveryIsDefined *bool  `json:"HeaderCompleteDeliveryIsDefined"`
	OverallDeliveryStatus           string `json:"OverallDeliveryStatus"`
}

type OrderID struct {
	ReferenceDocument               *int   `json:"ReferenceDocument"`
	OrderID                         *int   `json:"OrderID`
	HeaderCompleteDeliveryIsDefined *bool  `json:"HeaderCompleteDeliveryIsDefined"`
	OverallDeliveryStatus           string `json:"OverallDeliveryStatus"`
}

type CalculateDeliveryDocumentKey struct {
	ServiceLabel             string `json:"service_label"`
	FieldNameWithNumberRange string
}

type CalculateDeliveryDocumentQueryGets struct {
	ServiceLabel                 string `json:"service_label"`
	FieldNameWithNumberRange     string
	DeliveryDocumentLatestNumber *int
}

type CalculateDeliveryDocument struct {
	DeliveryDocumentLatestNumber *int
	DeliveryDocument             *int `json:"DeliveryDocument"`
}

type DeliveryDocumentItem struct {
	DeliveryDocumentItemNumber int `json:"DeliveryDocumentItemNumber"`
}

type ItemOrdersItemKey struct {
	OrderID                       *int   `json:"OrderID"`
	ItemCompleteDeliveryIsDefined *bool  `json:"ItemCompleteDeliveryIsDefined"`
	StockConfirmationStatus       string `json:"StockConfirmationStatus"`
}

type ItemOrdersItem struct {
	DeliveryDocument                       *int     `json:"DeliveryDocument`
	OrderID                                *int     `json:"OrderID"`
	ItemCompleteDeliveryIsDefined          bool     `json:"ItemCompleteDeliveryIsDefined"`
	OrderItem                              *int     `json:"OrderItem"`
	OrderItemCategory                      string   `json:"OrderItemCategory"`
	OrderItemText                          string   `json:"OrderItemText"`
	Product                                string   `json:"Product"`
	ProductStandardID                      string   `json:"ProductStandardID"`
	BaseUnit                               string   `json:"BaseUnit`
	OrderQuantityInBaseUnit                *string  `json:"OrderQuantityInBaseUnit`
	StockConfirmationPartnerFunction       *string  `json:"StockConfirmationPartnerFunction"`
	StockConfirmationBusinessPartner       *int     `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant                 *string  `json:"StockConfirmationPlant"`
	StockConfirmationPolicy                *string  `json:"StockConfirmationPolicy"`
	StockConfirmationStatus                *string  `json:"StockConfirmationStatus"`
	ItemGrossWeight                        *float32 `json:"ItemGrossWeight`
	ItemNetWeight                          *float32 `json:"ItemNetWeight"`
	ItemWeightUnit                         *string  `json:"ItemWeightUnit"`
	ProductGroup                           *string  `json:"ProductGroup"`
	ProductionPlantPartnerFunction         *string  `json:"ProductionPlantPartnerFunction"`
	ProductionPlantBusinessPartner         *string  `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                        *string  `json:"ProductionPlant"`
	ProductionPlantTimeZone                *string  `json:"ProductionPlantTimeZone"`
	ProductionPlantStorageLocation         *string  `json:"ProductionPlantStorageLocation`
	IssuingPlantPartnerFunction            *string  `json:"IssuingPlantPartnerFunction"`
	IssuingPlantBusinessPartner            *string  `json:"IssuingPlantBusinessPartner"`
	IssuingPlant                           *string  `json:"IssuingPlant"`
	IssuingPlantTimeZone                   *string  `json:"IssuingPlantTimeZone"`
	IssuingPlantStorageLocation            *string  `json:"IssuingPlantStorageLocation"`
	ReceivingPlantPartnerFunction          *string  `json:"ReceivingPlantPartnerFunction"`
	ReceivingPlantBusinessPartner          *string  `json:"ReceivingPlantBusinessPartner"`
	ReceivingPlant                         *string  `json:"ReceivingPlant"`
	ReceivingPlantTimeZone                 *string  `json:"ReceivingPlantTimeZone"`
	ReceivingPlantStorageLocation          *string  `json:"ReceivingPlantStorageLocation`
	ProductIsBatchManagedInProductionPlant *bool    `json:"ProductIsBatchManagedInProductionPlant"`
	ProductIsBatchManagedInIssuingPlant    *bool    `json:"ProductIsBatchManagedInIssuingPlant"`
	ProductIsBatchManagedInReceivingPlant  *bool    `json:"ProductIsBatchManagedInReceivingPlant"`
	BatchMgmtPolicyInProductionPlant       *string  `json:"BatchMgmtPolicyInProductionPlant"`
	BatchMgmtPolicyInIssuingPlant          *string  `json:"BatchMgmtPolicyInIssuingPlant"`
	BatchMgmtPolicyInReceivingPlant        *string  `json:"BatchMgmtPolicyInReceivingPlant`
	ProductionPlantBatch                   *string  `json:"ProductionPlantBatch"`
	IssuingPlantBatch                      *string  `json:"IssuingPlantBatch"`
	ReceivingPlantBatch                    *string  `json:"ReceivingPlantBatch"`
	ProductionPlantBatchValidityStartDate  *string  `json:"ProductionPlantBatchValidityStartDate"`
	ProductionPlantBatchValidityEndDate    *string  `json:"ProductionPlantBatchValidityEndDate"`
	IssuingPlantBatchValidityStartDate     *string  `json:"IssuingPlantBatchValidityStartDate"`
	IssuingPlantBatchValidityEndDate       *string  `json:"IssuingPlantBatchValidityEndDate"`
	ReceivingPlantBatchValidityStartDate   *string  `json:"ReceivingPlantBatchValidityStartDate"`
	ReceivingPlantBatchValidityEndDate     *string  `json:"ReceivingPlantBatchValidityEndDate"`
	Incoterms                              *string  `json:"Incoterms`
	BPTaxClassification                    string   `json:"BPTaxClassification"`
	ProductTaxClassification               string   `json:"ProductTaxClassification"`
	BPAccountAssignmentGroup               string   `json:"BPAccountAssignmentGroup"`
	ProductAccountAssignmentGroup          string   `json:"ProductAccountAssignmentGroup"`
	PaymentTerms                           string   `json:"PaymentTerms"`
	PaymentMethod                          string   `json:"PaymentMethod`
	Project                                *string  `json:"Project"`
	TaxCode                                *string  `json:"TaxCode"`
	TaxRate                                *float32 `json:"TaxRate"`
	CountryOfOrigin                        *string  `json:"CountryOfOrigin"`
}
