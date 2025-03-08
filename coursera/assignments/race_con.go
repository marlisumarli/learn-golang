// Go Routine: Withdraw and Deposit of a BankAccount
// Race Condition: Both routines share the same data "BankAccount.Balance".
// Without synchronization, the balance update may conflict, depending on when a deposit and withdrawn is done
// and at what point in time the Balance is checked. The outcome can vary.

package main

import (
	"fmt"
	"time"
)

type BankAccount struct {
	Balance *float64
}

func (ba *BankAccount) Withdraw(amount float64) {
	*ba.Balance -= amount
	time.Sleep(60 * time.Millisecond)
	fmt.Println("withdraw money", amount, "new balance", *ba.Balance)
}
func (ba *BankAccount) Deposit(amount float64) {
	*ba.Balance += amount
	time.Sleep(100 * time.Millisecond)
	fmt.Println("deposit money", amount, "new balance", *ba.Balance)
}

func main() {
	var amount = 0.0
	ba := BankAccount{
		Balance: &amount,
	}
	fmt.Printf("Balance is: %d \n", int(*ba.Balance))

	for i := 0; i < 10; i++ {
		fmt.Println("Transaction: Deposit ", 100, ", new Balance is: ", *ba.Balance)
		go ba.Deposit(100)
		fmt.Println("Transaction: Withdraw  ", 100, ", new Balance is: ", *ba.Balance)
		go ba.Withdraw(100)
	}
	fmt.Printf("Check Balance : %d \n", int(*ba.Balance))
}
