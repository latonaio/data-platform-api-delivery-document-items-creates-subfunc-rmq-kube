package requests

type OrderItemScheduleLineKey struct {
	OrderID                             []int    `json:"OrderID"`
	OrderItem                           []int    `json:"OrderItem"`
	ItemScheduleLineDeliveryBlockStatus bool     `json:"ItemScheduleLineDeliveryBlockStatus"`
	ConfirmedDeliveryDate               []string `json:"ConfirmedDeliveryDate"`
	ConfirmedDeliveryDateFrom           string   `json:"ConfirmedDeliveryDateFrom"`
	ConfirmedDeliveryDateTo             string   `json:"ConfirmedDeliveryDateTo"`
}
