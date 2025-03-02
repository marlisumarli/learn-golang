package main

import "fmt"

type Address struct {
	City, Country string
}

func main() {
	//i := Address{"Subang", "Indonesia"}
	//p := i // Copy
	//
	//p.City = "Jakarta"
	//
	//fmt.Printf("i  %T: &i=%p i=%v\n", i, &i, i)
	//fmt.Printf("p  %T: &p=%p p=%v\n", p, &p, p)

	i := Address{"Subang", "Indonesia"}
	p := &i // Pointer ke i

	p.City = "Jakarta"

	fmt.Printf("i -- %T: &i=%p i=%v\n", i, &i, i)
	fmt.Printf("p -- %T: &p=%p p=&i=%p  *p=i=%v\n", p, &p, p, *p) // Tidak menggunakan & pada %p karena ini pointer ke i
}
