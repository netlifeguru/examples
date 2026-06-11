package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gocql/gocql"
	"github.com/netlifeguru/db"
	scylla "github.com/netlifeguru/db-scylla"
)

const insertUserByIDQuery = `
	INSERT INTO users_by_id (id, email, name, active, created_at)
	VALUES (?, ?, ?, ?, ?)
`

const insertUserByEmailQuery = `
	INSERT INTO users_by_email (email, id, name, active, created_at)
	VALUES (?, ?, ?, ?, ?)
`

func InsertUserBatch(ctx context.Context, conn db.Conn, name string, email string, active bool) (string, error) {
	batchConn, ok := conn.(scylla.BatchConn)
	if !ok {
		return "", fmt.Errorf("connection does not support scylla batches")
	}

	id := gocql.TimeUUID()
	createdAt := time.Now().UTC()

	batch := batchConn.NewLoggedBatch(ctx)

	if err := batch.AddSQL(insertUserByIDQuery, id, email, name, active, createdAt); err != nil {
		return "", err
	}

	if err := batch.AddSQL(insertUserByEmailQuery, email, id, name, active, createdAt); err != nil {
		return "", err
	}

	if err := batch.Execute(); err != nil {
		return "", err
	}

	return id.String(), nil
}
