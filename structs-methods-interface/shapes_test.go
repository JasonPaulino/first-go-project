package main

import (
	"testing"
)

func TestPerimeter(t *testing.T){
	t.Run("rectangle perimeter", func (t *testing.T) {
		rectangle := Rectangle{2.0, 2.0}
		want:= 8.0

		checkPerimeter(t, rectangle, want)
	})
	t.Run("triangle perimeter", func (t *testing.T) {
		triangle := Triangle{3.0, 1.2, 2.2}
		want := 6.4

		checkPerimeter(t, triangle, want)
	})
}

func TestArea(t *testing.T) {
	t.Run("rectangle area", func(t *testing.T) {
		rectangle := Rectangle{2.0, 2.0}
		want := 4.0

		checkArea(t, rectangle, want)
	})
	t.Run("triangle area", func(t *testing.T) {
		triangle := Triangle{3.0, 1.2, 1}
		want := 1.7999999999999998

		checkArea(t, triangle, want)
	})
}

func checkArea(t testing.TB, shape Shape, want float64) {
	t.Helper()

	got := shape.Area()

	if got != want {
		t.Errorf("got %g, want %g", got, want)
	}
}

func checkPerimeter(t testing.TB, shape Shape, want float64) {
	t.Helper()

	got := shape.Perimeter()

	if got != want {
		t.Errorf("got %g, want %g", got, want)
	}
}