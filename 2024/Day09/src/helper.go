package main

import (
	"strconv"
)

type file struct {
	id        int
	blockSize int
}

type filesystem []file

func parseInput(input string) filesystem {
	fs := filesystem{}
	id := 0

	for i, char := range input {
		var f file
		var fId int

		fSize, _ := strconv.Atoi(string(char))

		if i%2 == 0 {
			fId = id
			id++
		} else {
			fId = -1
		}

		f = file{id: fId, blockSize: fSize}
		fs = append(fs, f)
	}

	return fs
}

func (fs filesystem) getBlockArray() []int {
	blocks := []int{}

	for _, f := range fs {
		for range f.blockSize {
			blocks = append(blocks, f.id)
		}
	}

	return blocks
}
