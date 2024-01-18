package main

import "fmt"

func main() {
	fmt.Println(sum(1, 3, 5, 7, 9))
}

func sum(nums ...int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}
