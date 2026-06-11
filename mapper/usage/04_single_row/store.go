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

func getUserByID(db *sql.DB, id int) (*User, error) {
	rows, err := db.Query(`
		SELECT id, name, email, active, created_at
		FROM users
		WHERE id = ?
		LIMIT 1
	`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result *User

	err = mapper.ScanStructRows[User](rows, func(user *User) error {
		u := *user
		result = &u
		return nil
	})
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, sql.ErrNoRows
	}

	return result, nil
}
