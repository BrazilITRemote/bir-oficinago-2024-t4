package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Lucas")

	got := buffer.String()
	want := "Hello, Lucas"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}