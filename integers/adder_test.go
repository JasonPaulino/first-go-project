package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Adder(3, 2)
	expected := 5

	assertCorrectMessage(t, sum, expected)
}

func ExampleAdder() {
	sum := Adder(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func assertCorrectMessage(t testing.TB, sum, expected int) {
	t.Helper()

	if sum != expected {
		t.Errorf("sum: %d, expected: %d", sum, expected)
	}
}