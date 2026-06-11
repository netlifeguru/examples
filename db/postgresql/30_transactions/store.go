package main

import (
	"context"
	"errors"
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

const insertUserQuery = `
	INSERT INTO users (name, email, active)
	VALUES ($1, $2, $3)
	RETURNING id
`

const updateUserQuery = `
	UPDATE users
	SET name = $1, email = $2, active = $3
	WHERE id = $4
`

const selectUserQuery = `
	SELECT *
	FROM users
	WHERE id = $1
	LIMIT 1
`

const deleteUserQuery = `
	DELETE FROM users
	WHERE id = $1
`

func RunUserTransaction(ctx context.Context, conn db.Conn) (int64, error) {
	var insertedID int64

	err := conn.TransactionCtx(ctx, func(tx db.Conn) error {
		id, found, err := db.Value[int64](
			ctx,
			tx,
			insertUserQuery,
			"Transaction User",
			"transaction.user@example.com",
			true,
		)
		if err != nil {
			return err
		}

		if !found {
			return errors.New("insert did not return id")
		}

		insertedID = id

		if _, err := db.Update(
			ctx,
			tx,
			updateUserQuery,
			"Updated Transaction User",
			"updated.transaction.user@example.com",
			false,
			insertedID,
		); err != nil {
			return err
		}

		user, found, err := db.Get[User](ctx, tx, selectUserQuery, insertedID)
		if err != nil {
			return err
		}

		if found {
			fmt.Printf("%d | %s | %s | active=%v | created_at=%s\n",
				user.ID,
				user.Name,
				user.Email,
				user.Active,
				user.CreatedAt.Format("2006-01-02 15:04:05"),
			)
		}

		if _, err := db.Delete(ctx, tx, deleteUserQuery, insertedID); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return insertedID, nil
}
