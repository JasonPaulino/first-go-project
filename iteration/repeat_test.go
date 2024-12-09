package iteration

import (
	"fmt"
	"testing"
)

// Tests
func TestRepeat(t *testing.T) {
	got := Repeat(5, "a")
	want := "aaaaa"

	assertCorrectMessage(t, got, want)
}

// Benchmarks
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat(2, "c")
	}
}

// Documentation
func ExampleRepeat() {
	got := Repeat(2, "a")
	fmt.Println(got)
	// Output: aa
}

func assertCorrectMessage(t testing.TB, got , want string) {
	t.Helper()

	if got != want {
		t.Errorf("got: %q, want: %q", got, want)
	}
}