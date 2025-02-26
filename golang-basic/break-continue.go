package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("Angka genap:", i)
			continue
		}

		fmt.Println("Angka ganjil", i)
	}
}
