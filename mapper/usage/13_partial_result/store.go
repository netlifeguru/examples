package main

import (
	"database/sql"

	"github.com/netlifeguru/mapper"
)

type UserListItem struct {
	ID     string `db:"id"`
	Name   string `db:"name"`
	Active bool   `db:"active"`
}

func getUserListItems(db *sql.DB) ([]UserListItem, error) {
	rows, err := db.Query(`
		SELECT id, name, active
		FROM users
		ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []UserListItem

	err = mapper.ScanStructRows[UserListItem](rows, func(user *UserListItem) error {
		users = append(users, *user)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return users, nil
}
