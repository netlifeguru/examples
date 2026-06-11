package main

import (
	"context"
	"fmt"

	"github.com/netlifeguru/db"
)

func InsertUser(ctx context.Context, conn db.Conn, queries Queries, name string, email string, active bool) (db.Result, error) {

	result, err := db.InsertReturnDialect(ctx, conn, queries.InsertUser, name, email, active)

	if err != nil {
		return nil, err
	}

	if result.RowsAffected() != 1 {
		return nil, fmt.Errorf("expected 1 rows affected, got %d", result.RowsAffected())
	}

	return result, nil
}
