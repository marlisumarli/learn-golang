package main

import "fmt"

func main() {
	month := [...]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "Desember"}

	slice1 := month[4:6]
	fmt.Println("slice1", slice1)

	slice2 := month[:3]
	fmt.Println("slice2", slice2)

	slice3 := month[3:]
	fmt.Println("slice3", slice3)

	slice4 := month[:]
	fmt.Println("slice4", slice4)

	days := [...]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	daySlice1 := days[5:]

	fmt.Println(daySlice1)

	daySlice1[0] = "Sabtu"
	daySlice1[1] = "Minggu"
	// days[0] = "Senin"

	fmt.Println(daySlice1)
	fmt.Println(days)

	daySlice2 := append(daySlice1, "Minggu Baru") // Append slice 1 ke new slice
	daySlice2[0] = "Sabtu Lama"                   // Tidak mengubah parent days, instead of silce 1

	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)

	var newSlice = make([]string, 2, 5)
	newSlice[0] = "Belajar"
	newSlice[1] = "Go"

	// newSlice[2] = "Go" // Error, harus menggunakan append

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Golang")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "Google"

	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	thisArray := [...]int{1, 2, 3, 4, 5}
	thisSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(thisArray)
	fmt.Println(thisSlice)
}
