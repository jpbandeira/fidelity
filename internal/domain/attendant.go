package domain

import "time"

type Attendant struct {
	ID        string
	Name      string
	Email     string
	Phone     string
	CreatedAt time.Time
}

func (a actions) CreateAttendant(attendant Attendant) (Attendant, error) {
	attendant, err := a.db.CreateAttendant(attendant)
	if err != nil {
		return Attendant{}, err
	}

	return attendant, nil
}

func (a actions) UpdateAttendant(attendant Attendant) (Attendant, error) {
	attendant, err := a.db.UpdateAttendant(attendant)
	if err != nil {
		return Attendant{}, err
	}

	return attendant, nil
}

func (a actions) ListAttendants(params []Param) ([]Attendant, error) {
	return a.db.ListAttendants(params)
}

func (a actions) DeleteAttendant(id string) error {
	return a.db.DeleteAttendant(id)
}
