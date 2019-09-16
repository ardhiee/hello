package hello

import "testing"

func TestHello(t *testing.T) {

	t.Run("will return simple hello world", func(t *testing.T) {
		got := Hello()
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q wants %q", got, want)
		}
	})
}
