package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat a character 5 times by default", func(t *testing.T) {
		got := Repeat("a", 0)
		want := "aaaaa"

		if got != want {
			t.Errorf("expected %q but got %q", got, want)
		}
	})

	t.Run("repeat a character for given times", func(t *testing.T) {
		got := Repeat("a", 10)
		want := "aaaaaaaaaa"

		if got != want {
			t.Errorf("expected %q but got %q", got, want)
		}
	})

}

func ExampleRepeat() {
	defaultRepeats := Repeat("b", 0)
	fmt.Println(defaultRepeats)
	// Output: bbbbb
}

func ExampleRepeat_givenTimes() {
	givenRepeatTimes := Repeat("b", 3)
	fmt.Println(givenRepeatTimes)
	// Output: bbb
}
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 1000)
	}
}
