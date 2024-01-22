package main

import "github.com/davecgh/go-spew/spew"

func main() {
	var vary interface{}
	vary = 123
	spew.Dump(vary)
	vary = "good morning"
	spew.Dump(vary)
	vary = map[string]int{"saechim": 10, "daeki": 29}
	spew.Dump(vary)
	vary = [3]string{"korea", "japan", "china"}
	spew.Dump(vary)

}
