package main

import (
	"fmt"
	"sync"
	"time"
)

type BankAccount struct {
	Balance float64
	mu      sync.Mutex
}

func (ba *BankAccount) Deposit(amount float64) {
	ba.mu.Lock()
	ba.Balance += amount
	ba.mu.Unlock()
	time.Sleep(50 * time.Millisecond)
	fmt.Println("Deposited:", amount, "New Balance:", ba.Balance)
}

func (ba *BankAccount) Withdraw(amount float64) {
	ba.mu.Lock()
	ba.Balance -= amount
	ba.mu.Unlock()
	time.Sleep(50 * time.Millisecond)
	fmt.Println("Withdrew:", amount, "New Balance:", ba.Balance)
}

func main() {
	account := &BankAccount{Balance: 1000}
	fmt.Println("Starting Balance:", account.Balance)

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(2)
		go func() {
			account.Deposit(200)
			wg.Done()
		}()
		go func() {
			account.Withdraw(150)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Final Balance:", account.Balance)
}
