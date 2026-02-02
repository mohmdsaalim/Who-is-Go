package main

import (
	"fmt"
	"math"
)

type Shape interface{
	Area() float64
}

type Rectangle struct {
	width, height float64
}
type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func calculateArea(r Shape) float64 {
	return r.Area()
}
func main() {
	rect := Rectangle{width: 10, height: 20}
	cir := Circle{radius: 10}
	// a := calculateArea(rect)
	// b := calculateArea(cir)

	fmt.Println(calculateArea(rect))
    fmt.Println(calculateArea(cir))

}

// interface composition and creating errors with the interface