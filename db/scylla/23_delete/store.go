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
	if _, err := db.Delete(ctx, conn, deleteUserByIDQuery, id); err != nil {
		return err
	}

	if _, err := db.Delete(ctx, conn, deleteUserByEmailQuery, email); err != nil {
		return err
	}

	return nil
}
