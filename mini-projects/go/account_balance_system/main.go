package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

func demoAccountBalanceSystem() {
	fmt.Printf("=== Demonstrating Thread-Safe Account Operations ===\n")

	bank := NewBank()

	err := bank.CreateAccount("acc1", 1000.0)
	if err != nil {
		log.Fatal(err)
	}

	err = bank.CreateAccount("acc2", 500.0)
	if err != nil {
		log.Fatal(err)
	}

	acc1, _ := bank.GetAccount("acc1")
	acc2, _ := bank.GetAccount("acc2")

	fmt.Printf("Initial balances:\n")
	fmt.Printf("Account 1: %.2f\n", acc1.GetBalance())
	fmt.Printf("Account 2: %.2f\n", acc2.GetBalance())
	fmt.Printf("Total: %.2f\n\n", bank.GetTotalBalance())

	// simulate concurrent operations
	const numGoroutines = 10
	const operationsPerGoroutine = 100

	var wg sync.WaitGroup

	fmt.Printf("Starting %d goroutines with %d operations each...\n", numGoroutines, operationsPerGoroutine)

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)

		go func(workerID int) {
			defer wg.Done()

			for j := 0; j < operationsPerGoroutine; j++ {
				switch rand.Intn(4) {
				case 0:
					acc1.Deposit(1.0)
				case 1:
					acc2.Deposit(1.0)
				case 2:
					acc1.Withdraw(0.5)
				case 3:
					if rand.Intn(2) == 0 {
						acc1.Transfer(acc2, 0.25)
					} else {
						acc2.Transfer(acc1, 0.25)
					}
				}
			}
		}(i)
	}

	wg.Wait()

	fmt.Printf("\nFinal balances after concurrent operations:\n")
	fmt.Printf("Account 1: %.2f\n", acc1.GetBalance())
	fmt.Printf("Account 2: %.2f\n", acc2.GetBalance())
	fmt.Printf("Total: %.2f\n", bank.GetTotalBalance())

	// demonstrate error handling
	fmt.Printf("\n=== Error Handling Demonstration ===\n")

	err = acc1.Withdraw(10000.0)
	if err == ErrInsufficientFunds {
		fmt.Printf("✓ Correctly prevented overdraft: %v\n", err)
	}

	err = acc1.Deposit(-100.0)
	if err == ErrInvalidAmount {
		fmt.Printf("✓ Correctly rejected negative deposit: %v\n", err)
	}

	_, err = bank.GetAccount("noneexistent")
	if err == ErrAccountNotFound {
		fmt.Printf("✓ Correctly handled missing account: %v\n", err)
	}
}

func main() {
	// rand.Seed(time.Now().UnixNano())
	rand.New(rand.NewSource(time.Now().UnixNano()))

	demoAccountBalanceSystem()
}
