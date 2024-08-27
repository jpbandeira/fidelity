package domain

type User struct {
	ID    string
	Name  string
	Email string
	Phone string
}

func (a actions) CreateUser(user User) (User, error) {
	user, err := a.db.CreateUser(user)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (a actions) UpdateUser(user User) (User, error) {
	user, err := a.db.UpdateUser(user)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (a actions) ListUsers(params []Param) ([]User, error) {
	return a.db.ListUsers(params)
}

func (a actions) DeleteUser(id string) error {
	return a.db.DeleteUser(id)
}
