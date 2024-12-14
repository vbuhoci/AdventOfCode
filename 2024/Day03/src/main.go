package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
    inputFilePath := "src/resources/input.txt"
    inputFileContent, errR := os.ReadFile(inputFilePath)

    if errR != nil {
        fmt.Printf("ERROR: Failed to read from input file '%s': %s\n", inputFilePath, errR)
        os.Exit(1)
    }

    fmt.Printf("* First half of the puzzle: sum of multiplications <%d>.\n", firstHalf(string(inputFileContent)))
    fmt.Printf("* Second half of the puzzle: sum of multiplications <%d>.\n", secondHalf(string(inputFileContent)))
}

func firstHalf(inputContent string) int {
    sum := 0

    mulOperationsRegex := regexp.MustCompile(`mul\(\d+,\d+\)`)
    mulOperandsRegex := regexp.MustCompile(`\d+`)

    mulOperationsStr := mulOperationsRegex.FindAllString(inputContent, -1)

    for _, mulOpStr := range mulOperationsStr {
        sum += execMulOperation(mulOperandsRegex, mulOpStr)
    }

    return sum
}

func secondHalf(inputContent string) int {
    sum := 0
    canMul := true

    mulOperationsRegex := regexp.MustCompile(`(mul\(\d+,\d+\))|(do\(\))|(don't\(\))`)
    mulOperandsRegex := regexp.MustCompile(`\d+`)

    mulOperationsStr := mulOperationsRegex.FindAllString(inputContent, -1)

    for _, mulOpStr := range mulOperationsStr {
        if mulOpStr == "do()" {
            canMul = true
        } else if mulOpStr == "don't()" {
            canMul = false
        } else if canMul {
            sum += execMulOperation(mulOperandsRegex, mulOpStr)
        }
    }

    return sum
}

func execMulOperation(operandsRegex *regexp.Regexp, mulOp string) int {
    operands := operandsRegex.FindAllString(mulOp, -1)
    leftOperand, _ := strconv.Atoi(operands[0])
    rightOperand, _ := strconv.Atoi(operands[1])

    return leftOperand * rightOperand
}
