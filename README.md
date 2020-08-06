# GolangTest

 Golang General Coding Test for the JazzNetworks Cloud Engineer Position for Premdaas

![GitHub Workflow Status (branch)](https://img.shields.io/github/workflow/status/JazzNetCodingTest/GolangTest/Go/master) 
[![codecov](https://codecov.io/gh/JazzNetCodingTest/GolangTest/branch/master/graph/badge.svg)](https://codecov.io/gh/JazzNetCodingTest/GolangTest)

# Question 1
##### Write a program in a language of your choice to calculate the sum of the first 100 even-valued Fibonacci numbers
Even fibonacci numbers can be obtained by F(n) = 4*f(n-1)+f(n-2) , with starting seeds of 0 and 2. Implementation can be found at [fibonacci.go](pkg/fibonacci.go) and tests are at [fibonacci_test.go](pkg/fibonacci_test.go)

# Question 2
##### Write a function in a language of your choice which, when passed two sorted arrays of integers returns an array of any numbers which appear in both. No value should appear in the returned array more than once. 
One of the Array of integers is converted in to Hashmap, so each element in other array can be looked up easily. Implementation can be found at [intersection.go](pkg/intersection.go) and tests are at [intersection_test.go](pkg/intersection_test.go)


# Question 3
##### Write a function in a language of your choice which, when passed a positive integer returns true if the decimal representation of that integer contains no odd digits and otherwise returns false.
Implementation using simple string.containsany search can be found at [odd_digit_finder.go](pkg/odd_digit_finder.go) and tests are at [odd_digit_finder_test.go](pkg/odd_digit_finder_test.go)

# Question 4
##### Write a function in a language of your choice which, when passed a decimal digit X, returns the value of X+XX+XXX+XXXX. E.g. if the supplied digit is 3 it should return 3702 (3+33+333+3333)

a Simp,le flexible implementation using loop can be found at [digits_sum.go](pkg/digits_sum.go) and tests are at [odd_digit_finder_test.go](pkg/digits_sum_test.go)