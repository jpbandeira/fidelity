package apimodel

type Service struct {
	ID          string  `json:"id"`
	Client      User    `json:"client"`
	Attendant   User    `json:"attendant"`
	Price       float32 `json:"price"`
	ServiceType string  `json:"service_type"`
	PaymentType string  `json:"payment_type"`
	Description string  `json:"description"`
}
