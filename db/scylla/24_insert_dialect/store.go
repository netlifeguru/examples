package main

import (
	"context"
	"time"

	"github.com/gocql/gocql"
	"github.com/netlifeguru/db"
)

func InsertUser(ctx context.Context, conn db.Conn, queries Queries, name string, email string, active bool) (string, error) {

	id := gocql.TimeUUID()

	createdAt := time.Now().UTC()

	_, err := db.InsertDialect(ctx, conn, queries.InsertUser, id, email, name, active, createdAt)

	if err != nil {
		return "", err
	}

	return id.String(), nil
}
