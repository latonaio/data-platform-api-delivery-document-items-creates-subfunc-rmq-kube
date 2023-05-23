package requests

type OrderItemScheduleLine struct {
	OrderID                             int      `json:"OrderID"`
	OrderItem                           int      `json:"OrderItem"`
	ItemScheduleLineDeliveryBlockStatus *bool    `json:"ItemScheduleLineDeliveryBlockStatus"`
	ConfirmedDeliveryDate               *string  `json:"ConfirmedDeliveryDate"`
	OrderItemScheduleLine               int      `json:"OrderItemScheduleLine"`
	ScheduleLineOrderQuantity           *float32 `json:"ScheduleLineOrderQuantity"`
	OpenConfdDelivQtyInOrdQtyUnit       *float32 `json:"OpenConfdDelivQtyInOrdQtyUnit"`
}
