package main

import (
	"errors"
	"sync"
)

type Bank struct {
	accounts map[string]*Account
	mu       sync.RWMutex // protect the accounts map
}

func NewBank() *Bank {
	return &Bank{
		accounts: make(map[string]*Account),
	}
}

func (b *Bank) CreateAccount(id string, initialBalance float64) error {
	account, err := NewAccount(id, initialBalance)
	if err != nil {
		return err
	}

	b.mu.Lock()
	defer b.mu.Unlock()

	if _, exists := b.accounts[id]; exists {
		return errors.New("account already exists")
	}

	b.accounts[id] = account
	return nil
}

func (b *Bank) GetAccount(id string) (*Account, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	account, exists := b.accounts[id]
	if !exists {
		return nil, ErrAccountNotFound
	}

	return account, nil
}

func (b *Bank) GetTotalBalance() float64 {
	b.mu.RLock()
	defer b.mu.RUnlock()

	var total float64

	accounts := make([]*Account, 0, len(b.accounts))
	for _, account := range b.accounts {
		accounts = append(accounts, account)
	}

	for _, account := range accounts {
		total += account.GetBalance()
	}

	return total
}

type AccountSnapshot struct {
	ID      string
	Balance float64
}

func (b *Bank) GetAllAccountSnapshots() []AccountSnapshot {
	b.mu.RLock()
	defer b.mu.RUnlock()

	snapshots := make([]AccountSnapshot, 0, len(b.accounts))

	for _, account := range b.accounts {
		snapshots = append(snapshots, AccountSnapshot{
			ID:      account.GetID(),
			Balance: account.GetBalance(),
		})
	}

	return snapshots
}
