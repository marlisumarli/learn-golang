/*
A race condition occurs when two or more goroutines access the same shared variable concurrently,
and at least one of them writes to it, leading to unpredictable behavior.
*/
package main

import (
	"fmt"
	"time"
)

type BankAccount struct {
	Balance float64
}

func (ba *BankAccount) Deposit(amount float64) {
	ba.Balance += amount
	time.Sleep(50 * time.Millisecond)
	fmt.Println("Deposited:", amount, "New Balance:", ba.Balance)
}

func (ba *BankAccount) Withdraw(amount float64) {
	ba.Balance -= amount
	time.Sleep(50 * time.Millisecond)
	fmt.Println("Withdrew:", amount, "New Balance:", ba.Balance)
}

func main() {
	account := &BankAccount{Balance: 1000}

	fmt.Println("Starting Balance:", account.Balance)

	for i := 0; i < 5; i++ {
		go account.Deposit(200)
		go account.Withdraw(150)
	}

	time.Sleep(1 * time.Second)

	fmt.Println("Final Balance:", account.Balance)
}
