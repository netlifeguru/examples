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

func SelectUser(ctx context.Context, conn db.Conn, queries Queries, id int) (*User, error) {
	return db.GetPtrDialect[User](ctx, conn, queries.GetUser, id)
}

func SelectUserExp(ctx context.Context, conn db.Conn, id int) (*User, error) {
	return db.GetPtrDialect[User](ctx, conn, db.DialectSQL{
		Postgres: `SELECT * FROM users WHERE id = $1 LIMIT 1`,
		Mysql:    `SELECT * FROM users WHERE id = ? LIMIT 1`,
	}, id)
}
