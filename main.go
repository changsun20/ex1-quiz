package main

import (
	"bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	csvPath := flag.String("csv", "", "Specify the path to the csv file of the quiz")

	flag.Parse()

	if *csvPath == "" {
		log.Fatal("Must specify the path to CSV file using -csv")
	}

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
		prompt, answerString := problem[0], problem[1]
		answer, err := strconv.Atoi(answerString)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Problem %d - %s: ", index+1, prompt)

		var userAnswer int
		_, err = fmt.Scan(&userAnswer)
		if err != nil {
			log.Fatal(err)
		}

		if userAnswer == answer {
			fmt.Println("Correct!")
			correctCount++
		} else {
			fmt.Println("Oh no, wrong answer")
		}
	}

	fmt.Printf("You got %d out of %d correct!\n", correctCount, total)
}
