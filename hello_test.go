package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Ardi", "")
		want := "Hello, Ardi"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, world' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Ardi", "Spanish")
		want := "Hola, Ardi"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Ardi", "French")
		want := "Bonjour, Ardi"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in indonesia", func(t *testing.T) {
		got := Hello("Ardi", "Indonesia")
		want := "Halo, Ardi"
		assertCorrectMessage(t, got, want)
	})
}
