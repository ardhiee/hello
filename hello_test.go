package hello

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q wants %q", got, want)
		}
	}

	t.Run("will return simple hello world", func(t *testing.T) {
		got := Hello()
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

}
