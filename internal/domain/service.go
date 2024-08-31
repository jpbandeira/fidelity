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

type ServiceList struct {
	Items []Service
	Total int
	Count int
}

func (a *actions) CreateService(service Service) (Service, error) {
	client, err := a.db.GetUser(service.Client.ID)
	if err != nil {
		return Service{}, err
	}

	attendant, err := a.db.GetUser(service.Attendant.ID)
	if err != nil {
		return Service{}, err
	}

	service, err = a.db.CreateService(service, attendant.ID, client.ID)
	if err != nil {
		return Service{}, err
	}

	return service, nil
}
