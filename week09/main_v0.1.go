package main

import "fmt"

func main() {

	a := 10
	var pa *int
	pa = &a

	fmt.println(a, *pa)
	fmt.println(a, pa)
	fmt.Printf("%T %T %T %T \n", a, *pa, &a, pa)
	fmt.println(&pa)
}