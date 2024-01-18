package main

import (
	"example.com/m/src/string_util"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"strings"
)

func main() {
	str := "seafood"

	fmt.Println(strings.Contains(str, "foo"))
	fmt.Println(strings.Contains(str, "bar"))

	x := "hello world"
	y := &x
	fmt.Printf("%#v\n", y)
	spew.Dump(y)

	println("\n==============\n")

	fmt.Println(string_util.MyContains("seafood", "foo"))
}
