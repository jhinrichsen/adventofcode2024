package adventofcode2024

import (
	"cmp"
	"slices"
)

func Day02(reports [][]uint, part1 bool) (n uint) {
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
