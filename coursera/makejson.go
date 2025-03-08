package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	userData := make(map[string]string)

	fmt.Print("Name: ")
	var name string
	fmt.Scan(&name)

	fmt.Print("Address: ")
	var address string
	fmt.Scan(&address)

	userData["name"] = name
	userData["address"] = address

	jsonData, err := json.Marshal(userData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("JSON Output: ")
	fmt.Println(string(jsonData))
}
