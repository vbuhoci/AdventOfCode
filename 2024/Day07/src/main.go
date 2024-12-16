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

	eqs := parseInput(input)
	fmt.Println(eqs)

	validEqs1 := filterValidEquations(eqs, OP_ADD, OP_MUL)
	validEqs2 := filterValidEquations(eqs, OP_ADD, OP_MUL, OP_CON)
	// fmt.Println(validEqs1)
	// fmt.Println(validEqs2)

	sum1, sum2 := 0, 0

	for _, validEq := range validEqs1 {
		sum1 += int(validEq.testResult)
	}

	for _, validEq := range validEqs2 {
		sum2 += int(validEq.testResult)
	}

	fmt.Printf("* First half: sum of valid equations is %d.\n", sum1)
	fmt.Printf("* Second half: sum of valid equations is %d.\n", sum2)
}
