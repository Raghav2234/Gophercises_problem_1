package main

import (
	"fmt"
	"math/rand"
	"time"
)

func quiz(records [][]string, correct *int, c chan bool) {
	for i := 0; i < len(records); i++ {
		question, answer := records[i][0], records[i][1]
		var inputAnswer string
		fmt.Println("Question no", i+1, " ", question)
		fmt.Scanln(&inputAnswer)
		if inputAnswer == answer {
			(*correct)++
		}
	}
	c <- true
}

func shuffle(records [][]string) {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range records {
		newPosition := r.Intn(len(records) - 1)
		records[i], records[newPosition] = records[newPosition], records[i]
	}
}
