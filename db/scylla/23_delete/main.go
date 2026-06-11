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

	id := "79f90f76-54ab-11f1-bac1-4ac3b511b962"
	email := "john@example.com"

	if err := DeleteUser(ctx, conn, id, email); err != nil {
		log.Fatal(err)
	}

	fmt.Println("user deleted")
}
