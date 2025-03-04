package main

import "fmt"

func log() {
	fmt.Println("End app")
	message := recover()
	fmt.Println("Terjadi error: ", message)
}

func run(error bool) {
	defer log()

	if error {
		panic("IF statement di-eksekusi")
	}
}

func main() {
	run(true)
	fmt.Println("Terus berjalan")
}
