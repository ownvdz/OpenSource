package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[1:]) // Arges <--- Arguments
	fmt.Println(os.Args[2])
}
