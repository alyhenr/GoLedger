package main

import (
	"context"
	"database/sql"
	"log"

	db "github.com/alyhenr/GoLedger/internal/db/generated"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgres://postgres:postgres@localhost:5433/goledger?sslmode=disable"
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.New(conn)

	// Test it
	listAccountsParams := db.ListAccountsParams{
		Limit:  10,
		Offset: 0,
	}
	account, err := store.ListAccounts(context.Background(), listAccountsParams)
	if err != nil {
		log.Println("Error fetching account:", err)
	} else {
		log.Printf("Account: %+v\n", account)
	}
}
