package main

import (
	"flag"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"runtime"
	"time"
)

func main() {
	envPtr := flag.String("env", "prod", "환경값(dev or prod)을 입력")
	cpuPtr := flag.Int("cpu", runtime.NumCPU(), "CPU수")
	flag.Parse()

	if *envPtr != "prod" {
		log.Logger = log.Output(zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		})
	}

	if *cpuPtr > runtime.NumCPU() {
		*cpuPtr = runtime.NumCPU()
	}

	if *cpuPtr > 0 {
		runtime.GOMAXPROCS(*cpuPtr)
	}

	log.Info().
		Str("env", *envPtr).
		Int("cpu", *cpuPtr).
		Msg("프로그램이 시작되었습니다")
}
