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

type Queries struct {
	ListUsers  db.DialectSQL `json:"ListUsers"`
	GetUser    db.DialectSQL `json:"GetUser"`
	CountUsers db.DialectSQL `json:"CountUsers"`
}

func LoadQueries(connections map[string]db.Conn, modelPath string, target any) error {
	for _, conn := range db.DistinctByDriver(connections) {
		if err := db.LoadModel(conn, modelPath, target); err != nil {
			return fmt.Errorf("load %s queries: %w", conn.DriverName(), err)
		}
	}

	return nil
}

func ListUsers(ctx context.Context, conn db.Conn, queries Queries, limit int) ([]User, error) {
	q, err := db.Dialect(conn, queries.ListUsers, limit)
	if err != nil {
		return nil, err
	}

	return db.ListQuery[User](ctx, conn, q)
}

func GetUser(ctx context.Context, conn db.Conn, queries Queries, id int64) (User, bool, error) {
	q, err := db.Dialect(conn, queries.GetUser, id, 1)
	if err != nil {
		return User{}, false, err
	}

	return db.GetQuery[User](ctx, conn, q)
}

func CountUsers(ctx context.Context, conn db.Conn, queries Queries) (int64, bool, error) {
	q, err := db.Dialect(conn, queries.CountUsers)
	if err != nil {
		return 0, false, err
	}

	return db.ValueQuery[int64](ctx, conn, q)
}
