package main

import (
	"context"
	"time"

	"github.com/netlifeguru/db"
)

type User struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Active    bool      `db:"active"`
	CreatedAt time.Time `db:"created_at"`
}

func getUsers(ctx context.Context, conn db.Conn) ([]User, error) {

	query := `SELECT * FROM users ORDER BY created_at DESC LIMIT $1`

	users, err := db.List[User](ctx, conn, query, 1)

	if err != nil {
		return nil, err
	}

	return users, nil
}
