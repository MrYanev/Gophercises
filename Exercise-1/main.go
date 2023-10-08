package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "A csv file in question-answer format")
	flag.Parse()
	file, err := os.Open(*csvFileName)
	if err != nil {
		fmt.Errorf("Couldn't open the file %s\n", *csvFileName)
		os.Exit(1)
	}
	_ = file

}
