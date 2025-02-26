package main

import "fmt"

type Discount func(price float64) float64

// Function as parameter
func countTotal(price float64, discount Discount) float64 {
	fmt.Println("Harga produk:", price)
	return discount(price)
}

// Discount function
func discount10(price float64) float64 {
	fmt.Println("Diskon 10%")
	return price * .90
}

func discount50k(price float64) float64 {
	fmt.Println("Potongan 50k")
	if price >= 200_000 {
		return price - 50_000
	}

	return price
}

func main() {
	product := 300_000.0

	total1 := countTotal(product, discount10)
	fmt.Println("Total setelah diskon:", total1)

	total2 := countTotal(product, discount50k)
	fmt.Println("Total setelah potongan:", total2)

	total3 := countTotal(product, func(price float64) float64 {
		fmt.Println("Diskon 20%")
		return price * .80
	})
	fmt.Println("Total setelah diskon:", total3)
}
