package adventofcode2024

import (
	"slices"
)

type tuple struct {
	before, after uint8
}

func Day05(lines []string, part1 bool) (sum uint) {
	var (
		nRules int
		rules  [1176]tuple

		digit = func(b byte) uint8 {
			return b - '0'
		}

		ten = func(b byte) uint8 {
			return 10 * digit(b)
		}
	)

	// section 1: rules

	for {
		line := lines[nRules]
		if len(line) == 0 {
			break
		}
		rules[nRules].before = ten(line[0]) + digit(line[1])
		rules[nRules].after = ten(line[3]) + digit(line[4])

		nRules++
	}

	// section 2: order

	for _, line := range lines[nRules+1:] {
		var (
			numbers    [23]uint8
			nNumbers   = (len(line) + 1) / 3
			sortRules  [253]tuple
			nSortRules = 0
			ordered    = true
		)

		// parse section 2: numbers
		for j, idx := 0, 0; j < nNumbers; j++ {
			numbers[j] = ten(line[idx]) + digit(line[idx+1])
			idx += 3
		}

		for j := range numbers[:nNumbers] {
			for k := j + 1; k < nNumbers; k++ {
				current := numbers[j]
				succ := numbers[k]
				for _, rule := range rules {
					bad := current == rule.after && succ == rule.before
					if bad {
						ordered = false
						if part1 { // done
							goto l1
						}
					}

					// for part 2, harvest all matching rules
					good := current == rule.before && succ == rule.after
					if good || bad {
						sortRules[nSortRules] = rule
						nSortRules++
					}
				}
			}
		}
	l1:
		if part1 && ordered {
			middle := numbers[nNumbers/2]
			sum += uint(middle)
		}
		if !part1 && !ordered {
			slices.SortFunc(numbers[:nNumbers], func(x, y uint8) int {
				for _, rule := range sortRules {
					// ascending or descending does not really matter, because we pick the middle element at the end
					if x == rule.before && y == rule.after {
						return -1
					}
					if x == rule.after && y == rule.before {
						return 1
					}
				}
				return 0
			})
			middle := numbers[nNumbers/2]
			sum += uint(middle)
		}
	}
	return
}
