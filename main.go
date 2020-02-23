package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("./import.dat")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(f)
	var line string
	var lineCount int64
	for {
		line, err = reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		if lineCount == 0 {
			lineCount++
			continue
		}
		lineCount++
		r := csv.NewReader(strings.NewReader(string(line)))
		r.Comma = '|'
		records, err := r.ReadAll()
		if err != nil {
			panic(err)
		}
		for _, row := range records {
			fmt.Println(row)
		}
	}

	fmt.Println("Finished")

}
