package hello

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q wants %q", got, want)
		}
	}

	t.Run("Can receive name parameter", func(t *testing.T) {
		got := Hello("Ardi", "")
		want := "Hello, Ardi"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Have a default value where name is empty", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in spanish additional language support", func(t *testing.T) {
		got := Hello("Ardi", "Spanish")
		want := "Hola, Ardi"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in french additional language support", func(t *testing.T) {
		got := Hello("Ardi", "French")
		want := "Bonjour, Ardi"
		assertCorrectMessage(t, got, want)
	})
}
