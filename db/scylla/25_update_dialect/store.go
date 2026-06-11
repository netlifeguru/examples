package main

import (
	"context"

	"github.com/netlifeguru/db"
)

func UpdateUser(ctx context.Context, conn db.Conn, queries Queries, id string, email string, active bool) error {

	_, err := db.UpdateDialect(ctx, conn, queries.UpdateUser, email, active, id)

	if err != nil {
		return err
	}

	return nil
}
