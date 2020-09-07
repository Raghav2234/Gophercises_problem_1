package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	// Setting Flags
	fileName := flag.String("file", "problem1.csv", "name of the file")
	timeL := flag.Int("timeLimit", 30, "an int")
	flag.Parse()

	//I/O operations
	file, err := os.Open(*fileName)
	if err != nil {
		log.Fatalln("File not found")
	}
	r := csv.NewReader(file)
	records, err := r.ReadAll()

	if err != nil {
		fmt.Println("Error occured Terminating ...")
		os.Exit(1)
	}
	var quizType string

	fmt.Println("Press 1 to randomize the questions and start or Press Enter to Start the quiz")

	for {
		fmt.Scanln(&quizType)
		if quizType == "1" {
			shuffle(records) //Function to shuffle quiz order
			break
		} else if quizType == "" {
			break
		} else {
			fmt.Println("Please enter correct input ")
			continue
		}
	}

	c := make(chan bool)

	correct := 0

	go quiz(records, &correct, c)

loop:
	for {
		select {
		case <-time.After(time.Duration(*timeL) * time.Second):
			break loop
		case <-c:
			break loop

		}
	}

	fmt.Println("You have answered", correct, "correct out of", len(records), "questions")

}
