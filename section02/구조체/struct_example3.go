package main

import "github.com/davecgh/go-spew/spew"

type Student struct {
	name   string
	age    int
	scores []Score
}

type Score struct {
	subject string
	score   int
}

func main() {
	s1 := Student{
		name: "saechimdaeki",
		age:  29,
		scores: []Score{
			{
				subject: "math",
				score:   10,
			},
			{
				subject: "music",
				score:   90,
			},
		},
	}
	spew.Dump(s1)
}
