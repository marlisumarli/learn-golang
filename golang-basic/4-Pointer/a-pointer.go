package main

import "fmt"

type Address struct {
	City, Country string
}

func main() {
	//i := Address{"Karawang", "Indonesia"}
	//p := i // Copy
	//
	//p.City = "Jakarta" // Ga bakal merubah i
	//
	//fmt.Printf("i  %T: &i=%p i=%v\n", i, &i, i)
	//fmt.Printf("p  %T: &p=%p p=%v\n", p, &p, p)

	i := Address{"Karawang", "Indonesia"}
	p := &i // Pointer ke i

	fmt.Println(i)

	p.City = "Jakarta"

	fmt.Printf("i -- %T: &i=%p i=%v\n", i, &i, i)
	fmt.Printf("p -- %T: &p=%p p=&i=%p  *p=i=%v\n", p, &p, p, *p) // Tidak menggunakan & pada %p karena ini pointer ke i
}
