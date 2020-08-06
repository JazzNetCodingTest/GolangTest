package main

import (
	"testing"
)


func TestWith3(t *testing.T) {
	got := sumOfdigits(3)
	want := 3702
	assertIntEqual(got, want, t)
}

func TestWith4(t *testing.T) {
	got := sumOfdigits(4)
	want := 4936
	assertIntEqual(got, want, t)
}

func TestWith9(t *testing.T) {
	got := sumOfdigits(9)
	want := 11106
	assertIntEqual(got, want, t)
}


func assertIntEqual(got int ,want int , t *testing.T){
	if (got != want) {
		t.Errorf("got %v want %v", got, want)
	}
}