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

	posts, err := getPostsWithUsers(db)
	if err != nil {
		log.Fatal(err)
	}

	for _, post := range posts {
		fmt.Printf(
			"Post: %s | Published: %t | Author: %s <%s> | Created: %s\n",
			post.Title,
			post.Published,
			post.UserName,
			post.UserEmail,
			post.CreatedAt.Format("2006-01-02 15:04:05"),
		)
	}
}
