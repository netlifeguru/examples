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

func GetUserByEmail(ctx context.Context, conn db.Conn, queries Queries, email string) (*User, error) {
	return db.GetPtrDialect[User](ctx, conn, queries.GetUserByEmail, email)
}

func GetUserByEmailExp(ctx context.Context, conn db.Conn, queries Queries, email string) (*User, error) {
	return db.GetPtrDialect[User](ctx, conn, db.DialectSQL{
		Scylla: `SELECT * FROM users_by_email WHERE email = ? LIMIT 1`,
	}, email)
}
