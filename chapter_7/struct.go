package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}



func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b) // Pythagorean theorem
}

func rectangleArea(x1, y1, x2, y2 float64) float64 {
	l := distance(x1, y1, x1, y2)
	w := distance(x1, y1, x2, y1)
	return l * w
}

func circleArea(x, y, r float64) float64 {
	return math.Pi * r * r // πr²
}

func rectangleAreaStruct(r Rectangle) float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func circleAreaStruct(c Circle) float64 {
	return math.Pi * c.r * c.r // πr²
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r // πr²
}

func main() {

	var rx1, ry1 float64 = 0, 0
	var rx2, ry2 float64 = 10, 10
	
	var cx, cy, cr float64 = 0, 0, 5

	fmt.Printf("Rectangle area: %f\n", rectangleArea(rx1, ry1, rx2, ry2))
	fmt.Printf("Circle area: %f\n", circleArea(cx, cy, cr))


	// Using structs
	fmt.Println("Using structs:")
	r := Rectangle{0, 0, 10, 10}
	c := Circle{0, 0, 5}

	fmt.Printf("Rectangle area: %f\n", rectangleArea(r.x1, r.y1, r.x2, r.y2))
	fmt.Printf("Circle area: %f\n", circleArea(c.x, c.y, c.r))

	fmt.Println("Using structs with functions that take structs:")
	fmt.Printf("Rectangle area: %f\n", rectangleAreaStruct(r))
	fmt.Printf("Circle area: %f\n", circleAreaStruct(c))

	// Using methods
	fmt.Println("Using methods:")
	fmt.Printf("Rectangle area: %f\n", r.area())
	fmt.Printf("Circle area: %f\n", c.area())
}