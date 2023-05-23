package requests

type OrderIDKey struct {
	OrderID                           *int   `json:"OrderID"`
	OrderItem                         *int   `json:"OrderItem"`
	IssuingPlantPartnerFunction       string `json:"IssuingPlantPartnerFunction"`
	IssuingPlantBusinessPartner       []*int `json:"IssuingPlantBusinessPartner"`
	IssuingPlantBusinessPartnerFrom   *int   `json:"IssuingPlantBusinessPartnerFrom"`
	IssuingPlantBusinessPartnerTo     *int   `json:"IssuingPlantBusinessPartnerTo"`
	ReceivingPlantPartnerFunction     string `json:"ReceivingPlantPartnerFunction"`
	ReceivingPlantBusinessPartner     []*int `json:"ReceivingPlantBusinessPartner"`
	ReceivingPlantBusinessPartnerFrom *int   `json:"ReceivingPlantBusinessPartnerFrom"`
	ReceivingPlantBusinessPartnerTo   *int   `json:"ReceivingPlantBusinessPartnerTo"`
	IssuingPlant                      []*int `json:"IssuingPlant"`
	IssuingPlantFrom                  *int   `json:"IssuingPlantFrom"`
	IssuingPlantTo                    *int   `json:"IssuingPlantTo"`
	ReceivingPlant                    []*int `json:"ReceivingPlant"`
	ReceivingPlantFrom                int    `json:"ReceivingPlantFrom"`
	ReceivingPlantTo                  int    `json:"ReceivingPlantTo"`
	ItemCompleteDeliveryIsDefined     bool   `json:"ItemCompleteDeliveryIsDefined"`
	ItemDeliveryStatus                string `json:"ItemDeliveryStatus"`
	ItemDeliveryBlockStatus           bool   `json:"ItemDeliveryBlockStatus"`
}
