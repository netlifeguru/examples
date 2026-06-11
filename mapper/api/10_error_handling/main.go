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

	fmt.Println("Example 1: no rows")
	_, err = getMissingUser(db)
	if err != nil {
		if errors.Is(err, mapper.ErrNoRows) {
			fmt.Println("Handled:", err)
		} else {
			log.Fatal(err)
		}
	}

	fmt.Println()
	fmt.Println("Example 2: too many rows")
	_, err = getOneUserWithoutLimit(db)
	if err != nil {
		if errors.Is(err, mapper.ErrTooManyRows) {
			fmt.Println("Handled:", err)
		} else {
			log.Fatal(err)
		}
	}

	fmt.Println()
	fmt.Println("Example 3: conversion error")
	_, err = getBrokenUsers(db)
	if err != nil {
		fmt.Println("Handled conversion error:")
		fmt.Println(err)
	}
}
