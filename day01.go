package adventofcode2024

import (
	"slices"
)

func Day01(buf []byte, part1 bool) uint {
	list := twoUints(buf)

	var sum uint
	if part1 {
		slices.Sort(list[0])
		slices.Sort(list[1])

		delta := func(x, y uint) uint {
			return uint(max(x, y) - min(x, y))
		}

		for i := range list[0] {
			sum += delta(list[0][i], list[1][i])
		}
	} else {
		histogram := make(map[uint]uint, len(list[1])) // len() is the worst case without any  duplicates
		for i := range list[1] {
			histogram[list[1][i]] += 1
		}

		for i := range list[0] {
			sum += uint(list[0][i]) * histogram[list[0][i]]
		}
	}
	return sum
}

func twoUints(buf []byte) [2][]uint {
	var dim uint
	for i := range buf {
		if buf[i] == '\n' {
			dim++
		}
	}
	is := [...][]uint{
		make([]uint, dim),
		make([]uint, dim),
	}

	for i, idx := 0, 0; i < len(buf); i++ {
		var n uint

		// left number
		for buf[i] != ' ' {
			n = 10*n + uint(buf[i]-'0')
			i++
		}
		is[0][idx] = n

		// 3 spaces as field separator
		i += 3

		// right number
		n = 0
		for buf[i] != '\n' {
			n = 10*n + uint(buf[i]-'0')
			i++
		}
		is[1][idx] = n
		idx++
	}

	return is
}
