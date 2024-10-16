package perimeter

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{7.0, 8.0}
	got := Perimeter(rectangle)
	want := 30.0

	assertCorrectMessage(t, got, want)
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, want: 36},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("\ngot	%g \nwant	%g", got, tt.want)
			}
		})
	}
}

func assertCorrectMessage(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("\ngot	%.2f \nwant	%.2f", got, want)
	}
}
