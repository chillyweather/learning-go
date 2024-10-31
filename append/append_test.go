package append

import (
	"reflect"
	"testing"
)

func TestAppend(t *testing.T) {
	testStruct := []cost{
		{0, 4.0},
		{1, 2.1},
		{5, 2.5},
		{1, 3.1},
	}
	got := getCostsByDay(testStruct)

	want := []float64{
		4.0,
		5.2,
		0.0,
		0.0,
		0.0,
		2.5,
	}

	isEqual := reflect.DeepEqual(got, want)

	if isEqual == false {
		t.Errorf("\nwant %f\ngot  %f\n", want, got)
	}
}
