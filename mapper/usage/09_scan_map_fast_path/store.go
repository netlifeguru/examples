package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/netlifeguru/mapper"
)

type User struct {
	ID        int64
	Name      string
	Email     string
	Active    bool
	CreatedAt time.Time
}

func (u *User) ScanMap(row mapper.Row) error {
	id, ok := row.Int64("id")
	if !ok {
		return fmt.Errorf("invalid id")
	}

	name, ok := row.String("name")
	if !ok {
		return fmt.Errorf("invalid name")
	}

	email, ok := row.String("email")
	if !ok {
		return fmt.Errorf("invalid email")
	}

	active, ok := row.Bool("active")
	if !ok {
		return fmt.Errorf("invalid active")
	}

	createdAt, ok := row.Time("created_at")
	if !ok {
		return fmt.Errorf("invalid created_at")
	}

	u.ID = id
	u.Name = name
	u.Email = email
	u.Active = active
	u.CreatedAt = createdAt

	return nil
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

	err = mapper.ScanMapRows(rows, func(row map[string]any) error {
		var user User

		if err := user.ScanMap(mapper.Row(row)); err != nil {
			return err
		}

		users = append(users, user)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return users, nil
}
