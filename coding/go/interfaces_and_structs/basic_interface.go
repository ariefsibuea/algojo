package main

import (
	"errors"
	"fmt"
)

// PaymentProcessor defines the contract for any payment processing system. This interface ensures all payment methods
// follow the same structure.
type PaymentProcessor interface {
	ProcessPayment(amount float64) error
	GetTransactionFee(amount float64) float64
}

// CreditCardProcessor implements PaymentProcessor for credit card payments
type CreditCardProcessor struct {
	CardNumber string
	Provider   string
}

// ProcessPayment handles credit card payment processing. We implement this method to satisfy the PaymentProcessor
// interface.
func (c CreditCardProcessor) ProcessPayment(amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}

	fmt.Printf("Processing credit card payment of $%.2f using %s\n", amount, c.Provider)
	// in real implementation, this would connect to payment gateway
	fmt.Printf("Payment successful for card ending in %s\n", c.CardNumber[len(c.CardNumber)-4:])

	return nil
}

// GetTransactionFee calculates credit card processing fees.
func (c CreditCardProcessor) GetTransactionFee(amount float64) float64 {
	// credit cards typically have higher fees (2.9% + $0.30)
	return amount*0.029 + 0.30
}

// BankTransferProcessor implements PaymentProcessor for bank transfers.
type BankTransferProcessor struct {
	BankCode    string
	AccountType string
}

// ProcessPayment handles bank transfer payment processing.
func (b BankTransferProcessor) ProcessPayment(amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}

	fmt.Printf("Processing bank transfer of $%.2f via %s\n", amount, b.BankCode)
	fmt.Printf("Transfer initiated for %s account\n", b.AccountType)

	return nil
}

// GetTransactionFee calculates bank transfer fees.
func (b BankTransferProcessor) GetTransactionFee(amount float64) float64 {
	// bank transfers usually have lower fees (flat $1.50)
	return 1.50
}

func ProcessMultiplePayments(processors []PaymentProcessor, amounts []float64) {
	fmt.Println("\n=== Processing Multiple Payments ===")

	for i, processor := range processors {
		if i < len(amounts) {
			amount := amounts[i]
			fee := processor.GetTransactionFee(amount)
			total := amount + fee

			fmt.Printf("\nPayment #%d:\n", i+1)
			fmt.Printf("Amount: $%.2f, Fee: $%.2f, Total: $%.2f\n", amount, fee, total)

			if err := processor.ProcessPayment(amount); err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		}
	}
}

func demoBasicInterface() {
	// create different payment processors
	creditCard := CreditCardProcessor{
		CardNumber: "4532123456789012",
		Provider:   "Visa",
	}

	bankTransfer := BankTransferProcessor{
		BankCode:    "BCA",
		AccountType: "Savings",
	}

	// demonstrate individual usage
	fmt.Println("=== Individual Payment Processing ===")

	err := creditCard.ProcessPayment(100.00)
	if err != nil {
		fmt.Printf("Credit card error: %v\n", err)
	}

	err = bankTransfer.ProcessPayment(250.00)
	if err != nil {
		fmt.Printf("Bank transfer error: %v\n", err)
	}

	// demonstrate polymorphism - both types can be treated as PaymentProcessor
	processors := []PaymentProcessor{creditCard, bankTransfer}
	amounts := []float64{75.00, 150.00}

	ProcessMultiplePayments(processors, amounts)

	// show how interfaces enable flexible code
	fmt.Println("\n=== Flexible Payment Selection ===")
	var selectedProcessor PaymentProcessor

	// in a real app, this choice might come from user input or configuration
	paymentMethod := "credit_card" // or "bank_transfer"

	switch paymentMethod {
	case "credit_card":
		selectedProcessor = creditCard
	case "bank_transfer":
		selectedProcessor = bankTransfer
	default:
		fmt.Println("Unknown payment method")
		return
	}

	// the same code works regardless of which concrete type we choose
	fee := selectedProcessor.GetTransactionFee(200.00)
	fmt.Printf("Selected payment method fee for $200.00: $%.2f\n", fee)
	selectedProcessor.ProcessPayment(200.00)
}
