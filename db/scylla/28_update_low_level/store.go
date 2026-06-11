package main

import (
	"context"

	"github.com/netlifeguru/db"
)

const updateUserByIDQuery = `
	UPDATE users_by_id
	SET name = ?, active = ?
	WHERE id = ?
`

const updateUserByEmailQuery = `
	UPDATE users_by_email
	SET name = ?, active = ?
	WHERE email = ?
`

func UpdateUser(ctx context.Context, conn db.Conn, id string, email string, name string, active bool) error {
	q, err := db.Raw(updateUserByIDQuery, name, active, id)
	if err != nil {
		return err
	}

	if _, err := conn.ExecCtx(ctx, q); err != nil {
		return err
	}

	q, err = db.Raw(updateUserByEmailQuery, name, active, email)
	if err != nil {
		return err
	}

	if _, err := conn.ExecCtx(ctx, q); err != nil {
		return err
	}

	return nil
}
