package main

import (
	"database/sql"

	"github.com/netlifeguru/mapper"
)

type User struct {
	ID        string `db:"id"`
	Name      string `db:"name"`
	Active    bool   `db:"active"`
	CreatedAt string `db:"created_at"`
}

func getUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query(`
		SELECT
			id,
			name,
			active,
			created_at
		FROM users
		ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return mapper.ScanStructSlice[User](rows)
}
