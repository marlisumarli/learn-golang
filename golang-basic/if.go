package main

import "fmt"

func main() {
	kecepatan := 1

	if kecepatan > 1 && kecepatan <= 5 {
		fmt.Println("Anda sedang berjalan 🚶🏻‍♂️")
	} else if kecepatan > 5 && kecepatan <= 15 {
		fmt.Println("Anda sedang naik sepeda Ontel 🚲")
	} else if kecepatan > 15 {
		fmt.Println("Anda sedang berkendara 🏍️")
	} else {
		fmt.Println("Sepertinya anda menaiki Kura-kura 😁")
	}

	// Short IF
	a := "aa"

	if length := len(a); length > 2 {
		fmt.Println("Berisik! Jangan teriakk!!")
	} else {
		fmt.Println("Good 👍")
	}
}
