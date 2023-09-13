package structs

import "testing"

func TestArea(t *testing.T) {

	t.Run("rectangles, circles and triangles", func(t *testing.T) {
		areas := []struct {
			name  string
			shape Shape
			want  float64
		}{
			{name: "Rectangle", shape: Rectangle{Width: 10, Height: 20}, want: 200},
			{name: "Circle", shape: Circle{Radius: 10}, want: 314.1592653589793},
			{name: "Triangle", shape: Triangle{Width: 10, Height: 20}, want: 100},
		}

		for _, a := range areas {
			t.Run(a.name, func(t *testing.T) {
				got := a.shape.Area()

				if got != a.want {
					t.Errorf("%#v got %g want %g", a, got, a.want)
				}
			})
		}

	})

}
