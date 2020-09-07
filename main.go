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
	shuffleRecord := flag.Bool("shuffle", false, "option to shuffle")
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
	if *shuffleRecord {
		shuffle(records)
	}
	fmt.Println("Press Enter to Start the quiz")
	fmt.Scanln(&quizType)
	// for {
	// 	fmt.Scanln(&quizType)
	// 	if quizType == "1" {
	// 		shuffle(records) //Function to shuffle quiz order
	// 		break
	// 	} else if quizType == "" {
	// 		break
	// 	} else {
	// 		fmt.Println("Please enter correct input ")
	// 		continue
	// 	}
	// }

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
