package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

/**
 * Basic Error Wrapping
 *
 * Error wrapping, introduced in Go 1.13, allows us to add context to errors while preserving the original error
 * information. This creates an error chain that can be inspected programmatically. The `%w` verb in `fmt.Errorf`
 * creates wrapped error that maintain the relationship between the original error and your additional context.
 */

var (
	ErrPaymentGatewayUnavailable = errors.New("payment gaetway is temporarily unavailable")
	ErrInsufficientFunds         = errors.New("insufficient funds in account")
	ErrInvalidAccount            = errors.New("account number is invalid")
)

// validateAccount simulates account validation
func validateAccount(accountID string) error {
	if accountID == "" {
		return fmt.Errorf("account validation failed: %w", ErrInvalidAccount)
	}

	if rand.Float32() < 0.3 {
		return fmt.Errorf("account validation failed for account %s: %w", accountID, ErrInvalidAccount)
	}

	return nil
}

// checkBalance simulates balance checking
func checkBalance(accountID string, amount float64) error {
	// simulate insufficient funds
	if amount > 1000 {
		return fmt.Errorf("balance check failed for account %s (requested: %.2f): %w",
			accountID, amount, ErrInsufficientFunds)
	}
	return nil
}

// callPaymentGateway simulates external payment gateway call
func callPaymentGateway(accountID string, amount float64) error {
	// simulate gateway unavailability
	if rand.Float32() < 0.4 {
		return fmt.Errorf("payment gateway request failed for account %s: %w", accountID, ErrPaymentGatewayUnavailable)
	}

	return nil
}

func processWalletPayment(accountID string, amount float64, description string) error {
	// Layer 1: Validate account
	if err := validateAccount(accountID); err != nil {
		return fmt.Errorf("wallet payment validation failed: %w", err)
	}

	// Layer 2: Check balance
	if err := checkBalance(accountID, amount); err != nil {
		return fmt.Errorf("wallet payment balance verification failed: %w", err)
	}

	// Layer 3: Process payment
	if err := callPaymentGateway(accountID, amount); err != nil {
		return fmt.Errorf("wallet payment processing failed for '%s': %w", description, err)
	}

	fmt.Printf("Payment processed successfully: %.2f for %s (Account: %s)\n",
		amount, description, accountID)
	return nil
}

func analyzeError(err error) {
	fmt.Printf("Full error: %s\n", err.Error())

	if errors.Is(err, ErrInsufficientFunds) {
		fmt.Println("Analysis: This is a funds-related issue. User needs to top up their wallet.")
	} else if errors.Is(err, ErrInvalidAccount) {
		fmt.Println("Analysis: This is an account validation issue. User should verify their account details.")
	} else if errors.Is(err, ErrPaymentGatewayUnavailable) {
		fmt.Println("Analysis: This is a temporary service issue. Payment should be retried later.")
	}

	// unwrap the error chain to see the root cause
	fmt.Println("\nError chain:")
	currentErr := err
	level := 1
	for currentErr != nil {
		fmt.Printf("  Level %d: %s\n", level, currentErr.Error())
		currentErr = errors.Unwrap(currentErr)
		level++
	}
}

func demoBasicErrorWrapping() {
	fmt.Println("=== Error Wrapping with fmt.Errorf ===")

	/* deprecated since Go1.20 */
	// rand.Seed(time.Now().UnixNano())
	// seed random number generator for consistent testing
	rand.New(rand.NewSource(time.Now().UnixNano()))

	testCases := []struct {
		accountID   string
		amount      float64
		description string
	}{
		{"ACC123", 150.00, "Coffee purchase"},
		{"", 50.00, "Invalid account test"},
		{"ACC456", 1500.00, "Expensive purchase"},
		{"ACC789", 75.00, "Normal purchase"},
	}

	for i, tc := range testCases {
		fmt.Printf("\n--- Test Case %d ---\n", i+1)
		fmt.Printf("Processing payment: %s (Account: %s, Amount: %.2f)\n", tc.description, tc.accountID, tc.amount)

		err := processWalletPayment(tc.accountID, tc.amount, tc.description)
		if err != nil {
			fmt.Println("\nError occurred:")
			analyzeError(err)
		}

		fmt.Println()
	}
}
