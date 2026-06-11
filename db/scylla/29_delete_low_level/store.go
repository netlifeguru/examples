package main

import (
	"context"

	"github.com/netlifeguru/db"
)

const deleteUserByIDQuery = `
	DELETE FROM users_by_id
	WHERE id = ?
`

const deleteUserByEmailQuery = `
	DELETE FROM users_by_email
	WHERE email = ?
`

func DeleteUser(ctx context.Context, conn db.Conn, id string, email string) error {
	q, err := db.Raw(deleteUserByIDQuery, id)
	if err != nil {
		return err
	}

	if _, err := conn.ExecCtx(ctx, q); err != nil {
		return err
	}

	q, err = db.Raw(deleteUserByEmailQuery, email)
	if err != nil {
		return err
	}

	if _, err := conn.ExecCtx(ctx, q); err != nil {
		return err
	}

	return nil
}
