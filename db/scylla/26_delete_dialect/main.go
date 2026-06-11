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

	err = DeleteUser(ctx, conn, queries, "01efebf6-64d4-11f1-9b56-4ac3b511b961")

	if err != nil {
		log.Fatal(err)
	}
}
