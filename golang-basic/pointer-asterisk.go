package main

import "fmt"

type Address struct {
	City, Country string
}

func main() {
	i := Address{"Subang", "Indonesia"}
	p := &i // Pointer ke i

	p.City = "Jakarta"

	//p = &Address{"USA", "America"}                                // Ini sama seperti copy, tidak mengacu pada i
	//fmt.Printf("p -- %T: &p=%p p=&i=%p  *p=i=%v\n", p, &p, p, *p) // Tidak menggunakan & pada %p karena ini pointer ke i

	fmt.Printf("i -- %T: &i=%p i=%v\n", i, &i, i)

	*p = Address{"USA", "America"}
	fmt.Printf("p -- %T: &p=%p p=&i=%p  *p=i=%v\n", p, &p, p, *p) // Tidak menggunakan & pada %p karena ini pointer ke i

	fmt.Printf("i -- %T: &i=%p i=%v", i, &i, i)
}
