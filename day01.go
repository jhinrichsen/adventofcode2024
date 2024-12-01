package adventofcode2023

import (
	"slices"
	"strconv"
	"strings"
)

func Day01(lines []string, part1 bool) uint {
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

	var sum uint
	if part1 {
		slices.Sort(left)
		slices.Sort(right)

		delta := func(x, y int) uint {
			return uint(max(x, y) - min(x, y))
		}

		for i := range lines {
			sum += delta(left[i], right[i])
		}
	} else {

		histogram := make(map[int]uint, len(lines)) // len(lines) is the worst case without any  duplicates
		for i := range right {
			histogram[right[i]] += 1
		}

		for i := range left {
			sum += uint(left[i]) * histogram[left[i]]
		}
	}
	return sum
}

func twoUint(buf []byte) [][2]uint {
	return nil
}
