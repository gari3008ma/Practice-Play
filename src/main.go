package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Structs for data models
type Account struct {
	AccountID      int    `json:"account_id"`
	DocumentNumber string `json:"document_number"`
}

type OperationType struct {
	OperationTypeID int    `json:"operation_type_id"`
	Description     string `json:"description"`
}

type Transaction struct {
	TransactionID   int       `json:"transaction_id"`
	AccountID       int       `json:"account_id"`
	OperationTypeID int       `json:"operation_type_id"`
	Amount          float64   `json:"amount"`
	EventDate       time.Time `json:"event_date"`
}

// In-memory data storage
var (
	accounts       = make(map[int]Account)
	transactions   = []Transaction{}
	currentAccID   = 1
	currentTransID = 1
)

// Handlers
func createAccountHandler(w http.ResponseWriter, r *http.Request) {
	var account Account
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	account.AccountID = currentAccID
	accounts[currentAccID] = account
	currentAccID++

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(account)
}

func getAccountHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	accountID, err := strconv.Atoi(params["accountId"])
	if err != nil {
		http.Error(w, "Invalid account ID", http.StatusBadRequest)
		return
	}

	account, exists := accounts[accountID]
	if !exists {
		http.Error(w, "Account not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(account)
}

func createTransactionHandler(w http.ResponseWriter, r *http.Request) {
	var transaction Transaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if _, exists := accounts[transaction.AccountID]; !exists {
		http.Error(w, "Account not found", http.StatusNotFound)
		return
	}

	transaction.TransactionID = currentTransID
	transaction.EventDate = time.Now()
	transactions = append(transactions, transaction)
	currentTransID++

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transaction)
}

func main() {
	r := mux.NewRouter()

	// API Endpoints
	r.HandleFunc("/accounts", createAccountHandler).Methods("POST")
	r.HandleFunc("/accounts/{accountId}", getAccountHandler).Methods("GET")
	r.HandleFunc("/transactions", createTransactionHandler).Methods("POST")

	// Start the server
	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", r)
}
