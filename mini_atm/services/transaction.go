package services

import "fmt"

// Account represents a bank account with balance information
type Account struct {
	Saldo int // Saldo is the current account balance in Rupiah
}

// Withdraw deducts the specified amount from the account balance
// Returns the remaining balance after withdrawal, or error if validation fails
func (a *Account) Withdraw(amount int) (int, error) {
	// Validate amount is positive
	if amount <= 0 {
		return 0, fmt.Errorf("nominal harus lebih besar dari 0")
	}
	
	// Validate sufficient balance
	if amount > a.Saldo {
		return 0, fmt.Errorf("saldo saat ini tidak mencukupi: Rp.%d", a.Saldo)
	}

	// Deduct amount from balance
	a.Saldo -= amount
	return a.Saldo, nil
}

// Deposit adds the specified amount to the account balance
// Returns the new balance after deposit, or error if validation fails
func (a *Account) Deposit(amount int) (int, error) {
	// Validate amount is positive
	if amount <= 0 {
		return 0, fmt.Errorf("nominal harus lebih besar dari 0")
	}
	
	// Add amount to balance
	a.Saldo += amount
	return a.Saldo, nil
}

// Transfer sends the specified amount to another account
// Returns the remaining balance and destination account number, or error if validation fails
func (a *Account) Transfer(amount int, destinationAccount string) (int, string, error) {
	// Validate amount is positive
	if amount <= 0 {
		return 0, "", fmt.Errorf("nominal harus lebih besar dari 0")
	}
	
	// Validate sufficient balance
	if amount > a.Saldo {
		return 0, "", fmt.Errorf("saldo saat ini tidak mencukupi: Rp.%d", a.Saldo)
	}

	// Deduct amount from balance for transfer
	a.Saldo -= amount
	return a.Saldo, destinationAccount, nil
}
