package iteration

import (
	"fmt"
	"testing"
)

// This is the experimental iteration module
func Example() {
	fmt.Println("Printing... ", RepeatFun("bubu", 2))
	//output: "bubububu"
}

func TestIteration(t *testing.T) {
	t.Run("Should return a 5 times", func(t *testing.T) {
		got := RepeatFun("a", 5)
		want := "aaaaa"

		if got != want {
			t.Errorf("espected %q, but got %q", want, got)
		}
	})
}

func BenchmarkIteration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatFun("a", 5)
	}
}
