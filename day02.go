package adventofcode2023

import (
	"slices"
	"strconv"
	"strings"
)

func Day02(lines []string) uint {
	const (
		atleast = 1
		atmost  = 3
	)
	var (
		n          uint
		increasing bool
		decreasing bool
		cp         []uint
	)

	for i := range lines {
		parts := strings.Fields(lines[i])
		levels := make([]uint, len(parts))
		for j := range parts {
			x, _ := strconv.Atoi(parts[j])
			levels[j] = uint(x)
		}

		// predicate 1: difference at least 1, at most 3
		for j := 1; j < len(levels); j++ {
			delta := max(levels[j-1], levels[j]) - min(levels[j-1], levels[j])
			if delta < atleast || delta > atmost {
				// take a shortcut w/o predicates
				goto skip
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

		if increasing || decreasing {
			n++
		}
	skip:
	}
	return n
}

/*
func abs(x int) uint {
	if x < 0 {
		return -n
	}
	return n
}
*/
