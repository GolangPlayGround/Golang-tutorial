package main

import "github.com/davecgh/go-spew/spew"

type rectangle struct {
	length  float64
	breadth float64
	color   string
}

func (r rectangle) area() float64 {
	return r.length * r.breadth
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.length + r.breadth)
}

func main() {
	s1 := rectangle{
		length:  3,
		breadth: 5,
		color:   "red",
	}
	spew.Dump(s1.area(), s1.perimeter())
}
