package hello

import "testing"

func TestHello(t *testing.T) {

	t.Run("say hello", func(t *testing.T) {
		got := Hello("Ardi", "")
		want := "Hello, Ardi"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("say hello return default value world", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("say hello in spanish", func(t *testing.T) {
		got := Hello("Ardi", "Spanish")
		want := "Hola, Ardi"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	})

	t.Run("say hello in france", func(t *testing.T) {
		got := Hello("Ardi", "France")
		want := "Bonjour, Ardi"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

}
