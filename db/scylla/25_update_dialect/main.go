package main

import (
	"context"
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

	err = UpdateUser(ctx, conn, queries, "22222222-2222-2222-2222-222222222222", "jane@example.com", false)

	if err != nil {
		log.Fatal(err)
	}
}
