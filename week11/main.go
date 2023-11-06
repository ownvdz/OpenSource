package main

import "fmt"

func main() {
	// var primes [3]int
	// primes[0] = 2
	// primes[1] = 3
	// primes[2] = 5
	// fmt.Println(primes, primes[1])

	// var primes [3]int = [3]int{2,3,5}
	// fmt.Println(primes, primes[1])

	primes := [3]int{2,3,5}
	fmt.Println(primes, primes[1])

	test := [5]bool{true, true, true}	
	fmt.Println(test[3])

	// for i := 0; i < 3; i++ {
	// 	fmt.Println(primes[i])
	// }

	// i := 0
	// for i < 3 {
	// 	fmt.Println(primes[i])
	// 	i++
	// }

	// for idx, prime := range primes {
	// 	fmt.Println(idx, prime)
	// }

	// for _, prime := range primes {
	// 	fmt.Println(prime)
	// }

	sum := 0
	for _, prime := range primes {
		sum = sum + prime
	}
	fmt.Println(sum)
	fmt.Println(float64(sum)/float64(len(primes)))
	fmt.Printf("%.2f\n", float64(sum)/float64(len(primes)))
}