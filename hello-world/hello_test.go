// https://github.com/quii/learn-go-with-tests/blob/main/hello-world.md
package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		// helper method for narrowing down issues/errors if and when it happens
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Chirag", "English")
		want := "Hello, Chirag"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Spanish speaking", func(t *testing.T) {
		got := Hello("Nirali", "Spanish")
		want := "Hola, Nirali"
		assertCorrectMessage(t, got, want)
	})
	t.Run("French speaking", func(t *testing.T) {
		got := Hello("Nirali", "french")
		want := "Bonjour, Nirali"
		assertCorrectMessage(t, got, want)
	})
}
