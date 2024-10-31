package main

import (
	"fmt"
	"math"
)

func checkIfPrime(num int) bool {
	result := true
	sqRoot := int(math.Sqrt(float64(num)))
	for i := 3; i <= sqRoot; i += 2 {
		if num%i == 0 {
			result = false
			break
		}
	}
	return result
}

func printPrimes(max int) {
	var primes []int
	for i := 2; i <= max; i++ {
		if i == 2 {
			fmt.Println(i)
			primes = append(primes, i)
		} else if i%2 == 0 {
			continue
		} else {
			if checkIfPrime(i) {
				fmt.Printf("%v\n", i)
			}
		}
	}
	fmt.Println(primes)
}

func main() {
	printPrimes(78)
}
