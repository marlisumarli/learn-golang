/*
Quiz 1:
- Buat penghitungan gaji karyawan perjam
- Hitung angka ganjil dan genap
*/
package main

import "fmt"

func main() {
	const subject = "Quiz 1"
	type MyKaryawan string

	fmt.Println(subject)

	fmt.Println("Perhitungan Gaji Perjam")
	var emp1 MyKaryawan = "Amba"
	var gaji = 10000
	var hour = 0

	fmt.Println("Karyawan: ", emp1)
	hour++
	hour++
	hour++
	fmt.Println("Gaji perjam: ", gaji)
	fmt.Println("Jam kerja: ", hour)
	fmt.Println("Stat gaji: ", gaji*hour)

	fmt.Println("Angka Ganjil dan Genap")
	var angka = 3

	var genap = (angka % 2) == 0
	var ganjil = (angka % 2) == 1

	fmt.Println("Angka", angka, "?")
	fmt.Println("Genap", genap)
	fmt.Println("Ganjil", ganjil)
}
