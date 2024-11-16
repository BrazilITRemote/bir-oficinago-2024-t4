package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Lucas")
	want := "Hello, Lucas"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}