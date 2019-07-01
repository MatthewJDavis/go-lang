package strings

import "testing"

func TestCaseConverter(t *testing.T) {

	assertCorrectCase := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("convert to lowercase", func(t *testing.T){
		lower := CaseConverter("Matt", "lower")
		expected := "matt"

		assertCorrectCase(t, lower, expected)
	})
	t.Run("convert to uppercase", func(t *testing.T){
		upper := CaseConverter("Matt", "upper")
		expected := "MATT"
		assertCorrectCase(t, upper, expected)
	})
	
}
