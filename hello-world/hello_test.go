package main

import "testing"


func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func (t *testing.T) {
		got := Hello("Jason")
		want := "Hello, Jason"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Hello, World' when there is no name input", func (t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}