package main

import (
	"flag"
	"github.com/davecgh/go-spew/spew"
	"runtime"
)

func main() {
	envPtr := flag.String("env", "prod", "환경값(dev or prod)을 입력")
	cpuPtr := flag.Int("cpu", runtime.NumCPU(), "CPU수")
	flag.Parse()

	if *envPtr != "prod" {

	}

	if *cpuPtr > runtime.NumCPU() {
		*cpuPtr = runtime.NumCPU()
	}

	if *cpuPtr > 0 {
		runtime.GOMAXPROCS(*cpuPtr)
	}

	spew.Dump(envPtr)
	spew.Dump(cpuPtr)
}
