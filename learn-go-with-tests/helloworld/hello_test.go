package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("Matt", "English")
		want := "Hello, Matt"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say Hello World when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Fred", "French")
		want := "Bonjour, Fred"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in German", func(t *testing.T){
		got := Hello("Bill", "German")
		want := "Gutentag, Bill"
		assertCorrectMessage(t, got, want)
	})

}
