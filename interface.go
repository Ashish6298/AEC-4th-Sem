package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func PrintArea(s Shape) {
	fmt.Printf("Area : %0.2f\n", s.Area())
}

func main() {
	/*
		Output:
		Area : 78.54
		Area : 24.00
	*/
	circle := Circle{radius: 5}
	Rectangle := Rectangle{width: 4, height: 6}

	PrintArea(circle)
	PrintArea(Rectangle)
}