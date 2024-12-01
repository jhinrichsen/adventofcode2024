package adventofcode2023

import (
	"slices"
	"strconv"
	"strings"
)

func Day01(lines []string) uint {
	atoi := func(s string) (int, int) {
		parts := strings.Fields(s)
		l, _ := strconv.Atoi(parts[0])
		r, _ := strconv.Atoi(parts[1])
		return l, r
	}

	left, right := make([]int, len(lines)), make([]int, len(lines))
	for i := range lines {
		left[i], right[i] = atoi(lines[i])
	}
	slices.Sort(left)
	slices.Sort(right)

	delta := func(x, y int) uint {
		return uint(max(x, y) - min(x, y))
	}

	var sum uint
	for i := range lines {
		sum += delta(left[i], right[i])
	}
	return sum
}
