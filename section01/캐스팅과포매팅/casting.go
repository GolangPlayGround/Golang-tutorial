package main

import (
	"fmt"
)

/**

var x int8 = 5 로 선언된 x를 int32로 변경하고 싶다면
var y int 32 = int32(x) 처럼 사용하면 된다.

다만, 캐스팅시 데이터 범위를 염두해둬야 한다

int8 -> int32의 경우처럼 넓은 범위로 캐스팅은 안전하지만

int8 -> uint8로 캐스팅 할 때는 음수때문에 문제가 생길 수 있다
int32 -> int8도 마찬가지

정수를 문자열형으로 변경하거나 면자열형을 정수로 변경해야할때는 strconv패키지를 사용해야한다

- int -> string : strconv.Itoa(값 또는 변수)
- string -> int : strconv.Atoi(값 또는 변수)
- float64 -> string : strconv.FormatFloat(값 또는 변수, 'f', -1, 64)
- string -> float64 : strconv.ParseFloat(값 또는 변수 , 64)

*/

func main() {
	casting_practice1()
	casting_practice2()
}

/*
*
x변수를 uint8 200으로 초기화 후
y변수는 int8로 선언 후 -100으로 초기화
z변수는 int형으로 선언후 x+y값 대입후 출력
*/
func casting_practice1() {
	var x uint8 = 200
	var y int8 = -100
	z := int(x) + int(y)
	fmt.Printf("%d\n", z)
}

/**
x변수를 uint8로 선언후 200으로 초기화
y변수는 float32로 선언후 -100.0으로 초기화
z변수는 float64로 선언후 x+y 대입
*/

func casting_practice2() {
	var x uint8 = 200
	var y float32 = -100.0
	z := float64(x) + float64(y)
	fmt.Printf("%f\n", z)
}
