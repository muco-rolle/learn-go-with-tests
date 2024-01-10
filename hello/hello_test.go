package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Tresor", "")
		want := "Hello, Tresor"
		assertEqual(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertEqual(t, got, want)
	})

	t.Run("saying hello to people in Spanish", func(t *testing.T) {
		got := Hello("Tresor", "es")
		want := "Hola, Tresor"
		assertEqual(t, got, want)
	})

	t.Run("saying hello to people in French", func(t *testing.T) {
		got := Hello("Tresor", "fr")
		want := "Salut, Tresor"
		assertEqual(t, got, want)
	})

	t.Run("saying hello to people in Kirundi", func(t *testing.T) {
		got := Hello("Tresor", "rn")
		want := "Yambu, Tresor"
		assertEqual(t, got, want)
	})
}

func assertEqual(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
