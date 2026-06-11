package main

import (
	"context"
	"time"

	"github.com/netlifeguru/db"
)

type User struct {
	Email     string    `db:"email"`
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	Active    bool      `db:"active"`
	CreatedAt time.Time `db:"created_at"`
}

const selectUserByEmailQuery = `
	SELECT *
	FROM users_by_email
	WHERE email = ?
	LIMIT 1
`

func GetUserByEmail(ctx context.Context, conn db.Conn, email string) (*User, error) {
	q, err := db.Raw(selectUserByEmailQuery, email)
	if err != nil {
		return nil, err
	}

	return db.GetPtrQuery[User](ctx, conn, q)
}
