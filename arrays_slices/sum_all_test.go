package main

import (
	"reflect"
	"testing"
)

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{5, 5, 5})
	want := []int{6, 15}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}