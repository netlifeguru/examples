package main

import (
	"context"

	"github.com/netlifeguru/db"
)

func DeleteUser(ctx context.Context, conn db.Conn, queries Queries, id string) error {

	_, err := db.DeleteDialect(ctx, conn, queries.DeleteUser, id)

	if err != nil {
		return err
	}

	return nil
}
