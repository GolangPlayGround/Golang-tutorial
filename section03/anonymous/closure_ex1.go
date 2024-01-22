package main

import "fmt"

func main() {
	for i := 10.0; i < 100.0; i += 10.0 {
		cv := func() float64 {
			return i * 39.370
		}
		fmt.Printf("%.0f미터 => %6.1f인치\n", i, cv())
	}
}
