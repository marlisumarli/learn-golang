package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Philosopher struct {
	id             int
	leftChopstick  *sync.Mutex
	rightChopstick *sync.Mutex
	eatCount       int
}

type Host struct {
	allowEating chan struct{}
}

func (p *Philosopher) dine(host *Host, wg *sync.WaitGroup) {
	defer wg.Done()

	for p.eatCount < 3 {
		host.allowEating <- struct{}{}

		if rand.Intn(2) == 0 {
			p.leftChopstick.Lock()
			p.rightChopstick.Lock()
		} else {
			p.rightChopstick.Lock()
			p.leftChopstick.Lock()
		}

		fmt.Println("starting to eat", p.id)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)+50))
		fmt.Println("finishing eating", p.id)
		p.eatCount++

		p.leftChopstick.Unlock()
		p.rightChopstick.Unlock()

		<-host.allowEating

		time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)+50))
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	chopsticks := make([]*sync.Mutex, 5)
	for i := 0; i < 5; i++ {
		chopsticks[i] = &sync.Mutex{}
	}

	host := &Host{allowEating: make(chan struct{}, 2)}

	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{
			id:             i + 1,
			leftChopstick:  chopsticks[i],
			rightChopstick: chopsticks[(i+1)%5],
			eatCount:       0,
		}
	}

	var wg sync.WaitGroup
	for _, p := range philosophers {
		wg.Add(1)
		go p.dine(host, &wg)
	}

	wg.Wait()
	fmt.Println("All philosophers have finished eating!")
}
