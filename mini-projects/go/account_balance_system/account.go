package main

import (
	"errors"
	"sync"
)

var (
	ErrInsufficientFunds = errors.New("insufficient funds for withdrawal")
	ErrInvalidAmount     = errors.New("amount must be positive")
	ErrAccountNotFound   = errors.New("account not found")
)

type Account struct {
	id      string
	balance float64
	mu      sync.RWMutex // RWMutex allows multiple concurrent readers
}

// NewAccount creates a new account with initial balance. This constructor ensures proper initialization
func NewAccount(id string, initialBalance float64) (*Account, error) {
	if initialBalance < 0 {
		return nil, ErrInvalidAmount
	}

	return &Account{
		id:      id,
		balance: initialBalance,
	}, nil
}

// GetBalance returns the current balance. Uses RLock for read-only operations, allowing concurrent reads.
func (a *Account) GetBalance() float64 {
	a.mu.RLock()         // acquired read lock
	defer a.mu.RUnlock() // ensure lock is released even if panic occurs
	return a.balance
}

// GetID returns the account ID (thread-safe since string is immutable).
func (a *Account) GetID() string {
	return a.id
}

// Deposit adds money to the account. Uses full Lock since we're modifying the balance.
func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}

	a.mu.Lock()         // acquire exclusive write lock
	defer a.mu.Unlock() // ensure lock is released

	a.balance += amount
	return nil
}

// Withdraw removes money from the account with overdraft protection. This is the most critical method for thread
// safety
func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	// Check for sufficient funds after acquiring lock. This prevents race conditions between balance check and
	// withdrawal.
	if a.balance < amount {
		return ErrInsufficientFunds
	}

	a.balance -= amount
	return nil
}

// Transfer performs atomic transfer between two accounts. This demonstrates handling multiple locks safely to avoid
// deadlocks
func (a *Account) Transfer(to *Account, amount float64) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}

	// Acquire locks in consistent order to prevent deadlocks.Always lock accounts in order of their ID to ensure
	// deterministic locking.
	var first, second *Account
	if a.id < to.id {
		first, second = a, to
	} else {
		first, second = to, a
	}

	first.mu.Lock()
	defer first.mu.Unlock()

	second.mu.Lock()
	defer second.mu.Unlock()

	if a.balance < amount {
		return ErrInsufficientFunds
	}

	a.balance -= amount
	to.balance += amount

	return nil
}
