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

const selectUserQuery = `
	SELECT *
	FROM users
	WHERE id = ?
	LIMIT 1
`

func GetUser(ctx context.Context, conn db.Conn, id int) (*User, error) {
	q, err := db.Raw(selectUserQuery, id)
	if err != nil {
		return nil, err
	}

	return db.GetPtrQuery[User](ctx, conn, q)
}
