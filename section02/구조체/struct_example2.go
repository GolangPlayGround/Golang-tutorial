package main

import "github.com/davecgh/go-spew/spew"

type rectangle_ex2 struct {
	length  float64
	breadth float64
	color   string
	geo     geometry
}

type geometry struct {
	area      float64
	perimeter float64
}

func main() {
	var rect1 rectangle_ex2
	rect1.length = 10
	rect1.breadth = 15
	rect1.geo.area = rect1.length * rect1.breadth
	spew.Dump(rect1)

	rect2 := rectangle_ex2{10, 20, "Green", geometry{200, 60}}
	spew.Dump(rect2)

	rect3 := rectangle_ex2{length: 5, breadth: 10}
	spew.Dump(rect3)

	rect4 := new(rectangle_ex2)
	rect4.length = 5
	rect4.breadth = 12
	spew.Dump(rect4)
}
