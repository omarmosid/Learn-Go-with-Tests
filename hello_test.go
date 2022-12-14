package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying Hello to people", func(t *testing.T) {
		got := Hello("Omar", "English")
		want := "Hello Omar"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying Hello World, when an empty string is passed", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola Elodie"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %s | Want %s", got, want)
	}
}
