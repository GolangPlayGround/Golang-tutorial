package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	rfp, err := os.Open("test.csv")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := rfp.Close(); err != nil {
			panic(err)
		}
	}()

	wfp, err := os.Create("test2.csv")
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := wfp.Close(); err != nil {
			panic(err)
		}
	}()

	reader := csv.NewReader(bufio.NewReader(rfp))
	writer := csv.NewWriter(bufio.NewWriter(wfp))

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		sum := 0.0
		for i := 1; i < len(row); i++ {
			if score, err := strconv.Atoi(row[i]); err == nil {
				sum += float64(score)
			}
		}
		avg := sum / 3.0
		row = append(row, strconv.FormatFloat(avg, 'f', 2, 64))
		fmt.Println(row)
		err = writer.Write(row)
		if err != nil {
			panic(err)
		}
		writer.Flush()
	}
}
