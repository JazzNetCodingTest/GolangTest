package main

import (
	"testing"
	"reflect"
	"sort"
)

func TestIntersectionForBothEmptyArrays(t *testing.T) {
	var a = []int{}
	var b = []int{}
	got := intersection(a, b)
	var want []int
	assertArrayEqual(got, want, t)
}


func TestIntersectionForBothOneEmptyArray(t *testing.T) {
	var a = []int{1}
	var b = []int{}
	got := intersection(a, b)
	var want []int
	assertArrayEqual(got, want, t)
}


func TestIntersectionForNonduplicateElements(t *testing.T) {
	var a = []int{2,4,6,8,10,12}
	var b = []int{3,6,9,12,15}
	got := intersection(a, b)
	want := []int {12,6}
	assertArrayEqual(got, want, t)
}

func TestIntersectionForWithduplicateElements(t *testing.T) {
	var a = []int{2,4,6,8,10,12,8,4,0}
	var b = []int{3,6,9,12,15,9,1}
	got := intersection(a, b)
	want := []int {6,12}
	assertArrayEqual(got, want, t)
}


func assertArrayEqual(got []int,want []int, t *testing.T){
	sort.Ints(got)
	sort.Ints(want)
	if (!reflect.DeepEqual(got,want) ) {
		t.Errorf("got %v want %v", got, want)
	}
}