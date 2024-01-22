package main

import "fmt"

type Polygons interface {
	Perimeter() float64
}

type Object interface {
	NumberOfSide() int
}

type Square float64

func (s Square) Perimeter() float64 {
	return float64(s * 4.0)
}

func (s Square) NumberOfSide() int {
	return 4
}

func main() {
	var s Polygons = Square(5.0)
	fmt.Println(s.Perimeter())
	var t = s.(Object)
	fmt.Println(t.NumberOfSide())
}
