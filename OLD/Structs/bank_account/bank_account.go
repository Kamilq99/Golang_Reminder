package main

import "fmt"

type BankAccount struct {
	Owner_name    string
	Owner_surname string
	Balance       float64
}

func (deposit *BankAccount) Deposit(amount float64) {

	deposit.Balance += amount
}

func (withdraw *BankAccount) Withdraw(amount float64) {
	if withdraw.Balance <= 0 {
		fmt.Print("You don't have enough money")
	} else if withdraw.Balance < amount {
		fmt.Print("You don't have enough money")
	} else {
		withdraw.Balance -= amount

		fmt.Println(withdraw.Balance)
	}
}

func main() {

	account_Jan := BankAccount{
		Owner_name:    "Jan",
		Owner_surname: "Kowalski",
		Balance:       1000,
	}

	fmt.Println(account_Jan.Balance)

	account_Jan.Deposit(500)

	fmt.Println(account_Jan.Balance)

	account_Jan.Withdraw(1500)
}
