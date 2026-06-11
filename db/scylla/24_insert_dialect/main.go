package main

import (
	"context"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found, I'm using system env variables")
	}

	conn, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}

	queries, err := LoadQueries(conn)

	if err != nil {
		log.Fatal(err)
	}

	id, err := InsertUser(ctx, conn, queries, "Alice Doe", "alice.doe@example.com", true)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(id)
}
