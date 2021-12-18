package lazy_lists

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNumberGenerator(t *testing.T) {

	t.Run("10 numbers", func(t *testing.T) {
		var g Generator = func(idx int) Element { return Int(idx) }
		listOfNumbers := NewLazyList(g)

		for i := 0; i < 10; i++ {
			var num Element
			num, listOfNumbers = listOfNumbers()
			require.Equal(t, num, Int(i))
		}
	})

	t.Run("10 alphabet list", func(t *testing.T) {
		var g Generator = func(idx int) Element {
			if idx < 26 {
				return String(fmt.Sprintf("%c", 'a'+idx))
			} else if idx == 26 {
				return nil
			}

			panic(1) // should never reach this point
		}
		list := NewLazyList(g)

		for _, alph := range []string{"a", "b", "c", "d", "e"} {
			var a Element
			a, list = list()
			require.Equal(t, String(alph), a)
		}
	})

	t.Run("alphabet loop using map", func(t *testing.T) {
		list := Numbers.Fmap(func(in Element) Element {
			num, ok := in.(Int)
			if !ok {
				return nil
			}

			return String(fmt.Sprintf("%c", 'a'+num%26))
		})

		for _, alph := range []string{"a", "b", "c", "d", "e"} {
			var a Element
			a, list = list()
			require.Equal(t, String(alph), a)
		}
	})

	t.Run("inefficient prime number list", func(t *testing.T) {
		numbers := NumbersFrom(2) // take [2, 3, 4, ...]

		removePrime := func(_primeNumber Element) Filter {
			primeNumber := int(_primeNumber.(Int))
			return func(in Element) bool {
				num, ok := in.(Int)
				if !ok {
					panic(1)
				}

				return int(num)%int(primeNumber) != 0
			}
		}

		for _, expectedPrime := range []int{2, 3, 5, 7, 11, 13, 17} {
			var prime Element
			prime, rest := numbers()
			numbers = rest.Filter(removePrime(prime))
			require.Equal(t, Int(expectedPrime), prime)
		}
	})

	t.Run("single filter prime number list", func(t *testing.T) {
		numbers := NumbersFrom(2) // take [2, 3, 4, 5,...]

		knownPrimes := []int{}
		var removePrimes Filter = func(in Element) bool {
			num, ok := in.(Int)
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

		for _, expectedPrime := range []int{2, 3, 5, 7, 11, 13, 17} {
			var prime Element
			prime, primeList = primeList()
			require.Equal(t, Int(expectedPrime), prime)

			knownPrimes = append(knownPrimes, int(prime.(Int)))

		}
	})

}
