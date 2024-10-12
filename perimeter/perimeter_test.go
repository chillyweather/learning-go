package perimeter

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(7.0, 8.0)
	want := 30.0

	assertCorrectMessage(t, got, want)
}

func assertCorrectMessage(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("\ngot	%.2f \nwant	%.2f", got, want)
	}
}
