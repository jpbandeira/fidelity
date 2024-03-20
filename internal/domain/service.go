package domain

type Service struct {
	ID          string
	Client      User
	Attendant   User
	Price       float32
	ServiceType uint8
	PaymentType uint8
	Description string
}
