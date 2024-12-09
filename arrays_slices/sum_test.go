package main

import (
	"fmt"
	"testing"
)

// Test
func TestSum(t *testing.T) {
	t.Run("Array with 5 elements", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
	
		sum := Sum(numbers)
		expected := 15
	
		assertCorrectMessage(t, sum, expected, numbers)
	})
	t.Run("Slice without fixed length", func(t *testing.T) {
		numbers := []int{1, 2, 3}
	
		sum := Sum(numbers)
		expected := 6
	
		assertCorrectMessage(t, sum, expected, numbers)
	})
}

func assertCorrectMessage(t testing.TB, sum, expected int, numbers []int) {
	t.Helper()

	if sum != expected {
		t.Errorf("sum: %v, expected: %v, given: %v", sum, expected, numbers)
	}
}

// Benchmarks
func BenchmarkSum(b *testing.B) {
	numbers := []int{1, 2, 3, 4, 5}

	for i := 0; i < b.N; i++ {
		Sum(numbers)
	} 
}

// Documentation
func ExampleSum() {
	numbers := []int{1, 2, 3, 4, 5}

	sum := Sum(numbers)
	fmt.Println(sum)
	// Output: 15
}