package hello_world

import (
	"testing"
)

/*
Let's go over the cycle again
Write a test
Make the compiler pass
Run the test, see that it fails and check the error message is meaningful
Write enough code to make the test pass
Refactor
*/
func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)

	})

	t.Run("Empty string defaults to 'world'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("In spanish", func(t *testing.T) {
		got := Hello("Juan", "Spanish")
		want := "Hola, Juan"
		assertCorrectMessage(t, got, want)
	})

	t.Run("In french", func(t *testing.T) {
		got := Hello("Mbappe", "French")
		want := "Bonjour, Mbappe"
		assertCorrectMessage(t, got, want)
	})

}

// testing TB is an interface that *testing.T and *testing.B both satisfy.
func assertCorrectMessage(t testing.TB, got, want string) {
	//t helper allow you to declare handler functions for your tests
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
