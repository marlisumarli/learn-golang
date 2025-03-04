package main

import "fmt"

func main() {
	var name [3]string

	name[0] = "A"
	name[1] = "B"
	name[2] = "C"

	fmt.Println(name)

	// Array panjang tetap, sama seperti diatas
	var number = [3]int{
		1,
		2,
		3,
	}

	fmt.Println(number)

	// Array literal, bebas mau berapapun
	var numberLtrl = [...]int{
		1,
		2,
		3,
		4,
		5,
	}

	fmt.Println(numberLtrl)
	fmt.Println(len(numberLtrl))
}
