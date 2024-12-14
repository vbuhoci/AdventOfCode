package main

import (
	"fmt"
	"os"
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

    safeReportsCount := 0

    for _, line := range strings.Split(string(inputFileContent), "\n") {
        numbers := []int{}
        numbersStr := strings.Split(line, " ")
        
        for _, numStr := range numbersStr {
            number, _ := strconv.Atoi(numStr)
            numbers = append(numbers, number)
        }

        if isReportSafe(numbers) {
            safeReportsCount++
        }
    }

    fmt.Println(isReportSafe([]int{62, 60, 63, 65, 66, 68, 71, 77}))
    fmt.Println(isReportSafe([]int{54, 53, 54, 55, 56}))

    fmt.Printf("* First half of the puzzle: amount of safe reports <%d>.\n", safeReportsCount)
}

/*
  Conditions:
    - The levels (numbers) are either all increasing or all decreasing.
    - Any two adjacent levels differ by at least one and at most three.
  Note: we do tolerate one bad level.
 */
func isReportSafe(report []int) bool {
    isSafe := isReportSequenceSafe(report)

    if isSafe {
        return true
    }

    for i := range len(report) {
        subreport := removeIndex(report, i)
        isSafe = isReportSequenceSafe(subreport)

        if isSafe {
            return true
        }
    }

    return false
}

func isReportSequenceSafe(subreport []int) bool {
    lastDelta := 0

    for i := 0; i < len(subreport) - 1; i++ {
        newDelta := subreport[i] - subreport[i+1]

        // Check if numbers are all in same order (either decreasing or increasing).
        if (lastDelta < 0 && newDelta > 0) || (lastDelta > 0 && newDelta < 0) {
            return false
        }

        // Check if delta belongs to [1, 3]; note that it can be negative, so abs() it.
        posNewDelta := abs(newDelta)
        if posNewDelta < 1 || posNewDelta > 3 {
            return false
        }

        // Prepare for next cycle.
        lastDelta = newDelta
    }

    return true
}

func abs(n int) int {
    if n >= 0 {
        return n
    } else {
        return n * (-1)
    }
}

func removeIndex(arr []int, index int) []int {
    ret := make([]int, 0, len(arr) - 1)
    ret = append(ret, arr[:index]...)
    ret = append(ret, arr[index+1:]...)

    return ret
}
