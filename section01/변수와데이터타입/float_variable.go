package main

import "fmt"

/**
실수형은 unsigned가 없음.

float32, float64, complex64, complex128만 존재

complex64와 complex128은 복소수 타입이라 현실에서 쓸 일 거의 없음
*/

func main() {
	x := 3.14159
	fmt.Printf("%g\n", x/2.71728)
}
