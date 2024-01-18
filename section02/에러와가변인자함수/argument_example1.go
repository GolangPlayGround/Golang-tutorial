package main

import "fmt"

func main() {

	variadicExample("red")
	variadicExample("yellow", "green", "blue")
}

func variadicExample(s ...string) {
	fmt.Printf("%#v\n", s)
}
