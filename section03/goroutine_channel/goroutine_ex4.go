package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func sum4(start, end int, sum *uint64, wg *sync.WaitGroup) {
	defer wg.Done()
	s := 0
	for i := start; i <= end; i++ {
		s += i
	}
	atomic.AddUint64(sum, uint64(s))
}

func main() {
	var s = uint64(0)
	var wg sync.WaitGroup
	wg.Add(5)
	go sum4(1234567, 3456789012, &s, &wg)
	go sum4(2345678, 3456789012, &s, &wg)
	go sum4(3456789, 3456789012, &s, &wg)
	go sum4(4567890, 3456789012, &s, &wg)
	go sum4(5678901, 3456789012, &s, &wg)
	wg.Wait()
	fmt.Println(atomic.LoadUint64(&s))
}
