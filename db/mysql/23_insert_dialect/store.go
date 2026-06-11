package main

import (
	"context"
	"fmt"
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

func InsertUser(ctx context.Context, conn db.Conn, queries Queries, name string, email string, active bool) (db.Result, error) {

	result, err := db.InsertDialect(ctx, conn, queries.InsertUser, name, email, active)

	if err != nil {
		return nil, err
	}

	if result.RowsAffected() != 1 {
		return nil, fmt.Errorf("expected 1 rows affected, got %d", result.RowsAffected())
	}

	return result, nil
}
