// Package tests contains integration tests for the API handlers.
package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/adnanlabib1509/go-transaction-engine/internal/api"
	"github.com/adnanlabib1509/go-transaction-engine/internal/models"
	"github.com/adnanlabib1509/go-transaction-engine/internal/store"
	"github.com/adnanlabib1509/go-transaction-engine/pkg/logger"
)

// TestCreateAccount tests the creation of a new account via the API.
func TestCreateAccount(t *testing.T) {
	// Initialize dependencies
	s := store.NewMemoryStore()
	l := logger.NewSimpleLogger()
	handler := api.NewHandler(s, l)

	// Create a test account
	account := &models.Account{
		ID:      "test123",
		Name:    "Test Account",
		Balance: 1000,
	}

	// Prepare the request
	body, _ := json.Marshal(account)
	req, err := http.NewRequest("POST", "/account", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	// Perform the request
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// Check the response body
	var responseAccount models.Account
	json.Unmarshal(rr.Body.Bytes(), &responseAccount)
	if responseAccount.ID != account.ID {
		t.Errorf("handler returned unexpected body: got %v want %v", responseAccount.ID, account.ID)
	}
}