package main

import (
	"database/sql"
	"time"

	"github.com/netlifeguru/mapper"
)

type User struct {
	ID        string
	Name      string
	Email     string
	Active    bool
	CreatedAt time.Time
}

func getUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query(`
		SELECT id, name, email, active, created_at
		FROM users
		ORDER BY created_at DESC
	`)
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
