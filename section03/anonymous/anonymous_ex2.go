package main

import (
	"github.com/davecgh/go-spew/spew"
	"sort"
)

func main() {
	people := []struct {
		Name string
		Age  int
	}{
		{"kim", 27},
		{"lee", 43},
		{"jung", 36},
	}
	sort.Slice(people, func(i, j int) bool { return people[i].Name < people[j].Name })
	spew.Dump(people)
}
