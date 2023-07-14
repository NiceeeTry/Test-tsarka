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
func (m UserModel) Update(id int, user *User) error {
	stmt := `UPDATE users SET name = ?, surname = ? WHERE id = ?`
	_, err := m.DB.Exec(stmt, user.Name, user.Surname, id)
	if err != nil {
		return err
	}
	return nil
}

func (m UserModel) Delete(id int) error {
	stmt := `DELETE FROM users WHERE id = ?`
	_, err := m.DB.Exec(stmt, id)
	return err
}
