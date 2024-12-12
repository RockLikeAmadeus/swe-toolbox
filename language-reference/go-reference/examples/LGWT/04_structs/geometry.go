package geometry

import "math"

// Interface implementation is implicit
type Shape interface {
	Area() float64
}

/* Rectangle */

type Rectangle struct {
	Width  float64
	Height float64
}

// Methods in Go are defined using *receivers*
// Convention is to call the receiver variable the first letter of the type
// (see [this](https://github.com/golang/go/wiki/CodeReviewComments#receiver-names).)
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

/* Circle */

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

/* Triangle */

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
