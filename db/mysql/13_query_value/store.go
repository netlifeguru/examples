package main

import (
	"context"

	"github.com/netlifeguru/db"
)

const countUsersQuery = `
	SELECT COUNT(*)
	FROM users
`

func CountUsers(ctx context.Context, conn db.Conn) (int64, bool, error) {
	q, err := db.Raw(countUsersQuery)
	if err != nil {
		return 0, false, err
	}

	return db.ValueQuery[int64](ctx, conn, q)
}
