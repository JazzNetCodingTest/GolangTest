package main

import (
	"math"
)

/*
* Using loop to have more flexile implementation than manual addition for each level
 */
func sumOfdigits(digit int) int {
	sum := digit
	result := 0
	for i := 1; i <= 3; i++ {
		result = digit*int(math.Pow10(i)) + result
		sum = sum + result + digit
	}
	return sum
}
