package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvHelpMessage := "a csv file in the format of 'question,answer'"
	csvFilename := flag.String("csv", "problems.csv", csvHelpMessage)
	timeLimit := flag.Int("limit", 30, "timer for the quiz")

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

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0
	// problemloop: // this is a loop label
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.question)

		answerChannel := make(chan string)

		// this will like run in the background so fmt.Scanf doesn't block
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer) // this method blocks the app flow and doesn't allow the timer to stop the call even when it has expired
			answerChannel <- answer
		}()

		// for whicherver comes first this select statement will know what to do
		select {
		case <-timer.C:
			fmt.Printf("\nThe time limit has been reached.\nYou scored %d out of %d.", correct, len(problems))
			return // we coulssd use break but we want to break out of the loop totally so we use return
			// break statement will just end this iteration and move on to the next iteration in the loop
			// break problemloop // if we don't want to use a return statememnt we can break out of the loop using the
			// break statement on the label of the loop
		case answer := <-answerChannel:
			if answer == p.answer {
				correct++
			}
		}
	}

	// if you use the break statement this line will be executed
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
