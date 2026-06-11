package main

import (
	"database/sql"

	"github.com/netlifeguru/mapper"
)

type User struct {
	ID    string `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
}

func getMissingUser(db *sql.DB) (*User, error) {
	rows, err := db.Query(`
		SELECT id, name, email
		FROM users
		WHERE id = ?
		LIMIT 1
	`, "missing-user-id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return mapper.ScanStructOne[User](rows)
}
