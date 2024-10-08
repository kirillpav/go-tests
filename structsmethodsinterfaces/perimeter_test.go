package main

import (
	"testing"
)

func TestPerimeter(t *testing.T){
	
		rectangle := Rectangle{10.0, 10.0}
		got := Perimeter(rectangle)
		want := 40.0

		if got != want {
			t.Errorf("got %.2g want %.2g", got, want)
		}
	
}

func TestArea(t *testing.T){
	// Table driven tests
	areaTests := []struct {
		shape Shape
		want float64
	} {
		{shape: Rectangle{Width: 12.0, Height: 6.0}, want: 72.0},
		{shape: Circle{Radius: 10}, want: 314.1592653589793},
		{shape: Triangle{Base: 12.0, Height: 6.0}, want: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want{
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}