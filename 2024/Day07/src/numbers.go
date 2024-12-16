package main

import (
	"strconv"
	"strings"
)

type equation struct {
	operands   []int64
	testResult int64
}

type operator rune

const (
	OP_ADD operator = '+'
	OP_MUL operator = '*'
	OP_CON operator = '|'
)

func parseInput(input string) []equation {
	equations := []equation{}

	for _, line := range strings.Split(input, "\n") {
		eqComponents := strings.Split(line, ":")
		eqTestResult, _ := strconv.ParseInt(strings.TrimSpace(eqComponents[0]), 10, 64)
		eqOperandsStr := strings.Split(strings.TrimSpace(eqComponents[1]), " ")
		eqOperands := []int64{}

		for _, eqOpStr := range eqOperandsStr {
			eqOp, _ := strconv.ParseInt(eqOpStr, 10, 64)
			eqOperands = append(eqOperands, eqOp)
		}

		eq := equation{operands: eqOperands, testResult: eqTestResult}
		equations = append(equations, eq)
	}

	return equations
}

func filterValidEquations(eqs []equation, operators ...operator) []equation {
	validEqs := []equation{}

	for _, eq := range eqs {
		if isEqValid(eq, operators...) {
			validEqs = append(validEqs, eq)
		}
	}

	return validEqs
}

func isEqValid(eq equation, operators ...operator) bool {
	// NOTE: Possible operators are "+", "*" and "|" (concat).

	allOps := generatePermutationsOf2(len(eq.operands)-1, operators...)

	for _, ops := range allOps {
		actualResult := performEquation(eq.operands, ops)
		if eq.testResult == actualResult {
			return true
		}
	}

	return false
}

func generatePermutationsOf2(n int, ops ...operator) [][]operator {
	return permutation(n, ops)
}

func permutation(n int, ops []operator) [][]operator {
	perms := [][]operator{}

	if n == 1 {
		for _, op := range ops {
			perms = append(perms, []operator{op})
		}
	} else {
		for _, op := range ops {
			for _, perm := range permutation(n-1, ops) {
				perms = append(perms, append([]operator{op}, perm...))
			}
		}
	}

	return perms
}

func performEquation(operands []int64, operators []operator) int64 {
	var result int64 = operands[0]

	for i, op := range operators {
		switch op {
		case OP_ADD:
			result += operands[i+1]
		case OP_MUL:
			result *= operands[i+1]
		case OP_CON:
			resultAsStr := strconv.FormatInt(result, 10) + strconv.FormatInt(operands[i+1], 10)
			result, _ = strconv.ParseInt(resultAsStr, 10, 64)
		}
	}

	return result
}
