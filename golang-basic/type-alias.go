package main

import "fmt"

func main() {
	type MyString string
	var myPhone MyString = "13131212"

	var ASCII = 8383

	var convert = MyString(ASCII)

	fmt.Println(myPhone)
	fmt.Println(convert)
}
