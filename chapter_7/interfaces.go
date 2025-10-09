package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) Area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func (c *Circle) Area() float64 {
	return math.Pi * c.r * c.r // πr²
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b) // Pythagorean theorem
}

func totalArea(shapes ...Shape) float64 {
	var total float64
	for _, s := range shapes {
		total += s.Area()
	}
	return total
}

func main() {
	r := &Rectangle{0, 0, 10, 10}
	c := &Circle{0, 0, 5}

	fmt.Printf("Total area: %f\n", totalArea(r, c))

}