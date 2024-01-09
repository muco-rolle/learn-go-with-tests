package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Tresor")
	want := "Hello, Tresor"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
