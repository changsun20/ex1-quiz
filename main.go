package main

import (
	"bytes"
	"encoding/csv"
	"flag"
	"fmt"
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
	quiz, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	total := len(quiz)
	var correctCount int

	fmt.Println("Welcome to the quiz. Now, answer the following questions.")
	for index, problem := range quiz {
		promptString, answerString := problem[0], problem[1]

		fmt.Printf("Problem %d - %s: ", index+1, promptString)

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
	}

	fmt.Printf("You got %d out of %d correct!\n", correctCount, total)
}
