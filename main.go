package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

type Problem struct {
	Question string
	Answer   string
}

func ParseProblem(record []string) (Problem, error) {
	if len(record) < 2 {
		return Problem{}, fmt.Errorf("invalid record: %v", record)
	}
	return Problem{Question: record[0], Answer: record[1]}, nil
}

func CheckAnswer(userAnswer, correctAnswer string) bool {
	return strings.TrimSpace(userAnswer) == strings.TrimSpace(correctAnswer)
}

func ReadInput(answerChannel chan<- string) {
	var userAnswer string
	_, err := fmt.Scan(&userAnswer)
	if err != nil {
		log.Fatal(err)
	}
	answerChannel <- userAnswer
}

func main() {
	csvPath := flag.String("csv", "problems.csv", "Specify the path to the csv file of the quiz")

	flag.Parse()

	file, err := os.Open(*csvPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	correctCount := 0
	problemNumber := 1

	fmt.Println("Welcome to the quiz. Now, answer the following questions.")
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		problem, err := ParseProblem(record)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Problem %d - %s: ", problemNumber, problem.Question)

		answerChannel := make(chan string, 1)

		go ReadInput(answerChannel)

		var userAnswer string
		select {
		case ans := <-answerChannel:
			userAnswer = ans
		case <-time.After(5 * time.Second):
			log.Fatal("You use too much time!")
		}

		if CheckAnswer(userAnswer, problem.Answer) {
			fmt.Println("Correct!")
			correctCount++
		} else {
			fmt.Println("Oh no, wrong answer")
		}

		problemNumber++
	}

	fmt.Printf("You got %d out of %d correct!\n", correctCount, problemNumber)
}
