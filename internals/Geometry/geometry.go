package geometry

import "math"

// interface untuk bentuk geometris
type Geometry interface {
	Area() float64
}

// implementasi untuk circle
type Cirlce struct {
	Radius float64
}

// method area untuk circle
func (c Cirlce) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// implementasi untuk rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// method area untuk rectangle
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

// fungsi kalkulator untuk menghitung total area
func CalculatorArea(shapes []Geometry) float64{
	total := 0.0
	for _, shape := range shapes{
		total += shape.Area()
	}
	return total
}
