package requests

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
