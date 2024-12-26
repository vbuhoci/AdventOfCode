package main

import (
	"strconv"
	"strings"
)

type position struct {
	x, y int
}

type topomap struct {
	grid           [][]int
	lines, columns int
	trailheads     []position
	summits        []position
}

type trail struct {
	start          position
	summits        []position
	distinctTracks int
}

func parseInput(input string) topomap {
	tm := topomap{}

	for i, line := range strings.Split(input, "\n") {
		tm.grid = append(tm.grid, []int{})

		for j, char := range line {
			height, err := strconv.Atoi(string(char))
			if err != nil {
				tm.grid[i] = append(tm.grid[i], 999)
			} else {
				tm.grid[i] = append(tm.grid[i], height)

				if height == 0 {
					tm.trailheads = append(tm.trailheads, position{x: i, y: j})
				} else if height == 9 {
					tm.summits = append(tm.summits, position{x: i, y: j})
				}
			}
		}

		tm.lines++
		tm.columns = len(tm.grid[i])
	}

	return tm
}

func (tm *topomap) findTrails() []trail {
	trails := []trail{}

	for _, start := range tm.trailheads {
		visitedSummits := []position{}

		tm.tryHike(start, &visitedSummits)

		trails = append(trails, trail{start: start, summits: getUniqueValues(visitedSummits), distinctTracks: len(visitedSummits)})
	}

	return trails
}

func (tm *topomap) tryHike(start position, visitedSummits *[]position) {
	if tm.grid[start.x][start.y] == 9 {
		*visitedSummits = append(*visitedSummits, start)
		return
	}

	if tm.isPlaceWalkable(start, position{x: start.x, y: start.y + 1}) {
		tm.tryHike(position{x: start.x, y: start.y + 1}, visitedSummits)
	}

	if tm.isPlaceWalkable(start, position{x: start.x + 1, y: start.y}) {
		tm.tryHike(position{x: start.x + 1, y: start.y}, visitedSummits)
	}

	if tm.isPlaceWalkable(start, position{x: start.x, y: start.y - 1}) {
		tm.tryHike(position{x: start.x, y: start.y - 1}, visitedSummits)
	}

	if tm.isPlaceWalkable(start, position{x: start.x - 1, y: start.y}) {
		tm.tryHike(position{x: start.x - 1, y: start.y}, visitedSummits)
	}
}

func (tm *topomap) isPlaceWalkable(currPlace position, potentialPlace position) bool {
	return potentialPlace.x >= 0 &&
		potentialPlace.x < tm.lines &&
		potentialPlace.y >= 0 &&
		potentialPlace.y < tm.columns &&
		tm.grid[potentialPlace.x][potentialPlace.y]-tm.grid[currPlace.x][currPlace.y] == 1
}

func getUniqueValues[T comparable](values []T) []T {
	uniquesMap := map[T]bool{}
	uniquesArr := []T{}

	for _, val := range values {
		uniquesMap[val] = true
	}

	for k := range uniquesMap {
		uniquesArr = append(uniquesArr, k)
	}

	return uniquesArr
}
