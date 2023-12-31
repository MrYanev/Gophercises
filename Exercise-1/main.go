package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type Problem struct {
	Question string
	Answer   string
}

func main() {
	csvFileName := flag.String("csv", "problems.csv", "A csv file in question-answer format")
	flag.Parse()
	timeLimit := flag.Int("limit", 30, "The time limit for the quiz in seconds")
	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Couldn't open the file %s\n", *csvFileName))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provide csv")
	}

	problems := parseLines(lines)

	//Setting up the timer for the quiz
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0
problemloop:
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.Question)
		ansChan := make(chan string)
		go func() {
			var ans string
			fmt.Scanf("%s\n", &ans)
			ansChan <- ans
		}()
		select {
		//A channel to keep track of time and interrup if needed

		case <-timer.C:
			fmt.Println("\nYou ran out of time")
			break problemloop
		case ans := <-ansChan:
			if ans == p.Answer {
				fmt.Println("Correct")
				correct++
			} else {
				fmt.Println("WRONG!")
			}
		}
	}
	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

// A function to parse the lines to the Problem struct
func parseLines(lines [][]string) []Problem {
	ret := make([]Problem, len(lines))
	for i, line := range lines {
		ret[i] = Problem{
			Question: line[0],
			Answer:   strings.TrimSpace(line[1]),
		}
	}
	return ret
}

// A function to exit and print out the message
// It will be used for error handling
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
