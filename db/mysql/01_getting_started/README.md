---
title: Getting Started
description: Learn how to connect to MySQL, create a connection pool, and map SQL rows into Go structs using the NetLife Guru DB and MySQL packages.
---

## Quick Start

This example shows how to connect to MySQL, create a database connection, and fetch rows into typed Go structs.

The example is split into three files:

- `main.go` - application entry point
- `connect.go` - database connection setup
- `store.go` - query logic and result mapping

### Install

```bash
go get github.com/netlifeguru/mysql
```

The MySQL package includes the required DB and mapper dependencies.

### Environment variables

```env
DB_HOST=localhost
DB_NAME=app
DB_USER=root
DB_PASSWORD=secret
```

### main.go

```go
package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"

	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()

	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found, using system environment variables")
	}

	conn, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}

	users, err := getUsers(ctx, conn)
	if err != nil {
		slog.Error("failed to get users", "error", err)
		return
	}

	fmt.Println(users)
}
```

### connect.go

```go
package main

import (
	"os"

	"github.com/netlifeguru/db"
	"github.com/netlifeguru/mysql"
)

func connectDB() (db.Conn, error) {
	conn := mysql.New()

	cfg := db.Config{
		Host:     os.Getenv("DB_HOST"),
		Database: os.Getenv("DB_NAME"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	if err := conn.CreatePool(cfg); err != nil {
		return nil, err
	}

	return conn.Fork(), nil
}
```

### store.go

```go
package main

import (
	"context"
	"time"

	"github.com/netlifeguru/db"
)

type User struct {
	ID        int64    `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Active    bool      `db:"active"`
	CreatedAt time.Time `db:"created_at"`
}

func getUsers(ctx context.Context, conn db.Conn) ([]User, error) {
	query := `SELECT * FROM users ORDER BY created_at DESC LIMIT ?`

	return db.List[User](ctx, conn, query, 10)
}
```

### Run

```bash
go run .
```

`List` executes the SQL query, passes query arguments safely, and maps the result rows into a typed `[]User` slice using `db` struct tags.