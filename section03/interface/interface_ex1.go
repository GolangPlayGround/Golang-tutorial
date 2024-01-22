package main

import "github.com/davecgh/go-spew/spew"

type Employee interface {
	PrintName() string
	PrintAddress() [3]string
}

type Emp struct {
	Name, PostNumber, Address1, Address2 string
}

func (e Emp) PrintName() string {
	return e.Name
}

func (e Emp) PrintAddress() [3]string {
	return [3]string{e.PostNumber, e.Address1, e.Address2}
}

func main() {
	var e Employee = Emp{"John", "12345", "address1", "address2"}
	spew.Dump(e.PrintName(), e.PrintAddress())
}
