package adventofcode2023

import (
	"slices"
	"strconv"
	"strings"
)

func Day02(lines []string, part1 bool) uint {
	var (
		n uint
		// part2 = !part1
	)

	for i := range lines {
		parts := strings.Fields(lines[i])
		levels := make([]uint, len(parts))
		for j := range parts {
			x, _ := strconv.Atoi(parts[j])
			levels[j] = uint(x)
		}

		if part1 {
			if safe(levels) {
				n++
			}
		} else {
			for j := range levels {
				// Delete() zeroes the last n elements from the parameter slice
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
	return n
}

func safe(levels []uint) bool {
	const (
		atleast = 1
		atmost  = 3
	)
	var (
		cp         []uint
		increasing bool
		decreasing bool
	)
	// predicate 1: difference at least 1, at most 3
	for j := 1; j < len(levels); j++ {
		delta := max(levels[j-1], levels[j]) - min(levels[j-1], levels[j])
		if delta < atleast || delta > atmost {
			return false
		}
	}

	// predicate 2: increasing
	cp = make([]uint, len(levels))
	copy(cp, levels)
	slices.Sort(cp)
	increasing = slices.Equal(levels, cp)

	// predicate 3: decreasing
	slices.Reverse(cp)
	decreasing = slices.Equal(levels, cp)

	return increasing || decreasing
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
