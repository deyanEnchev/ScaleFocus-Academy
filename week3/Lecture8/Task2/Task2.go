package main

import (
	"fmt"
	"math"
)

type Square struct {
	a float64
}

type Circle struct {
	r float64
}

func (s Square) NewSquare(a float64) Square {
	return Square{a}
}

func (c Circle) NewCircle(r float64) Circle {
	return Circle{r}
}

func (s Square) Area() float64 {
	return s.a * s.a
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.r, 2) 
}

type Shape interface {
	Area() float64
}

type Shapes []Shape 

func (s Shapes) LargestArea() float64 {
	var largestArea float64
	for _, v := range s {
		if v.Area() > largestArea {
			largestArea = v.Area()
		}
	}
	return largestArea
}

func main() {
	s := Square{5}
	c := Circle{5}

	var slice Shapes = []Shape{s,c}

	fmt.Println(slice.LargestArea())
}