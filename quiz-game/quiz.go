package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var defaultProblemsFile = "problems.csv"

func main() {
	var correctAnswers int
	var wrongAnswers int

	// Read problems csv
	fileBytes, err := ioutil.ReadFile(defaultProblemsFile)
	f := os.Stdin
	scanner := bufio.NewScanner(f)
	defer f.Close()

	// Conuld not read problems file
	if err != nil {
		log.Fatal(err)
	}

	content := fmt.Sprintf("%s", fileBytes)
	rows := strings.Split(content, "\n")

	for _, row := range rows {
		questionAndAnswer := strings.Split(row, ",")
		question := questionAndAnswer[0]
		answer := questionAndAnswer[1]

		fmt.Printf("What is %s?\n", question)

		scanner.Scan()
		userInput := string(scanner.Text())

		if userInput == answer {
			correctAnswers++
		} else {
			wrongAnswers++
		}
	}

	fmt.Printf("You had [%d] corrent answers and [%d] wrong answers", correctAnswers, wrongAnswers)
}
