package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		panic("2개의 인자가 필요")
	}
	if err := os.Rename(args[0], args[1]); err != nil {
		panic(err)
	}

	fmt.Printf("파일 이름 변경 성공하였습니다 : %s => %s", args[0], args[1])
}
