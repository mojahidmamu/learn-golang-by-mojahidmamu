package main

import "fmt"

type BankAccount struct { // Struct definition
	owner   string
	balance float64
}

func (b *BankAccount) Diposit(amount float64) {
	b.balance += amount
}

func (b *BankAccount) showBalance() {
	fmt.Printf("Account owner: %s, Balance: $%.2f\n", b.owner, b.balance)
}

func main() {
	// Start your code....
	account := BankAccount{owner: "Mojahid", balance: 1000.00}
	account.Diposit(500.00)
	account.showBalance()
}
