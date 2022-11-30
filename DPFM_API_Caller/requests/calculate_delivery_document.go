package requests

type CalculateDeliveryDocument struct {
	DeliveryDocumentLatestNumber *int
	DeliveryDocument             *int `json:"DeliveryDocument"`
}
