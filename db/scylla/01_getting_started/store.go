package main

import (
	"context"
	"github.com/netlifeguru/db"
	"time"
)

type UsersByEmail struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Active    bool      `db:"active"`
	CreatedAt time.Time `db:"created_at"`
}

func getUsers(ctx context.Context, conn db.Conn) ([]UsersByEmail, error) {

	query := `SELECT * FROM users_by_email `

	users, err := db.List[UsersByEmail](ctx, conn, query)

	if err != nil {
		return nil, err
	}

	return users, nil
}
