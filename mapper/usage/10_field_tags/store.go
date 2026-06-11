package main

import (
	"database/sql"

	"github.com/netlifeguru/mapper"
)

type UserDTO struct {
	UserID    string `db:"id"`
	FullName  string `db:"name"`
	EmailAddr string `db:"email"`
}

func getUserDTOs(db *sql.DB) ([]UserDTO, error) {
	rows, err := db.Query(`
		SELECT id, name, email
		FROM users
		ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []UserDTO

	err = mapper.ScanStructRows[UserDTO](rows, func(user *UserDTO) error {
		users = append(users, *user)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return users, nil
}
