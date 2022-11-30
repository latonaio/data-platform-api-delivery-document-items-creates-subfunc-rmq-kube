package requests

type ItemOrdersItemKey struct {
	OrderID                       *int   `json:"OrderID"`
	ItemCompleteDeliveryIsDefined *bool  `json:"ItemCompleteDeliveryIsDefined"`
	StockConfirmationStatus       string `json:"StockConfirmationStatus"`
}
