package main

import "fmt"

type Printer interface {
	Print()
}

type emp struct {
	name, address string
}

func (e *emp) Print() {
	fmt.Printf("Name: %s\naddress: %s\n", e.name, e.address)
}

func (e *emp) Assign(n, a string) {
	e.name = n
	e.address = a
}

func main() {
	var e emp
	e.Assign("saechim", "daeki good")
	var p Printer
	p = &e
	p.Print()
}
