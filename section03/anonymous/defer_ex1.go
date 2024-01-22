package main

import "fmt"

func main() {
	var fn = func(i int) {
		fmt.Println(i)
	}

	defer fn(3)
	defer fn(2)
	fn(1)
}
