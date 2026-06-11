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

const insertUserQuery = `
	INSERT INTO users (name, email, active)
	VALUES (?, ?, ?)
`

const updateUserQuery = `
	UPDATE users
	SET name = ?, email = ?, active = ?
	WHERE id = ?
`

const selectUserQuery = `
	SELECT *
	FROM users
	WHERE id = ?
	LIMIT 1
`

const deleteUserQuery = `
	DELETE FROM users
	WHERE id = ?
`

func RunUserTransaction(ctx context.Context, conn db.Conn) (int64, error) {
	var insertedID int64

	err := conn.TransactionCtx(ctx, func(tx db.Conn) error {
		q, err := db.Raw(insertUserQuery, "Transaction User", "transaction.user@example.com", true)
		if err != nil {
			return err
		}

		result, err := tx.ExecCtx(ctx, q)
		if err != nil {
			return err
		}

		insertedID = result.LastInsertId()

		q, err = db.Raw(updateUserQuery, "Updated Transaction User", "updated.transaction.user@example.com", false, insertedID)
		if err != nil {
			return err
		}

		if _, err := tx.ExecCtx(ctx, q); err != nil {
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

		q, err = db.Raw(deleteUserQuery, insertedID)

		if err != nil {
			return err
		}

		if _, err := tx.ExecCtx(ctx, q); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return insertedID, nil
}
