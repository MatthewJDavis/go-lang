package greeting

import "testing"

func TestGreeting(t *testing.T) {
	expected := "Hello Go"
	actual := Hi("Go")
	if actual != expected {
		t.Error("Test failed")
	}
}
