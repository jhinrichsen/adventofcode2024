package adventofcode2024

import (
	"cmp"
	"os"
	"slices"
)

// Day02Puzzle represents the parsed Day 2 puzzle data
type Day02Puzzle struct {
	reports [][]uint
}

// NewDay02 creates a Day02Puzzle from lines
func NewDay02(filename string) Day02Puzzle {
	buf, err := os.ReadFile(filename)
	if err != nil {
		panic(err) // This will be called from tests with proper error handling
	}
	
	var dim uint
	for i := range buf {
		if buf[i] == '\n' {
			dim++
		}
	}
	reports := make([][]uint, dim)
	for i := range reports {
		reports[i] = make([]uint, 0, 8)
	}

	isDigit := func(b byte) bool {
		return b >= '0' && b <= '9'
	}
	for i, y := 0, 0; i < len(buf); {
		var n uint

		for isDigit(buf[i]) {
			n = 10*n + uint(buf[i]-'0')
			i++
		}
		reports[y] = append(reports[y], n)

		for !isDigit(buf[i]) {
			if buf[i] == '\n' {
				y++
				i++
				break
			}
			i++
		}
	}
	
	return Day02Puzzle{reports: reports}
}

func Day02(puzzle Day02Puzzle, part1 bool) (n uint) {
	reports := puzzle.reports
	for i := range reports {
		levels := reports[i]
		if part1 {
			if safe(levels) {
				n++
			}
		} else {
			for j := range levels {
				// Delete() zero's out the last n elements from the parameter slice
				cropped := make([]uint, len(levels))
				copy(cropped, levels)
				cropped = slices.Delete(cropped, j, j+1)
				if safe(cropped) {
					n++
					break
				}
			}
		}
	}
	return
}

func safe(levels []uint) bool {
	const (
		atleast = 1
		atmost  = 3
	)
	// predicate 1: difference at least 1, at most 3
	for j := 1; j < len(levels); j++ {
		delta := max(levels[j-1], levels[j]) - min(levels[j-1], levels[j])
		if delta < atleast || delta > atmost {
			return false
		}
	}

	// predicate 2: increasing
	asc := cmp.Compare[uint]

	// predicate 3: decreasing
	desc := func(x, y uint) int {
		return cmp.Compare(y, x)
	}

	return slices.IsSortedFunc(levels, asc) || slices.IsSortedFunc(levels, desc)
}

// Diff returns number of differences between two slices.
// Unfortunately, the corresponding proposal has been rejected.
// https://github.com/golang/go/issues/58730, proposal: slices: add difference function
func Diff[T uint](s1, s2 []T) uint {
	l1, l2 := uint(len(s1)), uint(len(s2))
	var n uint
	for i := range min(l1, l2) {
		if s1[i] != s2[i] {
			n++
		}
	}
	if l1 == l2 {
		return n
	}
	return n + max(l1, l2) - min(l1, l2)
}

/*
func abs(x int) uint {
	if x < 0 {
		return -n
	}
	return n
}
*/
