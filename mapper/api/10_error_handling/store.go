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

type BrokenUser struct {
	ID        string `db:"id"`
	Name      string `db:"name"`
	CreatedAt int    `db:"created_at"`
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

func getOneUserWithoutLimit(db *sql.DB) (*User, error) {
	rows, err := db.Query(`
		SELECT id, name, email
		FROM users
		ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return mapper.ScanStructOne[User](rows)
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

	return mapper.ScanStructSlice[BrokenUser](rows)
}
