// Package store provides data storage implementations for the financial transaction system.
package store

import (
	"errors"
	"sync"

	"github.com/adnanlabib1509/go-transaction-engine/internal/models"
)

// MemoryStore implements an in-memory storage system for accounts and transactions.
type MemoryStore struct {
	accounts     map[string]*models.Account
	transactions []*models.Transaction
	mu           sync.RWMutex // Ensures thread-safe operations on the store
}

// NewMemoryStore creates and initializes a new MemoryStore.
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		accounts:     make(map[string]*models.Account),
		transactions: make([]*models.Transaction, 0),
	}
}

// CreateAccount adds a new account to the store.
// It returns an error if an account with the same ID already exists.
func (s *MemoryStore) CreateAccount(acc *models.Account) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.accounts[acc.ID]; exists {
		return errors.New("account already exists")
	}

	s.accounts[acc.ID] = acc
	return nil
}

// GetAccount retrieves an account from the store by its ID.
// It returns an error if the account is not found.
func (s *MemoryStore) GetAccount(id string) (*models.Account, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	acc, exists := s.accounts[id]
	if !exists {
		return nil, errors.New("account not found")
	}

	return acc, nil
}

// ProcessTransaction handles a financial transaction between two accounts.
// It updates account balances and stores the transaction record.
// Returns an error if the transaction cannot be completed (e.g., insufficient funds).
func (s *MemoryStore) ProcessTransaction(t *models.Transaction) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Verify sender account exists
	from, exists := s.accounts[t.From]
	if !exists {
		return errors.New("sender account not found")
	}

	// Verify recipient account exists
	to, exists := s.accounts[t.To]
	if !exists {
		return errors.New("recipient account not found")
	}

	// Check for sufficient funds
	if from.Balance < t.Amount {
		return errors.New("insufficient funds")
	}

	// Update account balances
	from.Balance -= t.Amount
	to.Balance += t.Amount

	// Store the transaction
	s.transactions = append(s.transactions, t)

	return nil
}

// GetTransactions returns all transactions stored in the system.
func (s *MemoryStore) GetTransactions() ([]*models.Transaction, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.transactions, nil
}