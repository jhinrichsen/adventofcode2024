package adventofcode2024

import (
	"strconv"
	"strings"
)

type Day25Puzzle struct {
	locks [][]int // each lock is [5]int of pin heights
	keys  [][]int // each key is [5]int of tooth heights
}

func NewDay25(lines []string) (Day25Puzzle, error) {
	var locks [][]int
	var keys [][]int

	// Parse schematics separated by empty lines
	i := 0
	for i < len(lines) {
		// Skip empty lines
		if strings.TrimSpace(lines[i]) == "" {
			i++
			continue
		}

		// Read a 7-row schematic
		schematic := make([]string, 0, 7)
		for i < len(lines) && strings.TrimSpace(lines[i]) != "" {
			schematic = append(schematic, strings.TrimSpace(lines[i]))
			i++
		}

		if len(schematic) != 7 {
			continue // Invalid schematic
		}

		// Determine if lock or key
		isLock := schematic[0] == "#####"
		heights := make([]int, 5)

		if isLock {
			// Lock: count # in each column below the top row
			for col := range 5 {
				count := 0
				for row := 1; row < 7; row++ {
					if schematic[row][col] == '#' {
						count++
					}
				}
				heights[col] = count
			}
			locks = append(locks, heights)
		} else {
			// Key: count # in each column above the bottom row
			for col := range 5 {
				count := 0
				for row := 0; row < 6; row++ {
					if schematic[row][col] == '#' {
						count++
					}
				}
				heights[col] = count
			}
			keys = append(keys, heights)
		}
	}

	return Day25Puzzle{locks: locks, keys: keys}, nil
}

func Day25(puzzle Day25Puzzle) string {
	count := 0

	// Check each lock/key pair
	for _, lock := range puzzle.locks {
		for _, key := range puzzle.keys {
			fits := true
			for col := range 5 {
				if lock[col]+key[col] > 5 {
					fits = false
					break
				}
			}
			if fits {
				count++
			}
		}
	}

	return strconv.Itoa(count)
}
