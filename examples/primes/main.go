package main

import (
	"fmt"

	ll "github.com/teja2010/lazy_lists"
)

func main() {
	fmt.Println("example to print primes")

	numbers := ll.NumbersFrom(2) // take [2, 3, 4, 5,...]

	knownPrimes := []int{}
	var removePrimes ll.Filter = func(in ll.Element) bool {
		num, ok := in.(ll.Int)
		if !ok {
			panic(1)
		}

		for _, primes := range knownPrimes {
			if int(num)%primes == 0 {
				return false
			}
		}
		return true
	}

	primeList := numbers.Filter(removePrimes)

	for {
		var prime ll.Element
		prime, primeList = primeList()
		if prime == nil {
			fmt.Printf("end of primesList") // should never happen
			return
		}
		knownPrimes = append(knownPrimes, int(prime.(ll.Int)))
		fmt.Printf("%d, ", prime)
	}
}
