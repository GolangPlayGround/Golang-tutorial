package main

import "github.com/davecgh/go-spew/spew"

type rectangle_ex1 struct {
	length  float64
	breadth float64
	color   string
}

func main() {
	spew.Dump(rectangle_ex1{10.5, 12.5, "red"})
}
