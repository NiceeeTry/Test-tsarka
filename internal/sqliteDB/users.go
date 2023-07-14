package sqlitedb

import (
	"database/sql"
	"fmt"
)

type UserModel struct {
	DB *sql.DB
}

type User struct {
	Name    string `json:"first_name"`
	Surname string `json:"last_name"`
}

func (m UserModel) Insert(user *User) error {
	query := `INSERT INTO users (name, surname) VALUES (?, ?)`
	id, err := m.DB.Exec(query, user.Name, user.Surname)
	if err != nil {
		return err
	}
	fmt.Println(id)
	return nil
}
