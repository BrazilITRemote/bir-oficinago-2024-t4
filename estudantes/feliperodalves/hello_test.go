package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Felipe")
	want := "Hello, Felipe"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
