package main

import "testing"

var dictionary = Dictionary{"test": "this is just a test"}

func TestDictionary(t *testing.T) {
	t.Run("known word", func (t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("uknown")
		if got == nil {
			t.Fatal("expected an error")
		}

		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	dictionary.Add("rizz", "a form of someone being suave")

	want := "a form of someone being suave"
	got, _ := dictionary.Search("rizz")

	assertStrings(t, got, want)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want.Error())
	}
}