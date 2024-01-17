package main

import "fmt"

func main() {
	a, b, c, d := fourArithmetic(11, 5)
	fmt.Println(a, b, c, d)

	num, str := 4, "hello"
	println(num, ":", str)
	change(&num, &str)
	println(num, ":", str)
}

func fourArithmetic(a int, b int) (int, int, int, int) {
	return a + b, a - b, a * b, a / b
}

func change(n *int, s *string) {
	*n *= 2
	*s += " World"
}
