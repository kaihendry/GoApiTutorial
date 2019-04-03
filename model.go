// model.go

package main

import (
	"github.com/jmoiron/sqlx"
)

type user struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Age  *int   `db:"age" json:"age,omitempty"`
}

func (u *user) getUser(db *sqlx.DB) error {
	err := db.Get(u, "SELECT * FROM users WHERE id=?", u.ID)
	return err
}

func (u *user) updateUser(db *sqlx.DB) error {
	_, err := db.NamedExec(`UPDATE users SET name=:name, age=:age WHERE id = :id`, u)
	return err
}

func (u *user) deleteUser(db *sqlx.DB) error {
	_, err := db.Exec("DELETE FROM users WHERE id=?", u.ID)
	return err
}

func (u *user) createUser(db *sqlx.DB) error {
	result, err := db.NamedExec(`INSERT INTO users (name, age) VALUES (:name, :age)`, u)
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
	if len(users) == 0 {
		return []user{}, nil
	}
	return users, nil
}
