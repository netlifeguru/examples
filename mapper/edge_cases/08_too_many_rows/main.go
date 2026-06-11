package main

import (
	"errors"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/netlifeguru/mapper"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found, I'm using system env variables")
	}

	db, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	user, err := getOneUserWithoutLimit(db)
	if err != nil {
		if errors.Is(err, mapper.ErrTooManyRows) {
			fmt.Println("Expected one user, but query returned multiple rows:", err)
			return
		}

		log.Fatal(err)
	}

	fmt.Println(user)
}
