package requests

type CalculateDeliveryDocumentQueryGets struct {
	ServiceLabel                 string `json:"service_label"`
	FieldNameWithNumberRange     string
	DeliveryDocumentLatestNumber *int
}
