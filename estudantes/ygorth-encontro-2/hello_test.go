package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Bocao")
	want := "Hello, Bocao"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
