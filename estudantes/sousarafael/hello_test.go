package main

import "testing"

//func TestHello(t *testing.T) {
//	got := Hello()
//	want := "Hello, world"
//
//	if got != want {
//		t.Errorf("got %q want %q", got, want)
//	}
//}

func TestHello(t *testing.T) {
	got := Hello("Chris")  //futuro
	want := "Hello, Chris" //o que eu quero que aconteça no futuro!

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
