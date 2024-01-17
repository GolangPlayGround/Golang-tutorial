package main

import "fmt"

func main() {
	/** numbers 슬라이스를 만들고 100부터 120까지 정수를 numbers에 추가**/
	sum := 0
	numbers := make([]int, 0)
	for i := 100; i <= 120; i++ {
		numbers = append(numbers, i)
	}
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	fmt.Println(sum)

	/** 다음 슬라이스를 사용하여 원소 개수가 2개씩만 가지도록 슬라이싱하여 출력**/
	alphabet := []string{"a", "b", "c", "d", "e"}
	for i := 0; i < len(alphabet)-1; i++ {
		fmt.Println(alphabet[i : i+2])
	}
}
