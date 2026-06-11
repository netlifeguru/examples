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

const selectUsersQuery = `
	SELECT *
	FROM users
	ORDER BY created_at DESC
	LIMIT $1
`

func ListUsers(ctx context.Context, conn db.Conn) ([]User, error) {
	return db.List[User](ctx, conn, selectUsersQuery, 10)
}
