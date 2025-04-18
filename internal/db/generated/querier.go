// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type Querier interface {
	CreateAccount(ctx context.Context, name string) (Account, error)
	CreateEntry(ctx context.Context, arg CreateEntryParams) (Entry, error)
	CreateTransaction(ctx context.Context, reference sql.NullString) (Transaction, error)
	GetAccount(ctx context.Context, id uuid.UUID) (Account, error)
	GetEntry(ctx context.Context, id uuid.UUID) (Entry, error)
	GetTranscation(ctx context.Context, id uuid.UUID) (Transaction, error)
	ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Account, error)
	ListEntries(ctx context.Context, arg ListEntriesParams) ([]Entry, error)
	ListTransactions(ctx context.Context, arg ListTransactionsParams) ([]Transaction, error)
}

var _ Querier = (*Queries)(nil)
