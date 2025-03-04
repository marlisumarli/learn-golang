package main

import "fmt"

func main() {
	months := [...]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

	// Membuat slice dari array months
	slice1 := months[4:6] // Mengambil elemen index 4 sampai 5 (May, June)
	slice2 := months[:3]  // Mengambil dari awal sampai index 2 (January, February, March)
	slice3 := months[3:]  // Mengambil dari index 3 sampai akhir
	slice4 := months[:]   // Mengambil seluruh elemen

	fmt.Println("slice1:", slice1)
	fmt.Println("slice2:", slice2)
	fmt.Println("slice3:", slice3)
	fmt.Println("slice4:", slice4)

	// Array hari
	days := [...]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	// Membuat slice dari array days
	daySlice1 := days[5:] // Mengambil elemen index 5 hingga akhir (Saturday, Sunday)
	fmt.Println("daySlice1 before modification:", daySlice1)

	// Mengubah nilai dalam slice (akan mempengaruhi array asli)
	daySlice1[0] = "Sabtu"
	daySlice1[1] = "Minggu"
	fmt.Println("daySlice1 after modification:", daySlice1)
	fmt.Println("days array after modification:", days)

	// Menambahkan elemen ke slice dengan append (membuat slice baru jika kapasitas penuh)
	daySlice2 := append(daySlice1, "Minggu Baru")
	daySlice2[0] = "Sabtu Lama" // Tidak mengubah days karena slice baru dibuat

	fmt.Println("daySlice1:", daySlice1)
	fmt.Println("daySlice2:", daySlice2)
	fmt.Println("days array:", days)

	// Membuat slice baru dengan make()
	newSlice := make([]string, 2, 5) // Slice dengan panjang 2 dan kapasitas 5
	newSlice[0] = "Belajar"
	newSlice[1] = "Go"
	fmt.Println("newSlice:", newSlice)
	fmt.Println("Length:", len(newSlice))
	fmt.Println("Capacity:", cap(newSlice))

	// Menambahkan elemen menggunakan append
	newSlice2 := append(newSlice, "Golang")
	fmt.Println("newSlice2:", newSlice2)
	fmt.Println("Length:", len(newSlice2))
	fmt.Println("Capacity:", cap(newSlice2))

	// Mengubah nilai dalam newSlice2 tidak mempengaruhi newSlice
	newSlice2[0] = "Google"
	fmt.Println("newSlice2 after modification:", newSlice2)
	fmt.Println("newSlice remains unchanged:", newSlice)

	// Menyalin slice ke slice lain
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)

	fmt.Println("fromSlice:", fromSlice)
	fmt.Println("toSlice:", toSlice)

	// Perbedaan array dan slice
	thisArray := [...]int{1, 2, 3, 4, 5} // Array dengan panjang tetap
	thisSlice := []int{1, 2, 3, 4, 5}    // Slice dengan panjang dinamis
	fmt.Println("Array:", thisArray)
	fmt.Println("Slice:", thisSlice)
}
