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

func (a actions) ListUsers() ([]User, error) {
	return []User{}, nil
}

func (a actions) DeleteUser(string) error {
	return nil
}
