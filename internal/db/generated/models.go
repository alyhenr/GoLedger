// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

import (
	"database/sql"

	"github.com/google/uuid"
)

type Account struct {
	ID        uuid.UUID    `json:"id"`
	Name      string       `json:"name"`
	CreatedAt sql.NullTime `json:"created_at"`
}

type Entry struct {
	ID            uuid.UUID      `json:"id"`
	TransactionID uuid.UUID      `json:"transaction_id"`
	AccountID     uuid.UUID      `json:"account_id"`
	Amount        string         `json:"amount"`
	Description   sql.NullString `json:"description"`
	CreatedAt     sql.NullTime   `json:"created_at"`
}

type Transaction struct {
	ID        uuid.UUID      `json:"id"`
	Reference sql.NullString `json:"reference"`
	CreatedAt sql.NullTime   `json:"created_at"`
}
