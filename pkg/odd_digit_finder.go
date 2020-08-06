package main

import (
	"strings"
	"strconv"
)


func noOdddigits(input int) (a bool) {
	inputs := strconv.Itoa(input)
	odddigits := "13579"
	return !strings.ContainsAny(inputs,odddigits)
}