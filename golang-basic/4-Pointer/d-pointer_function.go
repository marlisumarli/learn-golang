package main

import "fmt"

type User struct {
	Username, Password string
}

// Tidak berubah
//func changePassword(a User) {
//	a.Password = "golang09"
//}

func changePassword(a *User) {
	a.Password = "golang09"
}

func main() {
	user := User{}
	//changePassword(user)
	changePassword(&user)

	fmt.Println(user)
}
