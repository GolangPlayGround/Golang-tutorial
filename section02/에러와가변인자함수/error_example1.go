package main

import (
	"fmt"
	"strconv"
)

func main() {
	i1, err := divide("25", "5")
	if err != nil {
		panic(err)
	}
	fmt.Println(i1)
	l2, err := divide("25", "hello")
	if err != nil {
		panic(err)
	}
	fmt.Println(l2)
}

func divide(i1 string, i2 string) (int, error) {
	o1, err := strconv.Atoi(i1)
	if err != nil {
		return 0, err
	}
	o2, err := strconv.Atoi(i2)
	if err != nil {
		return 0, err
	}

	if o2 == 0 {
		return 0, fmt.Errorf("0으로 나누기 할 수 없다")
	}
	return o1 / o2, nil
}
