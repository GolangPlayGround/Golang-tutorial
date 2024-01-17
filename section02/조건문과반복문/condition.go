package main

import "fmt"

func main() {
	condition1(30)
	condition2(40)
	condition3()
	condition4()
}

/**
40대 이상이면 "40대 이상 출력"
*/

func condition1(age int) {
	if age < 20 {
		fmt.Println("10대")
	} else if age < 30 {
		fmt.Println("20대")
	} else if age < 40 {
		fmt.Println("30대")
	} else {
		fmt.Println("40대 이상")
	}
}

func condition2(age int) {
	switch {
	case age < 20:
		fmt.Println("10대")
	case age < 30:
		fmt.Println("20대")
	case age < 40:
		fmt.Println("30대")
	default:
		fmt.Println("40대이상")
	}
}

/**
1~200까지 홀수만 더하기
*/

func condition3() {
	sum := 0
	for i := 1; i <= 200; i += 2 {
		sum += i
	}
	fmt.Println(sum)
}

func condition4() {
	sum, i := 0, 1
	for i <= 200 {
		sum += i
		i += 2
	}
	fmt.Println(sum)
}
