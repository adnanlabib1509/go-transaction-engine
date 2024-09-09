package models

import (
	"time"

	"github.com/adnanlabib1509/go-transaction-engine/pkg/utils"
)

// TransactionType represents the type of financial transaction.
type TransactionType string

const (
	Deposit  TransactionType = "deposit"
	Withdrawal TransactionType = "withdrawal"
	Transfer TransactionType = "transfer"
)

// Transaction represents a financial transaction in the system.
type Transaction struct {
	ID        string          `json:"id"`
	Type      TransactionType `json:"type"`
	Amount    float64         `json:"amount"`
	FromID    string          `json:"from_id"`
	ToID      string          `json:"to_id"`
	Timestamp time.Time       `json:"timestamp"`
}

// NewTransaction creates a new Transaction with the given details.
func NewTransaction(transactionType TransactionType, amount float64, fromID, toID string) *Transaction {
	return &Transaction{
		ID:        generateTransactionID(),
		Type:      transactionType,
		Amount:    amount,
		FromID:    fromID,
		ToID:      toID,
		Timestamp: time.Now(),
	}
}

// generateTransactionID creates a unique identifier for the transaction.
// In a real application, this might use a UUID library or database-specific ID generation.
func generateTransactionID() string {
	return time.Now().Format("20060102150405") + generateRandomString(6) // Simple ID based on timestamp + random string
}

// generateRandomString creates a random string of the specified length.
// In a real application, you'd want to use a more robust method of generating random strings.
func generateRandomString(length int) string {
	return time.Now().Format("20060102150405") + utils.GenerateRandomString(6)
}