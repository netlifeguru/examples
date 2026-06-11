package main

import (
	"context"
	"time"

	"github.com/netlifeguru/db"
)

type Post struct {
	UserID    string    `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
	ID        string    `db:"id"`
	Title     string    `db:"title"`
	Body      string    `db:"body"`
	Published bool      `db:"published"`
}

type User struct {
	Email     string    `db:"email"`
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	Active    bool      `db:"active"`
	CreatedAt time.Time `db:"created_at"`
}

type Queries struct {
	ListPostsByUser  db.DialectSQL `json:"ListPostsByUser"`
	GetUserByEmail   db.DialectSQL `json:"GetUserByEmail"`
	CountPostsByUser db.DialectSQL `json:"CountPostsByUser"`
}

func LoadQueries(conn db.Conn) (Queries, error) {
	var queries Queries

	if err := db.LoadModel(conn, ".", &queries); err != nil {
		return Queries{}, err
	}

	return queries, nil
}

func ListPostsByUser(ctx context.Context, conn db.Conn, queries Queries, userID string, limit int) ([]Post, error) {
	q, err := db.Dialect(conn, queries.ListPostsByUser, userID, limit)
	if err != nil {
		return nil, err
	}

	return db.ListQuery[Post](ctx, conn, q)
}

func GetUserByEmail(ctx context.Context, conn db.Conn, queries Queries, email string) (User, bool, error) {
	q, err := db.Dialect(conn, queries.GetUserByEmail, email)
	if err != nil {
		return User{}, false, err
	}

	return db.GetQuery[User](ctx, conn, q)
}

func CountPostsByUser(ctx context.Context, conn db.Conn, queries Queries, userID string) (int64, bool, error) {
	q, err := db.Dialect(conn, queries.CountPostsByUser, userID)
	if err != nil {
		return 0, false, err
	}

	return db.ValueQuery[int64](ctx, conn, q)
}
