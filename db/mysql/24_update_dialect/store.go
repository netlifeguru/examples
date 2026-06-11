package main

import (
	"context"

	"github.com/netlifeguru/db"
)

func UpdateUser(ctx context.Context, conn db.Conn, queries Queries, id int64, name string, email string, active bool) (db.Result, error) {
	res, err := db.UpdateDialect(ctx, conn, queries.UpdateUser, name, email, active, id)

	if err != nil {
		return nil, err
	}

	return res, nil
}
