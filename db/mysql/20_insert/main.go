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

	if _, err := InsertUser(ctx, conn, "John Doe", "john.doe@example.com", true); err != nil {
		log.Fatal(err)
	}

	fmt.Println("user inserted")
}
