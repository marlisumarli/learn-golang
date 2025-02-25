package main

import "fmt"

func main() {
	ctrl := ""

	switch ctrl {
	case "w":
		fmt.Println("Move forward")
	case "a":
		fmt.Println("Move backward")
	case "s":
		fmt.Println("Tilt to the left")
	case "d":
		fmt.Println("Tilt to the right")
	default:
		fmt.Println("Motionless")
	}
}
