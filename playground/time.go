package main

import (
	"fmt"
	"time"
)

func sayHello() {
	time.Sleep(1 * time.Second)
	fmt.Println("Hello")
}

func main() {
	go sayHello()
	fmt.Println("Running")
	time.Sleep(2 * time.Second)
}
