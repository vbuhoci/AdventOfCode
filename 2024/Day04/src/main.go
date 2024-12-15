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
	gr := buildGrid(input)

	// TODO: just for debugging.
	gr.print()

	foundXmasWords, foundXMasWords := 0, 0

	for i, currRow := range gr {
		for j := range currRow {
			if count := beginsXmasWord(gr, i, j); count > 0 {
				foundXmasWords += count
				fmt.Printf("Found %d XMAS word(s) starting at pos [%d][%d] (zero-indexed)\n", count, i, j)
			}

			if beginsXMasWord(gr, i, j) {
				foundXMasWords++
				fmt.Printf("Found X-MAS starting at pos [%d][%d] (zero-indexed)\n", i, j)
			}
		}
	}

	fmt.Printf("* First half: found %d XMAS words.\n", foundXmasWords)
	fmt.Printf("* Second half: found %d X-MAS words.\n", foundXMasWords)
}

func beginsXmasWord(gr grid, charRow int, charCol int) int {
	count := 0

	// We accept words written both normally and backwards, so XMAS and SAMX are valid!

	i, j := charRow, charCol

	// Test horizontally, left to right:
	if j+3 < len(gr[i]) && isXmasWord(gr[i][j], gr[i][j+1], gr[i][j+2], gr[i][j+3]) {
		count++
	}

	// Test horizontally, right to left:
	if j-3 >= 0 && isXmasWord(gr[i][j], gr[i][j-1], gr[i][j-2], gr[i][j-3]) {
		count++
	}

	// Test vertically, top to bottom:
	if i+3 < len(gr) && isXmasWord(gr[i][j], gr[i+1][j], gr[i+2][j], gr[i+3][j]) {
		count++
	}

	// Test vertically, bottom to top:
	if i-3 >= 0 && isXmasWord(gr[i][j], gr[i-1][j], gr[i-2][j], gr[i-3][j]) {
		count++
	}

	// Test diagonally, top-left to bottom-right:
	if i+3 < len(gr) && j+3 < len(gr[i]) && isXmasWord(gr[i][j], gr[i+1][j+1], gr[i+2][j+2], gr[i+3][j+3]) {
		count++
	}

	// Test diagonally, bottom-left to top-right:
	if i-3 >= 0 && j+3 < len(gr[i]) && isXmasWord(gr[i][j], gr[i-1][j+1], gr[i-2][j+2], gr[i-3][j+3]) {
		count++
	}

	// Test diagonally, top-right to bottom-left:
	if i+3 < len(gr) && j-3 >= 0 && isXmasWord(gr[i][j], gr[i+1][j-1], gr[i+2][j-2], gr[i+3][j-3]) {
		count++
	}

	// Test diagonally, bottom-right to top-left:
	if i-3 >= 0 && j-3 >= 0 && isXmasWord(gr[i][j], gr[i-1][j-1], gr[i-2][j-2], gr[i-3][j-3]) {
		count++
	}

	return count
}

func isXmasWord(c1 rune, c2 rune, c3 rune, c4 rune) bool {
	if c1 == 'X' && c2 == 'M' && c3 == 'A' && c4 == 'S' {
		return true
	}

	return false
}

func beginsXMasWord(gr grid, charRow int, charCol int) bool {
	// Same function as before, but with a twist: we look for words crossing diagonally, forming an X.

	// We accept words written both normally and backwards, so, for instance, MAS and SAM forming an X is valid!

	i, j := charRow, charCol

	// Test top-left to bottom-right:
	if i+2 < len(gr) && j+2 < len(gr[i]) &&
		isMasOrSamWord(gr[i][j], gr[i+1][j+1], gr[i+2][j+2]) &&
		isMasOrSamWord(gr[i][j+2], gr[i+1][j+1], gr[i+2][j]) {
		return true
	}

	return false
}

func isMasOrSamWord(c1 rune, c2 rune, c3 rune) bool {
	if c1 == 'M' && c2 == 'A' && c3 == 'S' {
		return true
	}

	if c1 == 'S' && c2 == 'A' && c3 == 'M' {
		return true
	}

	return false
}
