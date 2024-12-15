package main

import "fmt"

type row []rune
type grid []row

func buildGrid(text string) grid {
	gr := grid{}
	gr = append(gr, row{})

	i := 0

	for k, char := range text {
		if char == '\n' {
			if k == len(text)-1 {
				break
			}

			i++
			gr = append(gr, row{})
		} else {
			gr[i] = append(gr[i], char)
		}
	}

	return gr
}

func (gr grid) print() {
	i, j := 0, 0

	for i = 0; i < len(gr); i++ {
		for j = 0; j < len(gr[i]); j++ {
			fmt.Print(string(gr[i][j]))
		}

		fmt.Println()
	}
}
