package main

import "fmt"

type Address struct {
	City, Country string
}

func main() {
	i := new(Address)
	p := i // Pointer ke i

	p.City = "Jakarta"

	fmt.Printf("i -- %T: &i=%p i=&struct=%v\n", i, i, i)          // Kalau pakai new, i bukan jadi variable melainkan pointer
	fmt.Printf("p -- %T: &p=%p p=&i=%p  *p=i=%v\n", p, &p, p, *p) // ini hanya menyalin i yg dimana i adalah pointer
}
