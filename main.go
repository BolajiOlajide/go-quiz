package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvHelpMessage := "a csv file in the format of 'question,answer'"
	csvFilename := flag.String("csv", "problems.csv", csvHelpMessage)

	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		errorExit(fmt.Sprintf("Failed to open the csv file %s\n", *csvFilename))
	}
	// _ = file
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		errorExit("Failed to parse the provided csv file")
	}

	problems := parseLines(lines)

	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.question)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.answer {
			correct++
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]), // to ensure data is clean from the csv
		}
	}
	return ret
}

type problem struct {
	question string
	answer   string
}

func errorExit(msg string) {
	fmt.Print(msg)
	os.Exit(1)
}
