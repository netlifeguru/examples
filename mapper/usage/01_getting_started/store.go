package main

import (
	"database/sql"
	"time"

	"github.com/netlifeguru/mapper"
)

type User struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Active    bool      `db:"active"`
	CreatedAt time.Time `db:"created_at"`
}

func getUsers(db *sql.DB) ([]User, error) {

	rows, err := db.Query(`SELECT * FROM users ORDER BY created_at DESC`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []User

	err = mapper.ScanStructRows[User](rows, func(user *User) error {
		users = append(users, *user)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return users, nil
}
