package main

import (
	"strings"
	"strconv"
)

/*
* sticking with simple string.containsany to find the odd digist since we have only 5 odd digits
*/
func noOdddigits(input int) (a bool) {
	inputs := strconv.Itoa(input)
	odddigits := "13579"
	return !strings.ContainsAny(inputs,odddigits)
}