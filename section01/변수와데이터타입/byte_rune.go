package main

import "fmt"

/**
byte는 uint8의 별칭이며, rune은 int32의 별칭

Go언어에는 하나의 문자를 표현하는 char같은 데이터 형이 없음

rune은 utf-8 형식의 문자 하나를 표현하는데 사용되고

byte타입은 ascii형식의 문자 하나를 표현하는데 사용
*/

func main() {
	var b byte = 'B'
	fmt.Printf("b = %s\n", string(b))
	R := '한'
	fmt.Printf("R = %s\n", string(R))
}
