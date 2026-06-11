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

	result, err := DeleteUser(ctx, conn, queries, 28)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("rows_affected=%d\n", result.RowsAffected())
}
