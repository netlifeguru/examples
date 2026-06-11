package main

import (
	"context"
	"time"

	"github.com/netlifeguru/db"
)

type User struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Active    bool      `db:"active"`
	CreatedAt time.Time `db:"created_at"`
}

const selectUserQuery = `
	SELECT *
	FROM users
	WHERE id = $1
	LIMIT 1
`

func SelectUser(ctx context.Context, conn db.Conn, id int) (*User, error) {
	return db.GetPtr[User](ctx, conn, selectUserQuery, id)
}
