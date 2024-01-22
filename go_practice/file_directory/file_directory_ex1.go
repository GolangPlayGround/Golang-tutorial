package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	const filename = "test.txt"
	fp, err := os.Create(filename)
	if err != nil {
		log.Panic().Err(err).Msg("파일 오픈 실패")
	}
	defer func() {
		if err := fp.Close(); err != nil {
			log.Panic().Err(err).Msg("파일 닫기 실패")
		}
	}()
	_, err = fmt.Fprintf(fp, "hello world\n")
	if err != nil {
		log.Panic().Err(err).Msg("파일 기록 실패")
	}

	contents, err := os.ReadFile(filename)
	if err != nil {
		log.Panic().Err(err).Msg("파일 내용 읽어오기")
	}

	spew.Dump(contents)
	fmt.Println(string(contents))
}
