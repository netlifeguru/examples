package main

import (
	"context"
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

func LoadQueries(conn db.Conn) (Queries, error) {
	var queries Queries

	if err := db.LoadModel(conn, ".", &queries); err != nil {
		return Queries{}, err
	}

	return queries, nil
}

func ListUsers(ctx context.Context, conn db.Conn, queries Queries, limit int) ([]User, error) {
	q, err := db.Dialect(conn, queries.ListUsers, limit)

	if err != nil {
		return nil, err
	}

	return db.ListQuery[User](ctx, conn, q)
}

func GetUser(ctx context.Context, conn db.Conn, queries Queries, id int64) (User, bool, error) {

	q, err := db.Dialect(conn, queries.GetUser, id)

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
