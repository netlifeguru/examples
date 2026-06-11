package main

import "github.com/netlifeguru/db"

type Queries struct {
	InsertUser db.DialectSQL `db:"InsertUser"`
	UpdateUser db.DialectSQL `db:"UpdateUser"`
	DeleteUser db.DialectSQL `db:"DeleteUser"`
}

func LoadQueries(conn db.Conn) (Queries, error) {
	var queries Queries

	if err := db.LoadModel(conn, ".", &queries); err != nil {
		return Queries{}, err
	}

	return queries, nil
}
