package main

import (
	"fmt"
)

// GenDisplaceFn returns a function to compute displacement.
func GenDisplaceFn(a, vo, so float64) func(t float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*t*t + vo*t + so
	}
}

func main() {
	var a, vo, so, t float64

	// Prompt user for input values
	fmt.Print("Enter acceleration: ")
	fmt.Scan(&a)
	fmt.Print("Enter initial velocity: ")
	fmt.Scan(&vo)
	fmt.Print("Enter initial displacement: ")
	fmt.Scan(&so)

	// Generate the displacement function
	fn := GenDisplaceFn(a, vo, so)

	// Prompt user for time and compute displacement
	fmt.Print("Enter time: ")
	fmt.Scan(&t)
	displacement := fn(t)
	fmt.Printf("The displacement after %.2f seconds is: %.2f\n", t, displacement)
}
