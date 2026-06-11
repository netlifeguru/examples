package main

import (
	"context"

	"github.com/netlifeguru/db"
)

func DeleteUser(ctx context.Context, conn db.Conn, queries Queries, id int64) (db.Result, error) {

	res, err := db.DeleteDialect(ctx, conn, queries.DeleteUser, id)

	if err != nil {
		return nil, err
	}

	return res, nil
}
