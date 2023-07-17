package mock

import (
	"errors"

	sqlitedb "Alikhan.Aitbayev/Data/sqliteDB"
)

var mockUser = &sqlitedb.User{
	Name:    "John",
	Surname: "Silver",
}

type UserModel struct{}

func (m UserModel) Insert(user *sqlitedb.User) (int, error) {
	return 2, nil
}

func (m UserModel) Get(id int) (*sqlitedb.User, error) {
	switch id {
	case 1:
		return mockUser, nil
	default:
		return nil, errors.New("error: no records")
	}
}

func (m UserModel) Update(id int, user *sqlitedb.User) error {
	return nil
}
func (m UserModel) Delete(id int) error {
	return nil
}
