package main

import (
	"errors"
	"fmt"
)

func Pembagian(n, divided int) (int, error) {
	if divided == 0 {
		return 0, errors.New("Pembagian Dengan NOL")
	} else {
		return n / divided, nil
	}
}

func main() {
	res, err := Pembagian(4, 1)

	if err == nil {
		fmt.Println("Hasil", res)
	} else {
		fmt.Println("Error", err.Error())
	}
}
