package main

import (
	"fmt"
	"slices"
	"time"
)

/**
 * Basic Custom Error Implementation
 *
 * This file contains example of simple custom error type.
 */

// PaymentError represents errors that occur during payment processing.
type PaymentError struct {
	Code      string    // error code for logging and debugging
	Message   string    // human-readable error message
	Timestamp time.Time // when the error occured
	UserID    string    // which user encountered this error
}

// Error implements the error interface. This method is required for any type to be considered an error.
func (e PaymentError) Error() string {
	return fmt.Sprintf("Payment Error [%s]: %s (User: %s, Time %s)",
		e.Code, e.Message, e.UserID, e.Timestamp.Format(time.RFC3339),
	)
}

// IsRetryable determines if this error indicates a retriable operation. This method provides additional behaviour
// beyond the basic error interface.
func (e PaymentError) IsRetryable() bool {
	retryableCodes := []string{"NETWORK_TIMEOUT", "TEMPORARY_UNAVAILABLE", "RATE_LIMITED"}
	// NOTE: old style
	// for _, code := range retryableCodes {
	// 	if e.Code == code {
	// 		return true
	// 	}
	// }
	if slices.Contains(retryableCodes, e.Code) {
		return true
	}
	return false
}

// processPayment simulates a payment processing function.
func processPayment(userID string, amount float64) error {
	if amount <= 0 {
		return PaymentError{
			Code:      "INVALID_AMOUT",
			Message:   "Payment amount must be greater than zero",
			Timestamp: time.Now(),
			UserID:    userID,
		}
	}

	if amount > 10000 {
		return PaymentError{
			Code:      "AMOUNT_EXCEEDED",
			Message:   "Payment amount exceeds daily limit",
			Timestamp: time.Now(),
			UserID:    userID,
		}
	}

	// simulate network timeout
	if userID == "network_issue_user" {
		return PaymentError{
			Code:      "NETWORK_TIMEOUT",
			Message:   "Connection to payment gateway timed out",
			Timestamp: time.Now(),
			UserID:    userID,
		}
	}

	fmt.Printf("Payment of %.2f processed successfully for user %s\n", amount, userID)
	return nil
}

func demoSimpleCustomError() {
	testCases := []struct {
		userID string
		amount float64
	}{
		{"user123", 100.50},
		{"user456", -50.00},
		{"user789", 15000.00},
		{"network_issue_user", 200.00},
	}

	for _, tc := range testCases {
		fmt.Printf("\n--- Processing payment for %s ---\n", tc.userID)

		err := processPayment(tc.userID, tc.amount)
		if err != nil {
			if paymentErr, ok := err.(PaymentError); ok {
				fmt.Printf("Error occured: %s\n", paymentErr.Error())

				if paymentErr.IsRetryable() {
					fmt.Println("This error is retryable. Will attempt retry...")
				} else {
					fmt.Println("This error is not retryable. Manual intervention required.")
				}
			} else {
				fmt.Printf("Unexpected error type: %s\n", paymentErr.Error())
			}
		}
	}
}
