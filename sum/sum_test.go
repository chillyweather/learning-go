package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size  	/", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(got []int, want []int, t *testing.T) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got	%v, want	%v", got, want)
		}
	}
	t.Run("make sum of some slices", func(t *testing.T) {
		got := SumAllTails([]int{2, 6}, []int{5, 6})
		want := []int{6, 6}

		checkSums(got, want, t)
	})

	t.Run("make sum with empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{5, 6})
		want := []int{0, 6}

		checkSums(got, want, t)
	})
}
