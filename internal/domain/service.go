package domain

type Service struct {
	ID          string
	Client      User
	Attendant   User
	Price       float32
	ServiceType string
	PaymentType string
	Description string
}

func (a *actions) CreateService(service Service) (Service, error) {
	service, err := a.db.CreateService(service)
	if err != nil {
		return Service{}, err
	}

	return service, nil
}
