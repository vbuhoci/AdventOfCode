package main

import (
	"errors"
	"strconv"
	"strings"
)

type OrderRule struct {
	Page1 int
	Page2 int
}

type PageNumbers []int

type Pages struct {
	Rules         []OrderRule
	ManualUpdates []PageNumbers
}

func ParseText(input string) Pages {
	pages := Pages{}
	parsingRules := true

	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			parsingRules = false
			continue
		}

		if parsingRules {
			rulePages := strings.Split(line, "|")

			rulePage1, _ := strconv.Atoi(rulePages[0])
			rulePage2, _ := strconv.Atoi(rulePages[1])

			orderRule := OrderRule{Page1: rulePage1, Page2: rulePage2}

			pages.Rules = append(pages.Rules, orderRule)
		} else {
			pageNumbers := strings.Split(line, ",")
			manualUpdate := PageNumbers{}

			for _, pageNumberStr := range pageNumbers {
				pageNumber, _ := strconv.Atoi(pageNumberStr)
				manualUpdate = append(manualUpdate, pageNumber)
			}

			pages.ManualUpdates = append(pages.ManualUpdates, manualUpdate)
		}
	}

	return pages
}

func (p *Pages) GetRulesCompliantNonCompliantLines() ([]PageNumbers, []PageNumbers) {
	compliantManualUpdates := []PageNumbers{}
	noncompliantManualUpdates := []PageNumbers{}

	for _, manualUpdate := range p.ManualUpdates {
		if isCompliant, _, _ := IsUpdateCompliant(p, manualUpdate); isCompliant {
			compliantManualUpdates = append(compliantManualUpdates, manualUpdate)
		} else {
			noncompliantManualUpdates = append(noncompliantManualUpdates, manualUpdate)
		}
	}

	return compliantManualUpdates, noncompliantManualUpdates
}

func IsUpdateCompliant(p *Pages, manualUpdate PageNumbers) (bool, int, int) {
	for _, rule := range p.Rules {
		pos1, err1 := getElemIndexInSlice(manualUpdate, rule.Page1)
		pos2, err2 := getElemIndexInSlice(manualUpdate, rule.Page2)

		// TODO: for debugging
		// fmt.Printf("Rule %v / [%d|%d]\n", rule, pos1, pos2)

		if err1 == nil && err2 == nil {
			if pos1 > pos2 {
				return false, pos1, pos2
			}
		}
	}

	return true, -1, -1
}

func getElemIndexInSlice(slice []int, item int) (int, error) {
	for i, elem := range slice {
		if elem == item {
			return i, nil
		}
	}

	return -1, errors.New("element not found in slice")
}
