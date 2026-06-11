package main

import (
	"context"

	"github.com/netlifeguru/db"
)

const updateUserQuery = `
	UPDATE users
	SET name = $1, email = $2, active = $3
	WHERE id = $4
`

func UpdateUser(ctx context.Context, conn db.Conn, id int, name string, email string, active bool) (db.Result, error) {
	q, err := db.Raw(updateUserQuery, name, email, active, id)
	if err != nil {
		return nil, err
	}

	result, err := conn.ExecCtx(ctx, q)
	if err != nil {
		return nil, err
	}

	return result, nil
}
