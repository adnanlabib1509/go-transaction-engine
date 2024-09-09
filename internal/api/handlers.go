package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/adnanlabib1509/go-transaction-engine/internal/models"
	"github.com/adnanlabib1509/go-transaction-engine/internal/store"
	"github.com/adnanlabib1509/go-transaction-engine/pkg/logger"
)

// Handler encapsulates the dependencies for API handlers.
type Handler struct {
	store  store.Store
	logger logger.Logger
}

// NewHandler creates and returns a new Handler with the given dependencies.
func NewHandler(s store.Store, l logger.Logger) http.Handler {
	h := &Handler{
		store:  s,
		logger: l,
	}

	r := mux.NewRouter()
	
	// Define routes
	r.HandleFunc("/account", h.createAccount).Methods("POST")
	r.HandleFunc("/account/{id}", h.getAccount).Methods("GET")
	r.HandleFunc("/transaction", h.processTransaction).Methods("POST")
	r.HandleFunc("/transactions", h.getTransactions).Methods("GET")

	// Apply middleware in order: CORS, Logging, Authentication, Rate Limiting
	handler := CORSMiddleware(r)
	handler = LoggingMiddleware(handler, l)
	handler = AuthenticationMiddleware(handler, l)
	handler = RateLimitingMiddleware(handler, l)

	return handler
}

// createAccount handles the creation of a new account.
func (h *Handler) createAccount(w http.ResponseWriter, r *http.Request) {
	var acc models.Account

	// Decode the request body into an Account struct
	if err := json.NewDecoder(r.Body).Decode(&acc); err != nil {
		h.logger.Error("Failed to decode account: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Attempt to create the account in the store
	if err := h.store.CreateAccount(&acc); err != nil {
		h.logger.Error("Failed to create account: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the created account
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(acc)
}

// getAccount retrieves account information for a given account ID.
func (h *Handler) getAccount(w http.ResponseWriter, r *http.Request) {
	// Extract the account ID from the URL parameters
	vars := mux.Vars(r)
	id := vars["id"]

	// Attempt to retrieve the account from the store
	acc, err := h.store.GetAccount(id)
	if err != nil {
		h.logger.Error("Failed to get account: %v", err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Return the account information
	json.NewEncoder(w).Encode(acc)
}

// processTransaction handles the processing of a new transaction.
func (h *Handler) processTransaction(w http.ResponseWriter, r *http.Request) {
	var t models.Transaction

	// Decode the request body into a Transaction struct
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		h.logger.Error("Failed to decode transaction: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Attempt to process the transaction
	if err := h.store.ProcessTransaction(&t); err != nil {
		h.logger.Error("Failed to process transaction: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Return the processed transaction
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(t)
}

// getTransactions retrieves all transactions.
func (h *Handler) getTransactions(w http.ResponseWriter, r *http.Request) {
	// Attempt to retrieve all transactions from the store
	transactions, err := h.store.GetTransactions()
	if err != nil {
		h.logger.Error("Failed to get transactions: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the list of transactions
	json.NewEncoder(w).Encode(transactions)
}