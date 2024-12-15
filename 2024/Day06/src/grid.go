package main

import (
	"fmt"
	"strings"
)

const (
	reset = "\033[0m"

	bold      = "\033[1m"
	underline = "\033[4m"
	strike    = "\033[9m"
	italic    = "\033[3m"

	cRed    = "\033[31m"
	cGreen  = "\033[32m"
	cYellow = "\033[33m"
	cBlue   = "\033[34m"
	cPurple = "\033[35m"
	cCyan   = "\033[36m"
	cWhite  = "\033[37m"

	guard       = '^'
	obstacle    = '#'
	patrolled   = 'X'
	obstruction = 'O'

	moveUp    = 0
	moveRight = 1
	moveDown  = 2
	moveLeft  = 3
)

type cell rune

type row []cell

type position struct {
	x, y int
}

type lab struct {
	grid                        []row
	guardCurrPos                position
	guardCurrDir                int
	cellsWhereRotationsOccurred []position
}

func parseText(input string) lab {
	newLab := lab{guardCurrDir: moveUp}

	for i, line := range strings.Split(input, "\n") {
		newLab.grid = append(newLab.grid, make(row, 0, len(line)))

		for j, character := range line {
			newLab.grid[i] = append(newLab.grid[i], cell(character))

			if character == '^' {
				newLab.guardCurrPos = position{x: i, y: j}
			}
		}
	}

	return newLab
}

func (l *lab) print() {
	for i := range len(l.grid) {
		for j := range len(l.grid[i]) {
			char := string(l.grid[i][j])

			if char == string(guard) {
				fmt.Print(cRed + l.getGuardDisplayChar())
			} else if char == string(obstacle) {
				fmt.Print(cBlue + char)
			} else if char == string(patrolled) {
				fmt.Print(cCyan + char)
			} else {
				fmt.Print(reset + char)
			}
		}

		fmt.Println()
	}
}

func (l *lab) getGuardDisplayChar() string {
	switch l.guardCurrDir {
	case moveUp:
		return "^"
	case moveRight:
		return ">"
	case moveDown:
		return "v"
	case moveLeft:
		return "<"
	}

	return "^"
}

func (l *lab) moveGuard() {
	for {
		if exitedGrid := l.moveGuardOnce(true); exitedGrid {
			break
		}
	}
}

func (l *lab) moveGuardUntilItLoops() int {
	loopCount := 0
	initialGuardPos := position{x: l.guardCurrPos.x, y: l.guardCurrPos.y}
	initialGuardDir := l.guardCurrDir

	for i := range len(l.grid) {
		for j := range len(l.grid[i]) {
			if l.grid[i][j] == '.' {
				l.grid[i][j] = obstruction

				for {
					if exitedGrid := l.moveGuardOnce(false); exitedGrid {
						break
					} else {
						if l.isGuardStuckInLoop() {
							loopCount++
							break
						}
					}
				}

				l.grid[i][j] = '.'
				l.grid[l.guardCurrPos.x][l.guardCurrPos.y] = '.'
				l.guardCurrPos = initialGuardPos
				l.guardCurrDir = initialGuardDir
				l.cellsWhereRotationsOccurred = []position{}
			}
		}
	}

	return loopCount
}

func (l *lab) moveGuardOnce(leaveTrail bool) bool {
	switch l.guardCurrDir {
	case moveUp:
		if l.guardCurrPos.x-1 >= 0 {
			if l.grid[l.guardCurrPos.x-1][l.guardCurrPos.y].isCellWalkable() {
				if leaveTrail {
					l.grid[l.guardCurrPos.x][l.guardCurrPos.y] = patrolled
				} else {
					l.grid[l.guardCurrPos.x][l.guardCurrPos.y] = '.'
				}
				l.grid[l.guardCurrPos.x-1][l.guardCurrPos.y] = guard
				l.guardCurrPos.x--
			} else {
				l.rotateGuard()
			}
		} else {
			// Guard crossed the edge and left the grid.
			return true
		}
	case moveRight:
		if l.guardCurrPos.y+1 < len(l.grid[l.guardCurrPos.x]) {
			if l.grid[l.guardCurrPos.x][l.guardCurrPos.y+1].isCellWalkable() {
				if leaveTrail {
					l.grid[l.guardCurrPos.x][l.guardCurrPos.y] = patrolled
				} else {
					l.grid[l.guardCurrPos.x][l.guardCurrPos.y] = '.'
				}
				l.grid[l.guardCurrPos.x][l.guardCurrPos.y+1] = guard
				l.guardCurrPos.y++
			} else {
				l.rotateGuard()
			}
		} else {
			// Guard crossed the edge and left the grid.
			return true
		}
	case moveDown:
		if l.guardCurrPos.x+1 < len(l.grid) {
			if l.grid[l.guardCurrPos.x+1][l.guardCurrPos.y].isCellWalkable() {
				if leaveTrail {
					l.grid[l.guardCurrPos.x][l.guardCurrPos.y] = patrolled
				} else {
					l.grid[l.guardCurrPos.x][l.guardCurrPos.y] = '.'
				}
				l.grid[l.guardCurrPos.x+1][l.guardCurrPos.y] = guard
				l.guardCurrPos.x++
			} else {
				l.rotateGuard()
			}
		} else {
			// Guard crossed the edge and left the grid.
			return true
		}
	case moveLeft:
		if l.guardCurrPos.y-1 >= 0 {
			if l.grid[l.guardCurrPos.x][l.guardCurrPos.y-1].isCellWalkable() {
				if leaveTrail {
					l.grid[l.guardCurrPos.x][l.guardCurrPos.y] = patrolled
				} else {
					l.grid[l.guardCurrPos.x][l.guardCurrPos.y] = '.'
				}
				l.grid[l.guardCurrPos.x][l.guardCurrPos.y-1] = guard
				l.guardCurrPos.y--
			} else {
				l.rotateGuard()
			}
		} else {
			// Guard crossed the edge and left the grid.
			return true
		}
	}

	return false
}

func (c cell) isCellWalkable() bool {
	return c != obstacle && c != obstruction
}

func (l *lab) rotateGuard() {
	// "Rotate" 90 degrees to change direction.
	l.guardCurrDir = (l.guardCurrDir + 1) % 4

	l.cellsWhereRotationsOccurred = append(l.cellsWhereRotationsOccurred, l.guardCurrPos)
}

func (l *lab) getPatrolledCells() []cell {
	cells := []cell{}

	for _, row := range l.grid {
		for _, cell := range row {
			if cell == patrolled {
				cells = append(cells, cell)
			}
		}
	}

	return cells
}

func (l *lab) isGuardStuckInLoop() bool {
	frequencyMap := map[position]int{}

	for _, visitedPos := range l.cellsWhereRotationsOccurred {
		frequencyMap[visitedPos]++
	}

	for _, val := range frequencyMap {
		if val > 2 {
			return true
		}
	}

	return false
}
