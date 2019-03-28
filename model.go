// model.go

package main

import (
	"github.com/jmoiron/sqlx"
)

type user struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (u *user) getUser(db *sqlx.DB) error {
	err := db.Get(u, "SELECT * FROM users WHERE id=?", u.ID)
	return err
}

func (u *user) updateUser(db *sqlx.DB) error {
	_, err := db.Exec("UPDATE users SET name=?, age=? WHERE id=?", u.Name, u.Age, u.ID)
	return err
}

func (u *user) deleteUser(db *sqlx.DB) error {
	_, err := db.Exec("DELETE FROM users WHERE id=?", u.ID)
	return err
}

func (u *user) createUser(db *sqlx.DB) error {
	result, err := db.Exec("insert into users(name, age) values(?,?)",
		u.Name,
		u.Age)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	u.ID = int(id)
	return err
}

func getUsers(db *sqlx.DB, startid, count int) (users []user, err error) {
	err = db.Select(&users, "SELECT * FROM users WHERE id >= ? ORDER BY id LIMIT ?", startid, count)
	if err != nil {
		return users, err
	}
	return users, nil
}
