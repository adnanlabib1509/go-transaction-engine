// Package service provides business logic for the transaction engine.
package service

import (
	"github.com/adnanlabib1509/go-transaction-engine/internal/models"
	"github.com/adnanlabib1509/go-transaction-engine/internal/store"
)

// TransactionService handles the business logic for transactions.
type TransactionService struct {
	store store.Store
}

// NewTransactionService creates a new TransactionService with the given store.
func NewTransactionService(s store.Store) *TransactionService {
	return &TransactionService{store: s}
}

// ProcessTransaction handles the business logic for processing a transaction.
// It delegates the actual storage operation to the underlying store.
func (s *TransactionService) ProcessTransaction(t *models.Transaction) error {
	// Additional business logic can be added here
	return s.store.ProcessTransaction(t)
}

// GetTransactionHistory retrieves the history of all transactions.
func (s *TransactionService) GetTransactionHistory() ([]*models.Transaction, error) {
	return s.store.GetTransactions()
}