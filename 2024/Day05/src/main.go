package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"reflect"
)

func main() {
	inputFilePath := "src/resources/input.txt"
	inputFileContent, errR := os.ReadFile(inputFilePath)

	if errR != nil {
		log.Fatalf("Failed to read input file '%s' due to: %s", inputFilePath, errR)
	}

	input := string(inputFileContent)

	// TODO: just for debugging.
	pages := ParseText(input)
	// fmt.Println(pages.Rules)

	compliantLines, noncompliantLines := pages.GetRulesCompliantNonCompliantLines()
	// fmt.Println(compliantLines)

	compliantSum, noncompliantSum := 0, 0

	for _, line := range compliantLines {
		compliantSum += getMiddleElemInSlice(line)
	}

	for _, line := range noncompliantLines {
		for {
			isCompliant, pos1, pos2 := IsUpdateCompliant(&pages, line)
			if !isCompliant && pos1 >= 0 && pos2 >= 0 {
				swapElementsInSlice(line, pos1, pos2)
			} else {
				break
			}
		}

		noncompliantSum += getMiddleElemInSlice(line)
	}

	fmt.Printf("* First half: sum of middle numbers in compliant manual updates is %d.\n", compliantSum)
	fmt.Printf("* Second half: sum of middle numbers in noncompliant manual updates is %d.\n", noncompliantSum)
}

func getMiddleElemInSlice(slice []int) int {
	sliceLen := len(slice)
	midElemPos := int(math.Floor(float64(sliceLen) / 2))

	return slice[midElemPos]
}

func swapElementsInSlice(slice []int, pos1 int, pos2 int) {
	swapperFunc := reflect.Swapper(slice)
	swapperFunc(pos1, pos2)
}
