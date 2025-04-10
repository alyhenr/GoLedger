package ledger

import (
	"time"

	"github.com/google/uuid"
)

type Entry struct {
	ID        uuid.UUID
	AccountID string
	Amount    float64 // Can be negative (debit) or positive (credit)
	Timestamp time.Time
	Reference string // Optional external ref
}

type Transaction struct {
	ID        uuid.UUID
	Entries   []Entry // Must always have 2+ entries that sum to 0
	CreatedAt time.Time
}
