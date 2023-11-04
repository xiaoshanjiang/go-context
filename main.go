package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/mysql")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	// Set a 2 seconds timeout for database query
	// if the query takes more than 2 seconds to run,
	// it will receieve a cancellation signal from the context and interupt the db query
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel()

	rows, err := db.QueryContext(ctx, "SELECT * FROM users")

	if err != nil {
		fmt.Println("Error when querying: ", err)
		return
	}

	defer rows.Close()

	// handle the query result
}
