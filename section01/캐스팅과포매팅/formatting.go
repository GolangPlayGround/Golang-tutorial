package main

import (
	"fmt"
	"strconv"
)

/**
논리형 포매팅의 경우 %t를 사용

문자열의 경우 %s를 사용

ex) %3s 란 문자열형을 3자리에 맞추어 출력하는데 3자리가 안되면 공백을 앞에 채워준다.
다만 3자리가 넘어가서 5글자라면 5글자 그대로 출력


정수형의 경우 십진수로 출력하고 싶을때 %d 2진수 %b 8진수 %o 16진수 %x를 사용

실수형의 경우 %f 사용 %9.3f라고 하면 전체길이는 9자리이고 실수부는 3자리에 맞춰 출력


특정 변수나 값이 어떤 타입인지를 알려면 %T

이변수가 정수형인지 논리형인지 자동 판단하여 알아서 값 출력 %v
*/

func main() {
	formatting_practice1()
	formatting_practice2()
	formatting_practice3()
	formatting_practice4()
}

/**
x 변수를 문자열형 "200"으로 초기화한후 int형 캐스팅한후 출력
y변수는 float64로 선언후 -100.0으로 초기화 문자열형 캐스팅후 출력
*/

func formatting_practice1() {
	x := "200"
	intX, err := strconv.Atoi(x)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d\n", intX)
	y := -100.0
	strY := strconv.FormatFloat(y, 'f', -1, 64)
	fmt.Printf("%s\n", strY)
}

/*
*
x: 뱐수를 논리형 타입으로 선언 한 후에 값의 초기화 없이 그값을 출력
x 변수를 true로 대입한 후에 다시 그값을 출력
*/
func formatting_practice2() {
	var x bool
	fmt.Printf("%t\n", x)
	x = true
	fmt.Printf("%t\n", x)
}

/**
a, b, c, d 변수를 문자혈 타입으로 선언한 후 "nice" "to" "meet" "you"
로 초기화 후 각각"[%3s]"로 출력
*/

func formatting_practice3() {
	a, b, c, d := "nice", "to", "meet", "you"
	fmt.Printf("[%3s] [%3s] [%3s] [%3s]\n", a, b, c, d)
}

/**
x 변수를 31로 초괴화 한후에 x변수를 2진수,8진수,16진수,10진수로 출력
y를 12로 초기화후 "[%03d]", "[%3d]"로 출력
*/

func formatting_practice4() {
	x := 31
	fmt.Printf("%b %o %x %d\n", x, x, x, x)
	y := 12
	fmt.Printf("[%03d] [%3d]\n", y, y)
}
