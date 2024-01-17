package main

import "fmt"

/**
주소 연산자 &
해당 변수의 주소를 가리키게되며 변수에만 사용이 가능. 포인터 연산자*는 주소 연산자가 가리키는
주소의 값을 가져오거나, 변수 선언 시 포인터 변수라는 것을 선언할 때 사용
*/

func main() {
	pointer_practice()
}

/**
문자열 포인터를 선언하여 "hello"에서 "world"로 변환되는 과정을 출력
*/

func pointer_practice() {
	var x = "hello"
	var p *string
	p = &x
	fmt.Println(x)
	*p = "world"
	fmt.Println(x)
}
