package main

import "fmt"

/**
Signed는 음수표현 가능하지만 unsigned는 0부터 시작하는 수.
Unsigned라면 0에서 65535까지, signed라면 -32768~32767까지 표현 가능

int는 signed integer이고 uint는 unsigned integer.
*/

func main() {
	var x uint8 = 200
	fmt.Printf("x = %d\n", x+56) // overflow
	var y int8 = -128
	fmt.Printf("y = %d\n", y-3) //underflow
}
