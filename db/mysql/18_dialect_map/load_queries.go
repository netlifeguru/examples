package main

import "github.com/netlifeguru/db"

type Queries struct {
	ListUsers  db.DialectSQL `json:"ListUsers"`
	GetUser    db.DialectSQL `json:"GetUser"`
	CountUsers db.DialectSQL `json:"CountUsers"`
}

func LoadQueries(conn db.Conn) (Queries, error) {
	var queries Queries

	if err := db.LoadModel(conn, ".", &queries); err != nil {
		return Queries{}, err
	}

	return queries, nil
}
