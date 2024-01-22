package main

import "fmt"

func main() {
	n := 10
	c := make(chan int, n)
	go func() {
		x, y := 0, 1
		for i := 0; i < n; i++ {
			c <- x
			x, y = y, x+y
		}
		close(c)
	}()
	for i := range c {
		fmt.Println(i)
	}
}
