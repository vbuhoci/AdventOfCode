package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	inputFilePath := "src/resources/input.txt"
	inputFileContent, errR := os.ReadFile(inputFilePath)

	if errR != nil {
		log.Fatalf("Failed to read input file '%s' due to: %s", inputFilePath, errR)
	}

	input := string(inputFileContent)

	// TODO: just for debugging.
	// fmt.Println(input)

	fs := parseInput(input)
	// fmt.Println(fs)

	blocks := fs.getBlockArray()
	// fmt.Println(blocks)
	// printBlocks(blocks)

	compBlocks := compressBlocks(blocks)
	// printBlocks(compBlocks)

	compWholeBlocks := compressBlocksWhole(blocks)
	// printBlocks(compWholeBlocks)

	fmt.Printf("* First half: checksum is %d.\n", computeChecksum(compBlocks))
	fmt.Printf("* Second half: checksum is %d.\n", computeChecksum(compWholeBlocks))
}

func printBlocks(blocks []int) {
	for _, b := range blocks {
		if b < 0 {
			fmt.Print(".")
		} else {
			fmt.Print(b)
		}
	}

	fmt.Println()
}

func compressBlocks(blocks []int) []int {
	compressed := append([]int{}, blocks...)

	l := len(blocks)
	rlIndex := l - 1 // Right-to-left index.
	lrIndex := 0     // Left-to-right index.

	for {
		if rlIndex <= lrIndex {
			break
		}

		// Find the right-most block.
		for ; rlIndex >= 0; rlIndex-- {
			if compressed[rlIndex] > 0 {
				break
			}
		}

		// Find the left-most empty space.
		for ; lrIndex < l; lrIndex++ {
			if compressed[lrIndex] < 0 {
				break
			}
		}

		if rlIndex <= lrIndex {
			break
		}

		// Swap elements.
		compressed[lrIndex], compressed[rlIndex] = compressed[rlIndex], compressed[lrIndex]
	}

	return compressed
}

func compressBlocksWhole(blocks []int) []int {
	compressed := append([]int{}, blocks...)

	l := len(blocks)
	rlIndex := l - 1 // Right-to-left index.
	lrIndex := 0     // Left-to-right index.

	for {
		if rlIndex <= lrIndex {
			break
		}

		// Find the right-most block.
		for ; rlIndex >= 0; rlIndex-- {
			if compressed[rlIndex] > 0 {
				break
			}
		}

		// Find starting index of blocks to move to the left.
		rlIndexStart, rlIndexEnd := rlIndex, rlIndex
		for {
			rlIndexStart--

			if rlIndexStart < 0 || blocks[rlIndex] != blocks[rlIndexStart] {
				rlIndexStart += 1
				break
			}
		}

		for {
			// Find the left-most empty space.
			for ; lrIndex < l; lrIndex++ {
				if compressed[lrIndex] < 0 {
					break
				}
			}

			if rlIndex <= lrIndex {
				lrIndex = 0
				rlIndex = rlIndexStart - 1
				break
			}

			// Find ending index of empty space on the left side.
			lrIndexStart, lrIndexEnd := lrIndex, lrIndex
			for {
				lrIndexEnd++

				if lrIndexEnd >= rlIndexStart || blocks[lrIndexEnd] > 0 {
					lrIndexEnd--
					break
				}
			}

			// Ensure there is enough free space for the whole chunk of blocks to fit in.
			if lrIndexEnd-lrIndexStart < rlIndexEnd-rlIndexStart {
				if lrIndexStart >= rlIndexStart {
					rlIndex = rlIndexStart - 1
					break
				} else {
					lrIndex = lrIndexEnd + 1
					continue
				}
			}

			// Swap elements.
			for i := rlIndexStart; i <= rlIndexEnd; i++ {
				compressed[lrIndexStart], compressed[i] = compressed[i], compressed[lrIndexStart]
				lrIndexStart++
			}

			lrIndex = 0
			rlIndex = rlIndexStart - 1
			break
		}
	}

	return compressed
}

func computeChecksum(blocks []int) int64 {
	var checksum int64 = 0

	for i, val := range blocks {
		if val > 0 {
			checksum += int64(i * val)
		}
	}

	return checksum
}
