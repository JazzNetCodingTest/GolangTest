package main

import (
	"math/big"
	"testing"
)

func TestFibonacciSumForZero(t *testing.T) {
	got := getEvenFibonacciSum(0)
	want := big.NewInt(0)
	assertBigIntEqual(got,want,t)
}


func TestFibonacciSumForOne(t *testing.T) {
	got := getEvenFibonacciSum(1)
	want := big.NewInt(2)
	assertBigIntEqual(got,want,t)
}

func TestFibonacciSumForTwenty(t *testing.T) {
	got := getEvenFibonacciSum(20)
	want := big.NewInt(2026369768940)
	assertBigIntEqual(got,want,t)
}


func TestFibonacciSumForHundred(t *testing.T) {
	got := getEvenFibonacciSumForLength100()
	want,_ := big.NewInt(0).SetString("290905784918002003245752779317049533129517076702883498623284700",0)
	assertBigIntEqual(got,want,t)
}


func assertBigIntEqual(got *big.Int, want *big.Int,t *testing.T){
	if (got.Cmp(want) != 0 ) {
		t.Errorf("got %q want %q", got, want)
	}
}
