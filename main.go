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

	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel()

	go func() {
		<-ctx.Done()

		fmt.Println("Received cancellation signal, rolling back transaction...")

		tx.Rollback()

	}()

	// execute long running transactions
	time.Sleep(3 * time.Second)

	err = tx.Commit()
	if err != nil {
		fmt.Println("Error when committing transaction: ", err)
		return
	}

	fmt.Println("Transaction committed successfully.")
}
