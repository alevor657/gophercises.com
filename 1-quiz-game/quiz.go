package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

var problemsFile string
var timeout uint

func main() {
	flag.StringVar(&problemsFile, "problems", "problems.csv", "Provide a path to csv file")
	flag.UintVar(&timeout, "timeout", 30, "Timeout for a quiz game")
	flag.Parse()

	// Read problems csv
	fileBytes, err := ioutil.ReadFile(problemsFile)

	// Could not read problems file
	if err != nil {
		log.Fatal(err)
	}

	content := fmt.Sprintf("%s", fileBytes)
	rows := strings.Split(content, "\n")

	correctAnswers, wrongAnswers := startQuiz(rows)

	fmt.Printf("You had [%d] corrent answers and [%d] wrong answers", correctAnswers, wrongAnswers)
}

func startQuiz(rows []string) (correctAnswers uint, wrongAnswers uint) {
	f := os.Stdin
	defer f.Close()
	scanner := bufio.NewScanner(f)
	totalQuestions := uint(len(rows))

	fmt.Printf("You have %d seconds to complete quiz with %d questions in it, press any key to continue", timeout, totalQuestions)
	scanner.Scan()

	timer := time.NewTimer(time.Duration(timeout) * time.Second)
	answerChanel := make(chan string)

	for _, row := range rows {
		questionAndAnswer := strings.Split(row, ",")
		question := questionAndAnswer[0]
		answer := questionAndAnswer[1]

		fmt.Printf("What is %s?\n", question)

		go func() {
			scanner.Scan()
			userInput := strings.TrimSpace(scanner.Text())

			answerChanel <- userInput
		}()

		select {
		case <-timer.C:
			return correctAnswers, totalQuestions - correctAnswers
		case userInput := <-answerChanel:
			if userInput == answer {
				correctAnswers++
			}
		}
	}

	return correctAnswers, totalQuestions - correctAnswers
}
