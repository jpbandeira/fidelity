package apimodel

type Service struct {
	ID          string  `json:"ID"`
	Client      User    `json:"Client"`
	Attendant   User    `json:"Attendant"`
	Price       float32 `json:"Price"`
	ServiceType string  `json:"ServiceType"`
	PaymentType string  `json:"PaymentType"`
	Description string  `json:"Description"`
}
