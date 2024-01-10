package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	result := Add(3, 2)
	expected := 5

	if result != expected {
		t.Errorf("expected '%d' but got '%d'", expected, result)
	}
}

func ExampleAdd() {
	result := Add(3, 2)
	fmt.Println(result)
	// Output: 5
}
