// Package models defines the data structures used in the financial transaction system.
package models

import (
	"time"
)

// Account represents a user's account in the system.
type Account struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewAccount creates a new Account with the given name and initial balance.
func NewAccount(name string, initialBalance float64) *Account {
	now := time.Now()
	return &Account{
		ID:        generateID(), // Implement this function to generate unique IDs
		Name:      name,
		Balance:   initialBalance,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// generateID creates a unique identifier for the account.
// In a real application, this might use a UUID library or database-specific ID generation.
func generateID() string {
	return time.Now().Format("20060102150405") // Simple ID based on current timestamp
}