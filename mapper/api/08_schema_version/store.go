package main

import (
	"database/sql"
	"time"

	"github.com/netlifeguru/mapper"
)

const usersListCacheKey = "users:list:created_at_desc"

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

	var users []User

	err = mapper.ScanStructRowsWithCacheKey[User](
		rows,
		usersListCacheKey,
		func(user *User) error {
			users = append(users, *user)
			return nil
		},
	)
	if err != nil {
		return nil, err
	}

	return users, nil
}
