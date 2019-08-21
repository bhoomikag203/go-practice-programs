package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s' ", got, want)
		}
	}
	t.Run("Say hello to people", func(t *testing.T) {
		got := Hello("Raju", "")
		want := "Hello Raju"	
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Hello world' when an empty string is supplied",  func(t *testing.T) {
		got := Hello("", "")
		want := "Hello world"	
		assertCorrectMessage(t, got, want)
	})

	t.Run("In Spanish", func(t *testing.T) {
		got := Hello("Raju", "Spanish")
		want := "Hola Raju"
		assertCorrectMessage(t, got, want)
	})

	t.Run("In French", func(t *testing.T) {
		got := Hello("Raju", "French")
		want := "Bonjour Raju"
		assertCorrectMessage(t, got, want)
	})

	t.Run("In Kannada", func(t *testing.T) {
		got := Hello("Raju", "Kannada")
		want := "Namaskara Raju"
		assertCorrectMessage(t, got, want)
	})
}
