package injection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
Greet(&buffer, "Matt")

got := buffer.String()
want := "Hello, Matt"

if got != want {
	t.Errorf("wanted '%s' but got '%s'", want, got)
}
}