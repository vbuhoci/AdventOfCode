package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	inputFilePath := "src/resources/input.txt"
	inputFileContent, errR := os.ReadFile(inputFilePath)

	if errR != nil {
		log.Fatalf("Failed to read input file '%s' due to: %s", inputFilePath, errR)
	}

	input := string(inputFileContent)

	// TODO: just for debugging.
	// fmt.Println(input)

	tm := parseInput(input)
	// fmt.Printf("%+v\n", tm)

	score, rating := 0, 0
	trails := tm.findTrails()

	for _, trail := range trails {
		score += len(trail.summits)
		rating += trail.distinctTracks
	}

	fmt.Printf("* First half: score is %d.\n", score)
	fmt.Printf("* Second half: rating is %d.\n", rating)
}
