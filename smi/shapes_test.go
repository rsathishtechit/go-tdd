package smi

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 8}, want: 48},
		{name: "Circle", shape: Circle{Radius: 10}, want: 314.1592653589793},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
		}
	}
}
