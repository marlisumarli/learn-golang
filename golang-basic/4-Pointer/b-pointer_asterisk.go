package main

import "fmt"

type Fruit struct {
	Name, Color string
}

func main() {
	i := Fruit{"Apple", "Red"}
	p := &i // Pointer ke i

	p.Name = "Strawberry"

	////Ini sama seperti copy, tidak mengacu pada i
	//p = &Fruit{"Apple", "Green"}
	//// Tidak menggunakan & pada %p karena ini pointer ke i
	//fmt.Printf("p -- %T: &p=%p p=&i=%p  *p=i=%v\n", p, &p, p, *p)

	fmt.Printf("i -- %T: &i=%p i=%v\n", i, &i, i)

	// Yang menggunakan address bakal ganti semua
	*p = Fruit{"Orange", "Yellow"}
	// Tidak menggunakan & pada %p karena ini pointer ke i
	fmt.Printf("p -- %T: &p=%p p=&i=%p  *p=i=%v\n", p, &p, p, *p)

	fmt.Printf("i -- %T: &i=%p i=%v", i, &i, i)
}
