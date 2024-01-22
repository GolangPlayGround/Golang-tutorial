package main

import "fmt"

func main() {
	finished := make(chan bool) // 채널 생성
	go func() {
		fmt.Println("Hello World")
		finished <- true // 채널에 데이터 보냄
	}()
	<-finished // 채널에서 데이터를 바등ㄹ때까지 block

	fmt.Println("Program Ended")
}
