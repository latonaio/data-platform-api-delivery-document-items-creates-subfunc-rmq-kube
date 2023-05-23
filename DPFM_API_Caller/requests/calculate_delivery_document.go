package requests

type CalculateDeliveryDocument struct {
	DeliveryDocumentLatestNumber *int `json:"DeliveryDocumentLatestNumber"`
	DeliveryDocument             int  `json:"DeliveryDocument"`
}
