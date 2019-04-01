package main

type user struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Age  *int   `db:"age" json:"age,omitempty"`
}
