package main

import (
	"fmt"
	"strconv"
)

func GenDisplaceFn(a, vo, so float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*t*t + vo*t + so
	}
}

func main() {
	var a, vo, so float64

	fmt.Print("Enter acceleration (a): ")
	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		return
	}

	fmt.Print("Enter initial velocity (vo): ")
	_, err = fmt.Scan(&vo)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		return
	}

	fmt.Print("Enter initial displacement (so): ")
	_, err = fmt.Scan(&so)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		return
	}

	fn := GenDisplaceFn(a, vo, so)

	fmt.Print("Enter time (t): ")
	var t float64
	_, err = fmt.Scan(&t)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		return
	}

	fmt.Println("Displacement after", strconv.FormatFloat(t, 'f', 2, 64), "seconds:", fn(t))
}
