package fizzbuzz

import (
	"fmt"
)

// FizzBuzz function returns "Fizz" for multiples of 3,
// "Buzz" for multiples of 5, "FizzBuzz" for multiples of both 3 and 5,
// and the number itself for all other cases.
func FizzBuzz(n int) string {
	switch {
	case n%15 == 0:
		return "FizzBuzz"
	case n%3 == 0:
		return "Fizz"
	case n%5 == 0:
		return "Buzz"
	default:
		return fmt.Sprintf("%d", n)
	}
}
