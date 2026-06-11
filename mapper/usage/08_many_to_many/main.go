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

	postTags, err := getPostTags(db)
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range postTags {
		fmt.Printf(
			"Post: %s | Tag: %s\n",
			item.PostTitle,
			item.TagName,
		)
	}
}
