package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
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

	users, err := getUserDTOs(db)
	if err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		fmt.Printf(
			"UserID: %s | FullName: %s | EmailAddr: %s\n",
			user.UserID,
			user.FullName,
			user.EmailAddr,
		)
	}
}
