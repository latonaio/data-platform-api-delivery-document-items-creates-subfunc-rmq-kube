package requests

type OrderIDKey struct {
	ReferenceDocument               *int   `json:"ReferenceDocument"`
	HeaderCompleteDeliveryIsDefined *bool  `json:"HeaderCompleteDeliveryIsDefined"`
	OverallDeliveryStatus           string `json:"OverallDeliveryStatus"`
}
