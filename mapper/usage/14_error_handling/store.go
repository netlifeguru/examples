package main

import (
	"database/sql"

	"github.com/netlifeguru/mapper"
)

type BrokenUser struct {
	ID        string `db:"id"`
	Name      string `db:"name"`
	CreatedAt int    `db:"created_at"`
}

func getBrokenUsers(db *sql.DB) ([]BrokenUser, error) {
	rows, err := db.Query(`
		SELECT id, name, created_at
		FROM users
		ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []BrokenUser

	err = mapper.ScanStructRows[BrokenUser](rows, func(user *BrokenUser) error {
		users = append(users, *user)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return users, nil
}
