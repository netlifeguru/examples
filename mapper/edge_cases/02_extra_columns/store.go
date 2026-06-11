package main

import (
	"database/sql"

	"github.com/netlifeguru/mapper"
)

type UserSummary struct {
	ID   string `db:"id"`
	Name string `db:"name"`
}

func getUserSummaries(db *sql.DB) ([]UserSummary, error) {
	rows, err := db.Query(`
		SELECT id, name, email, active, created_at
		FROM users
		ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return mapper.ScanStructSlice[UserSummary](rows)
}
