package main

import "fmt"

type Profile struct {
	Name string
}

type ProfileInt interface{}

func main() {
	var p1, p2 ProfileInt = Profile{"iswbm"}, Profile{"iswbm"}
	var p3, p4 ProfileInt = &Profile{"iswbm"}, &Profile{"iswbm"}

	fmt.Printf("p1 --> type: %T, data: %v \n", p1, p1)
	fmt.Printf("p2 --> type: %T, data: %v \n", p2, p2)
	fmt.Println(p1 == p2) // true

	fmt.Printf("p3 --> type: %T, data: %p \n", p3, p3)
	fmt.Printf("p4 --> type: %T, data: %p \n", p4, p4)
	fmt.Println(p3 == p4) // false
}
