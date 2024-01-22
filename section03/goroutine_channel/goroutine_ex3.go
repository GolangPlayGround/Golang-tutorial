package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func sum3(start, end int) {
	defer wg.Done()
	s := 0
	for i := start; i <= end; i++ {
		s += i
	}
	fmt.Printf("sum(%d ~ %d) = %d\n", start, end, s)
}

func main() {
	wg.Add(5)
	s1 := time.Now().UnixMilli()
	go sum3(1234567, 3456789012)
	go sum3(2345678, 3456789012)
	go sum3(3456789, 3456789012)
	go sum3(4567890, 3456789012)
	go sum3(5678901, 3456789012)
	wg.Wait()
	fmt.Printf("Elapsed Time: %dms\n", time.Now().UnixMilli()-s1)
}
