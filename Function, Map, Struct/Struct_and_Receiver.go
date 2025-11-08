package main

import "fmt"

type BankAccount struct { // Struct definition
	owner   string
	balance float64
}

func (b *BankAccount) Diposit(amount float64) { // Receiver function with pointer receiver
	b.balance += amount
}

func (b *BankAccount) showBalance() { // Receiver function to show balance	
	fmt.Printf("Account owner: %s, Balance: $%.2f\n", b.owner, b.balance)
}

func main() { 
	// Start your code....
	account := BankAccount{owner: "Mojahid", balance: 1000.00} // Struct initialization
	account.Diposit(500.00)
	account.showBalance()
}
