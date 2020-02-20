package main

import "fmt"

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLenth float64
}

type shape interface {
	getArea() float64
}

func main() {
	s := square{sideLenth: 4}
	t := triangle{base: 4, height: 8}

	printArea(s)
	printArea(t)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64 {
	return .5 * t.height * t.base
}

func (s square) getArea() float64 {
	return s.sideLenth * s.sideLenth
}
