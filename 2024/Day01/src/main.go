package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
    inputFilePath := "src/resources/input.txt"
    inputFileContent, errR := os.ReadFile(inputFilePath)

    if errR != nil {
        fmt.Printf("ERROR: Failed to read from input file '%s': %s\n", inputFilePath, errR)
        os.Exit(1)
    }

    list1 := []int{};
    list2 := []int{};
    
    for _, pairPerLine := range strings.Split(string(inputFileContent), "\n") {
        if len(pairPerLine) == 0 {
            break
        }

        pairElements := strings.Split(pairPerLine, "   ")

        elem1, _ := strconv.Atoi(pairElements[0])
        elem2, _ := strconv.Atoi(pairElements[1])

        list1 = append(list1, elem1)
        list2 = append(list2, elem2)
    }

    slices.Sort(list1)
    slices.Sort(list2)

    deltaSum := 0

    for i := range len(list1) {
        delta := math.Abs(float64(list1[i] - list2[i]))
        deltaSum += int(delta)
    }

    fmt.Printf("* First half of the puzzle: sum of distances between minimum values in left and right lists is <%d>.\n", deltaSum)

    simScore := 0

    for _, number1 := range list1 {
        count := 0

        for _, number2 := range list2 {
            if number1 == number2 {
                count ++
            }
        }

        simScore += number1 * count
    }

    fmt.Printf("* Second half of the puzzle: similarity score is <%d>.\n", simScore)
}
