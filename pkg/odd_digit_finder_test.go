package main

import (
	"testing"
)

func TestWithEvenDigits(t *testing.T) {
	got := noOdddigits(2468)
	want := true
	assertBoolEqual(got, want, t)
}

func TestWithOddDigits(t *testing.T) {
	got := noOdddigits(324868421)
	want := false
	assertBoolEqual(got, want, t)
}

func assertBoolEqual(got bool, want bool, t *testing.T) {
	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}
