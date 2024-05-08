package main

import (
	"fmt"
	"log"
	"os"

	reading "github.com/edoardottt/go-readingtime"
)

const (
	argsNumber = 2
)

func main() {
	if len(os.Args) != argsNumber {
		log.Fatal("Usage: readt filepath")
	}

	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	t := reading.HumanEstimate(string(file))

	fmt.Printf("It would take about %s.\n", t)
}
