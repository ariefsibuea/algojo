package main

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

/**
 * Advanced Custom Error with Wrapped Causes
 *
 * Sometimes we need to preserve the original error while adding your own context. This file contains a more
 * sophisticated approach.
 */

// TransactionError represents database transaction errors with context
type TransactionError struct {
	Operation string    // what operation was being performed
	UserID    string    // which user's transaction
	Amount    float64   // transaction amount
	Timestamp time.Time // when it occured
	Cause     error     // the underlying error that caused this
}

func (e TransactionError) Error() string {
	return fmt.Sprintf("Transaction failed [%s] for user %s (amount: %.2f): %v",
		e.Operation, e.UserID, e.Amount, e.Cause)
}

// Unwrap returns the underlying error, enabling error chain inspection. This is important for Go 1.13+ error handling
// features.
func (e TransactionError) Unwrap() error {
	return e.Cause
}

// Is enables checking if this error matches a target error type.
func (e TransactionError) Is(target error) bool {
	_, ok := target.(TransactionError)
	return ok
}

// execDBOperation mimics a database operation that might fail.
func execDBOperation(query string) error {
	switch query {
	case "INVALID_SYNTAX":
		return fmt.Errorf("syntax error near 'INVALID'")
	case "CONNECTION_LOST":
		return sql.ErrConnDone
	case "NO_ROWS":
		return sql.ErrNoRows
	default:
		return nil
	}
}

func transerMoney(fromUserID, toUserID string, amount float64) error {
	operation := "TRANSFER"

	// Step 1: Validate sender's balance
	err := execDBOperation("SELECT balance FROM accounts WHERE user_id = ?")
	if err != nil {
		return TransactionError{
			Operation: operation + "_BALANCE_CHECK",
			UserID:    fromUserID,
			Amount:    amount,
			Timestamp: time.Now(),
			Cause:     err,
		}
	}

	// Step 2: Begin transaction
	err = execDBOperation("BEGIN TRANSACTION")
	if err != nil {
		return TransactionError{
			Operation: operation + "_BEGIN",
			UserID:    fromUserID,
			Amount:    amount,
			Timestamp: time.Now(),
			Cause:     err,
		}
	}

	// Step 3: Debit sender (simulate syntax error for demonstration)
	err = execDBOperation("INVALID_SYNTAX")
	if err != nil {
		return TransactionError{
			Operation: operation + "_DEBIT",
			UserID:    fromUserID,
			Amount:    amount,
			Timestamp: time.Now(),
			Cause:     err,
		}
	}

	fmt.Printf("Transfer of %.2f from %s to %s completed successfully\n",
		amount, fromUserID, toUserID)
	return nil
}

func demoAdvancedCustomError() {
	fmt.Println("=== Testing Advanced Custom Error Types ===")

	err := transerMoney("Alice", "Bob", 250.00)
	if err != nil {
		fmt.Printf("Transfer failed: %s\n", err.Error())

		var transErr TransactionError

		if errors.As(err, &transErr) {
			fmt.Printf("Error details:\n")
			fmt.Printf("  Operation: %s\n", transErr.Operation)
			fmt.Printf("  User: %s\n", transErr.UserID)
			fmt.Printf("  Amount: %.2f\n", transErr.Amount)
			fmt.Printf("  Time: %s\n", transErr.Timestamp.Format(time.RFC3339))

			// Check the underlying cause
			if transErr.Cause != nil {
				fmt.Printf("  Root cause: %s\n", transErr.Cause.Error())
			}
		}
	}
}
