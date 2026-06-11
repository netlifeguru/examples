package main

import (
	"database/sql"
	"time"

	"github.com/netlifeguru/mapper"
)

type PostWithUser struct {
	PostID    string    `db:"post_id"`
	Title     string    `db:"title"`
	Body      string    `db:"body"`
	Published bool      `db:"published"`
	UserID    string    `db:"user_id"`
	UserName  string    `db:"user_name"`
	UserEmail string    `db:"user_email"`
	CreatedAt time.Time `db:"created_at"`
}

func getPostsWithUsers(db *sql.DB) ([]PostWithUser, error) {
	rows, err := db.Query(`
		SELECT
			p.id AS post_id,
			p.title,
			p.body,
			p.published,
			u.id AS user_id,
			u.name AS user_name,
			u.email AS user_email,
			p.created_at
		FROM posts p
		JOIN users u ON u.id = p.user_id
		ORDER BY p.created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []PostWithUser

	err = mapper.ScanStructRows[PostWithUser](rows, func(post *PostWithUser) error {
		posts = append(posts, *post)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return posts, nil
}
