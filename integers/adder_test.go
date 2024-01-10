package integers

import "testing"

func TestAdder(t *testing.T) {
	result := Add(3, 2)
	expected := 5

	if result != expected {
		t.Errorf("expected '%d' but got '%d'", expected, result)
	}
}
