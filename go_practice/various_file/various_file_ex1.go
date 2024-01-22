package main

import (
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"os"
)

type Employees struct {
	Employees []Employee `json:"employees"`
}

type Employee struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	data1, err := os.ReadFile("test.json")
	if err != nil {
		panic(err)
	}
	emp := Employees{}
	err = json.Unmarshal(data1, &emp)
	if err != nil {
		panic(err)
	}
	spew.Dump(emp)
	data2, err := json.MarshalIndent(emp, "", "	")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("test2.json", data2, 0644)
	if err != nil {
		panic(err)
	}
}
