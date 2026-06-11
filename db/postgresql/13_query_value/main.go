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

	count, found, err := CountUsers(ctx, conn)
	if err != nil {
		log.Fatal(err)
	}

	if !found {
		log.Println("no scalar value was returned")
		return
	}

	fmt.Printf("users count: %d\n", count)
}
