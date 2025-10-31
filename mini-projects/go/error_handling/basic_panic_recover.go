package main

import (
	"fmt"
	"runtime/debug"
	"strings"
	"time"
)

// CriticalError represents error that might cause panics
type CriticalError struct {
	Component string
	Message   string
	Timestamp time.Time
	Stack     []byte
}

func (e CriticalError) Error() string {
	return fmt.Sprintf("CRITICAL [%s]: %s at %s", e.Component, e.Message, e.Timestamp.Format(time.RFC3339))
}

// safeExecute wraps potentially panicking operations with recovery
func safeExecute(operation string, fn func() error) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = CriticalError{
				Component: operation,
				Message:   fmt.Sprintf("Panic recovered: %v", r),
				Timestamp: time.Now(),
				Stack:     debug.Stack(),
			}

			fmt.Printf("  -> PANIC RECOVERED in %s: %v\n", operation, r)
			// fmt.Printf("Stack trace:\n%s\n", debug.Stack())
		}
	}()

	// execute the potentially panicking function
	err = fn()
	return err
}

// riskyDatabaseOperation simulates operations that might panic
func riskyDatabaseOperation(query string) error {
	fmt.Printf("  Executing database query: %s\n", query)

	switch query {
	case "PANIC_TEST":
		panic("simulated division by zero")
	case "ARRAY_BOUNDS":
		// simulate panic due to array bounds violation
		arr := []int{1, 2, 3}
		fmt.Println(arr[10]) // This will panic
	case "DIVIDE_BY_ZERO":
		// NOTE: This doesn't actually panic in Go, but let's simulate it
		x := 0
		if x == 0 {
			panic("simulated division by zero")
		}
	case "VALID_QUERY":
		fmt.Println("  -> Query executed successfully")
		return nil
	default:
		return fmt.Errorf("unknown query: %s", query)
	}

	return nil
}

// processTransactionWithRecovery demonstrates panic recovery in financial operations
func processTransactionWithRecovery(transactionID string, operations []string) error {
	fmt.Printf("\n=== Processing Transaction %s ===\n", transactionID)

	for i, operation := range operations {
		step := fmt.Sprintf("Transaction_%s_Step_%d", transactionID, i+1)
		fmt.Printf("Step %d: %s\n", i+1, operation)

		err := safeExecute(step, func() error {
			return riskyDatabaseOperation(operation)
		})

		if err != nil {
			// check if this was a critical error (panic recovery)
			if critErr, ok := err.(CriticalError); ok {
				fmt.Printf("  -> CRITICAL ERROR in step %d: %s\n", i+1, critErr.Message)
				fmt.Printf("  -> Transaction %s TERMINATED due to critical error\n", transactionID)

				// Return error to stop THIS transaction, but allow program to continue
				return fmt.Errorf("transaction %s failed due to critical system error at step %d", transactionID, i+1)
			} else {
				// Handle regular errors
				fmt.Printf("  -> Regular error in step %d: %s\n", i+1, err.Error())
				return fmt.Errorf("transaction %s failed at step %d: %w", transactionID, i+1, err)
			}
		}

		// completedSteps = append(completedSteps, step)
		fmt.Printf("  -> Step %d completed successfully\n\n", i+1)
	}

	fmt.Printf("  -> Transaction %s completed successfully\n", transactionID)
	return nil
}

func demoBasicPanicRecover() {
	fmt.Println("=== Basic Panic/Recover Implementation ===")
	fmt.Println("Demonstrating that each transaction is processed independently")

	testCases := []struct {
		id         string
		operations []string
	}{
		{
			id:         "TXN001",
			operations: []string{"VALID_QUERY", "VALID_QUERY", "VALID_QUERY"},
		},
		{
			id:         "TXN002",
			operations: []string{"VALID_QUERY", "PANIC_TEST", "VALID_QUERY"},
		},
		{
			id:         "TXN003",
			operations: []string{"VALID_QUERY", "ARRAY_BOUNDS", "VALID_QUERY"},
		},
		{
			id:         "TXN004",
			operations: []string{"DIVIDE_BY_ZERO"},
		},
	}

	// track overall system state
	totalTransactions := len(testCases)
	successfulTransactions := 0

	// Process each transaction independently
	for i, ts := range testCases {
		fmt.Printf("\n>>> Starting transaction %d of %d\n", i+1, totalTransactions)

		// Process this specific transaction
		err := processTransactionWithRecovery(ts.id, ts.operations)

		// Handle the result for this transaction
		if err != nil {
			fmt.Printf("❌ Final result for %s: FAILED - %s\n", ts.id, err.Error())
		} else {
			fmt.Printf("✅ Final result for %s: SUCCESS\n", ts.id)
			successfulTransactions++
		}

		// Show that we're continuing to the next transaction
		if i < len(testCases)-1 {
			fmt.Printf(">>> System continues to next transaction...\n")
		}
	}

	// Show overall system resilience
	fmt.Printf("\n" + strings.Repeat("=", 50))
	fmt.Printf("\n🎯 SYSTEM SUMMARY:")
	fmt.Printf("\n   Total transactions processed: %d", totalTransactions)
	fmt.Printf("\n   Successful transactions: %d", successfulTransactions)
	fmt.Printf("\n   Failed transactions: %d", totalTransactions-successfulTransactions)
	fmt.Printf("\n   System status: OPERATIONAL (failures were isolated)")
	fmt.Printf("\n" + strings.Repeat("=", 50) + "\n")
}
