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
	rows, err := db.Query(`
		SELECT id, name, email, active, created_at
		FROM users
		ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return mapper.ScanStructSlice[User](rows)
}
