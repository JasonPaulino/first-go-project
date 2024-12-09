package main

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("In Spanish", func (t *testing.T) {
		got := Hello ("Jason", spanish)
		want := "Hola, Jason"
		assertCorrectMessage(t, got, want)
	})

	t.Run("In French", func (t *testing.T) {
		got := Hello ("Jason", french)
		want := "Salut, Jason"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying hello to people", func (t *testing.T) {
		got := Hello("Jason", english)
		want := "Hello, Jason"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Hello, World' when there is no name input", func (t *testing.T) {
		got := Hello("", english)
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}

// Documentation test
func ExampleHello() {
	got := Hello("Kevin", "Spanish")
	fmt.Println(got)
	// Output: Hola, Kevin
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}