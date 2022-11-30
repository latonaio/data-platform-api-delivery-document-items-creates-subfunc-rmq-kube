package requests

type CalculateDeliveryDocumentKey struct {
	ServiceLabel             string `json:"service_label"`
	FieldNameWithNumberRange string
}
