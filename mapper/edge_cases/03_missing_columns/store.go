package main

import (
	"database/sql"

	"github.com/netlifeguru/mapper"
)

type User struct {
	ID     string `db:"id"`
	Name   string `db:"name"`
	Email  string `db:"email"`
	Active bool   `db:"active"`
}

func getPartialUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query(`
		SELECT id, name
		FROM users
		ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return mapper.ScanStructSlice[User](rows)
}
