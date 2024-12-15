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

	l := parseText(input)
	l.print()

	fmt.Println()

	// l.moveGuard()
	// l.print()

	patrolledCells := l.getPatrolledCells()

	fmt.Println()

	loopCount := l.moveGuardUntilItLoops()

	fmt.Println()

	fmt.Printf("* First half: there are %d patrolled cells.\n", len(patrolledCells)+1)
	fmt.Printf("* Second half: there are %d possible obstruction locations.\n", loopCount)
}
