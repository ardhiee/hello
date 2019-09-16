package hello

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q wants %q", got, want)
		}
	}

	t.Run("Hello with input name", func(t *testing.T) {
		got := Hello("Ardi")
		want := "Hello, Ardi"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Hello with empty input shall return default value of world", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})
}
