package main

import "fmt"

func main() {
	sum := 0
	fruits := map[string]int{
		"apple":  1000,
		"banana": 3000,
	}
	for _, val := range fruits {
		sum += val
	}
	fmt.Println(sum)
}
