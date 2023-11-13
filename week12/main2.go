package main

import "fmt"

func main() {
	// a := []string{"a", "b", "c", "d"}
	// aS := a[:2]
	// aS[1] = "Z"
	// fmt.Println(a, aS)

	// b := [4]int{4, 3, 2, 1}
	// bS := b[1:3]
	// fmt.Println(b, bS)

	a := make([]string, 4, 5) // type, length, capacity
	a[0] = "a"
	a[1] = "b"
	a[2] = "c"
	aS := a[0:2]
	aS[1] = "Z"
	// c := append(a, "y")
	c := append(a, "y", "x")

	c[0] = "Q"
	fmt.Println(a, len(a), cap(a))
	fmt.Println(c, len(c), cap(c))
	fmt.Printf("%x %x %x\n", &a[0], &aS[0], &c[0])
}