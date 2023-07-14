package sqlitedb

import (
	"database/sql"
)

type UserModel struct {
	DB *sql.DB
}

type User struct {
	Name    string `json:"first_name"`
	Surname string `json:"last_name"`
}

func (m UserModel) Insert(user *User) (int, error) {
	query := `INSERT INTO users (name, surname) VALUES (?, ?)`
	res, err := m.DB.Exec(query, user.Name, user.Surname)
	if err != nil {
		return -1, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}
	return int(id), nil
}

func (m UserModel) Get(id int) (*User, error) {
	user := &User{}
	stmt := `SELECT name, surname FROM users WHERE id = ?`

	row := m.DB.QueryRow(stmt, id)
	err := row.Scan(&user.Name, &user.Surname)
	if err != nil {
		return nil, err
	}
	return user, nil
}
