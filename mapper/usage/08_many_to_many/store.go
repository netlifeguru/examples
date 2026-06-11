package main

import (
	"database/sql"

	"github.com/netlifeguru/mapper"
)

type PostTag struct {
	PostID    string `db:"post_id"`
	PostTitle string `db:"post_title"`
	TagID     string `db:"tag_id"`
	TagName   string `db:"tag_name"`
}

func getPostTags(db *sql.DB) ([]PostTag, error) {
	rows, err := db.Query(`
		SELECT
			p.id AS post_id,
			p.title AS post_title,
			t.id AS tag_id,
			t.name AS tag_name
		FROM posts p
		JOIN post_tags pt ON pt.post_id = p.id
		JOIN tags t ON t.id = pt.tag_id
		ORDER BY p.created_at DESC, t.name ASC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var postTags []PostTag

	err = mapper.ScanStructRows[PostTag](rows, func(postTag *PostTag) error {
		postTags = append(postTags, *postTag)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return postTags, nil
}
