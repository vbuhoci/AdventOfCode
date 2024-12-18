package main

import (
	"fmt"
	"strings"
)

type frequency rune

type position struct {
	x, y int
}

type grid struct {
	// Map/grid size.
	lines, columns int
	// For each frequency we can have multiple antennas at different locations.
	antennas map[frequency][]position
	// Antinode locations.
	antinodes  []position
	antinodesH []position
}

func parseInput(input string) grid {
	g := grid{lines: 0, columns: 0, antennas: map[frequency][]position{}}

	for i, line := range strings.Split(input, "\n") {
		for j, item := range line {
			if item != '.' {
				// We have an antenna of a certain frequency.

				f := frequency(item)
				g.antennas[f] = append(g.antennas[f], position{x: i, y: j})
			}
		}

		g.lines++
		g.columns = len(line)
	}

	// For each frequency, gather and store the antinodes of that group of antennas.
	for _, antennaPositions := range g.antennas {
		g.antinodes = append(g.antinodes, computeAntinodes(g.lines, g.columns, antennaPositions)...)
		g.antinodesH = append(g.antinodesH, computeAntinodesWithHarmonics(g.lines, g.columns, antennaPositions)...)
	}

	return g
}

/*
For a slice of antenna positions, get a slice of antinode positions.
*/
func computeAntinodes(gLines, gCols int, anteps []position) []position {
	antips := []position{}
	l := len(anteps)

	// Pair every antenna with each other to find their specific antinode.

	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if i != j {
				// Skip this cycle so we don't work with a single antenna..

				deltaX := anteps[j].x - anteps[i].x
				deltaY := anteps[j].y - anteps[i].y

				ax := anteps[i].x + deltaX*2
				ay := anteps[i].y + deltaY*2

				// Ensure this antinode doesn't fall outside the grid.
				if ax >= 0 && ax < gLines && ay >= 0 && ay < gCols {
					antips = append(antips, position{x: ax, y: ay})
				}
			}
		}
	}

	return antips
}

/*
For a slice of antenna positions, get a slice of antinode positions.
*/
func computeAntinodesWithHarmonics(gLines, gCols int, anteps []position) []position {
	antips := []position{}
	l := len(anteps)

	// Pair every antenna with each other to find their specific antinodes.

	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if i != j {
				// Skip this cycle so we don't work with a single antenna..

				deltaX := anteps[j].x - anteps[i].x
				deltaY := anteps[j].y - anteps[i].y

				ax := anteps[i].x
				ay := anteps[i].y

				for {
					ax += deltaX
					ay += deltaY

					// Ensure this antinode doesn't fall outside the grid.
					if ax >= 0 && ax < gLines && ay >= 0 && ay < gCols {
						antips = append(antips, position{x: ax, y: ay})
					} else {
						break
					}
				}
			}
		}
	}

	return antips
}

func (g *grid) getUniqueAntinodes() []position {
	return filterUniqueValues(g.antinodes)
}

func (g *grid) getUniqueAntinodesWithHarmonics() []position {
	return filterUniqueValues(g.antinodesH)
}

func filterUniqueValues(slice []position) []position {
	uniques := map[position]bool{}
	for _, pos := range slice {
		uniques[pos] = true
	}

	fmt.Println(uniques)

	finalArr := make([]position, 0, len(uniques))
	for pos, _ := range uniques {
		finalArr = append(finalArr, pos)
	}

	return finalArr
}
