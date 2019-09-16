package hello

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q wants %q", got, want)
		}
	}

	t.Run("will return simple hello + name", func(t *testing.T) {
		got := Hello("Ardi", "")
		want := "Hello, Ardi"
		assertCorrectMessage(t, got, want)
	})

	t.Run("will return default value world if name is empty", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Ardi", "Spanish")
		want := "Hola, Ardi"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in french", func(t *testing.T) {
		got := Hello("Ardi", "French")
		want := "Bonjour, Ardi"
		assertCorrectMessage(t, got, want)
	})

}
