package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	result := Repeat("a", 7)
	expected := "aaaaaaa"

	if result != expected {
		t.Errorf("\n\nexpected	%q, \ngot		%q\n\n", expected, result)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 7)
	}
}

func ExampleRepeat() {
	result := Repeat("za", 4)
	fmt.Println(result)
	//Output: zazazaza

}
