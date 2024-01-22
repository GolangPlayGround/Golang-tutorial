package main

import "github.com/davecgh/go-spew/spew"

func main() {
	var area = func(l, b int) int {
		return l * b
	}

	spew.Dump(area(20, 30))

	func(l int, b int) {
		spew.Dump(l * b)
	}(20, 20)
}
