package main

import (
	"bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	csvPath := flag.String("csv", "problems.csv", "Specify the path to the csv file of the quiz")

	flag.Parse()

	csvData, err := os.ReadFile(*csvPath)
	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(bytes.NewReader(csvData))
	correctCount := 0
	problemNumber := 1

	fmt.Println("Welcome to the quiz. Now, answer the following questions.")
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		problemString, answerString := line[0], line[1]

		fmt.Printf("Problem %d - %s: ", problemNumber, problemString)

		var userAnswer string
		_, err = fmt.Scan(&userAnswer)
		if err != nil {
			log.Fatal(err)
		}

		if userAnswer == answerString {
			fmt.Println("Correct!")
			correctCount++
		} else {
			fmt.Println("Oh no, wrong answer")
		}

		problemNumber++
	}

	fmt.Printf("You got %d out of %d correct!\n", correctCount, problemNumber)
}
