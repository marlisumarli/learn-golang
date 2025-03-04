package main

import "fmt"

type Wallet struct {
	User    string
	Balance float64
}

func main() {
	i := new(Wallet)
	p := i // Pointer ke i

	p.User = "zero"
	p.Balance = 20_000.21

	fmt.Printf("i -- %T: &i=%p i=&struct=%v\n", i, i, i)          // Kalau pakai new, i bukan jadi variable melainkan pointer
	fmt.Printf("p -- %T: &p=%p p=&i=%p  *p=i=%v\n", p, &p, p, *p) // ini hanya menyalin i yg dimana i adalah pointer
}
