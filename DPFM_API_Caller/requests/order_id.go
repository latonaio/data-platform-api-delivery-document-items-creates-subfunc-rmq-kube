package requests

type OrderID struct {
	ReferenceDocument               *int   `json:"ReferenceDocument"`
	OrderID                         *int   `json:"OrderID`
	HeaderCompleteDeliveryIsDefined *bool  `json:"HeaderCompleteDeliveryIsDefined"`
	OverallDeliveryStatus           string `json:"OverallDeliveryStatus"`
}
