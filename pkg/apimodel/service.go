package apimodel

type Service struct {
	ID        string  `json:"id"`
	Client    User    `json:"client"`
	Attendant User    `json:"attendant"`
	Price     float32 `json:"price"`
	// change to string and convert to uint in the domain layer
	ServiceType string `json:"service_type"`
	// change to string and convert to uint in the domain layer
	PaymentType string `json:"payment_type"`
	Description string `json:"description"`
}
