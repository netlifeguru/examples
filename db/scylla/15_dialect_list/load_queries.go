package main

import "github.com/netlifeguru/db"

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
