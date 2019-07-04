package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	got := dictionary.Search("test")
	want := "this is just a test"
	assertStrings(t, got, want)
}
func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if want != got {
		t.Errorf("wanted '%s', but got '%s'", want, got)
	}
}
