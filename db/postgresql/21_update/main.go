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

	id := 10

	result, err := UpdateUser(ctx, conn, id, "Updated User", "updated.user@example.com", true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("updated rows: %d\n", result.RowsAffected())
}
