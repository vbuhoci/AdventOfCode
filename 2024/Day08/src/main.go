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

	g := parseInput(input)
	u1 := g.getUniqueAntinodes()
	u2 := g.getUniqueAntinodesWithHarmonics()

	fmt.Printf("* First half: number of antinodes is %d.\n", len(u1))
	fmt.Printf("* Second half: number of antinodes is %d.\n", len(u2))
}
