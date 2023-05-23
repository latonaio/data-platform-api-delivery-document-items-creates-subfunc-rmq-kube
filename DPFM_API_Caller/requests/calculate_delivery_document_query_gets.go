package requests

type CalculateDeliveryDocumentQueryGets struct {
	ServiceLabel                 string `json:"service_label"`
	FieldNameWithNumberRange     string `json:"FieldNameWithNumberRange"`
	DeliveryDocumentLatestNumber *int   `json:"DeliveryDocumentLatestNumber"`
}
