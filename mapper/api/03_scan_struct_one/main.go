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

	user, err := getUserByID(db, 1)
	if err != nil {
		if errors.Is(err, mapper.ErrNoRows) {
			log.Println("user not found")
			return
		}

		log.Fatal(err)
	}

	fmt.Printf(
		"ID: %d | Name: %s | Email: %s | Active: %t | Created: %s\n",
		user.ID,
		user.Name,
		user.Email,
		user.Active,
		user.CreatedAt.Format("2006-01-02 15:04:05"),
	)
}
