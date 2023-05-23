package requests

type OrderID struct {
	OrderID                       int     `json:"OrderID"`
	OrderItem                     int     `json:"OrderItem"`
	IssuingPlantPartnerFunction   *string `json:"IssuingPlantPartnerFunction"`
	IssuingPlantBusinessPartner   *int    `json:"IssuingPlantBusinessPartner"`
	ReceivingPlantPartnerFunction *string `json:"ReceivingPlantPartnerFunction"`
	ReceivingPlantBusinessPartner *int    `json:"ReceivingPlantBusinessPartner"`
	IssuingPlant                  *int    `json:"IssuingPlant"`
	ReceivingPlant                *int    `json:"ReceivingPlant"`
	ItemCompleteDeliveryIsDefined *bool   `json:"ItemCompleteDeliveryIsDefined"`
	ItemDeliveryStatus            *string `json:"ItemDeliveryStatus"`
	ItemDeliveryBlockStatus       *bool   `json:"ItemDeliveryBlockStatus"`
}
