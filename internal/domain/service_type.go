package domain

type ServiceType struct {
	Description string
}

type ServiceTypeList struct {
	Items []ServiceType
	Total int
	Count int
}

// func (a *actions) CreateServiceType(service ServiceType) (ServiceType, error) {
// 	return a.db.CreateServiceType(service)
// }

func (a actions) ListServiceTypes(params []Param) ([]ServiceType, error) {
	return a.db.ListServiceTypes(params)
}
