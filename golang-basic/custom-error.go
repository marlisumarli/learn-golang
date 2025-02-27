package main

import (
	"fmt"
)

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func saveData(id string, data any) error {
	if id == "" {
		return &validationError{"Validation Error"}
	} else if id != "golang" {
		return &notFoundError{"Not Found Error"}
	} else {
		return nil
	}
}

func main() {
	err := saveData("", nil)

	// Contoh menggunakan if
	if err != nil {
		// if valErr, ok := err.(*validationError); ok {
		// 	fmt.Println(valErr.Message)
		// } else if err404, ok := err.(*notFoundError); ok {
		// 	fmt.Println(err404)
		// } else {
		// 	fmt.Println("Unknown", err.Error())
		// }

		switch finalError := err.(type) {
		case *validationError:
			fmt.Println(finalError.Error())
		case *notFoundError:
			fmt.Println(finalError.Error())
		default:
			fmt.Println("Unknown", finalError.Error())
		}
	} else {
		fmt.Println("Sukses")
	}
}
