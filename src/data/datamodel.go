package models

import "time"

// Account represents a customer account.
type Account struct {
	AccountID      int    `json:"account_id"`      // Unique identifier for the account
	DocumentNumber string `json:"document_number"` // Unique document number of the customer
}

// OperationType represents the type of operation or transaction.
type OperationType struct {
	OperationTypeID int    `json:"operation_type_id"` // Unique identifier for the operation type
	Description     string `json:"description"`       // Description of the operation type
}

// Transaction represents a transaction made on an account.
type Transaction struct {
	TransactionID   int       `json:"transaction_id"`    // Unique identifier for the transaction
	AccountID       int       `json:"account_id"`        // Associated account ID
	OperationTypeID int       `json:"operation_type_id"` // Associated operation type ID
	Amount          float64   `json:"amount"`            // Transaction amount (negative for debit, positive for credit)
	EventDate       time.Time `json:"event_date"`        // Date and time of the transaction
}
