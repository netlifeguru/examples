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

	var users []UserSummary

	err = mapper.ScanStructRows[UserSummary](rows, func(user *UserSummary) error {
		users = append(users, *user)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return users, nil
}
