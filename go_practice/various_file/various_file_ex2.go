package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rfp, err := os.Open("test2.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := rfp.Close(); err != nil {
			panic(err)
		}
	}()

	wfp, err := os.Create("test3.txt")
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := wfp.Close(); err != nil {
			panic(err)
		}
	}()

	rsc := bufio.NewScanner(rfp)
	rsc.Split(bufio.ScanLines)
	wsc := bufio.NewWriter(wfp)
	for rsc.Scan() {
		line := rsc.Text()
		fmt.Println(line)
		_, err = wsc.WriteString(line + "\n")
		if err != nil {
			panic(err)
		}
		_ = wsc.Flush()
	}
}
