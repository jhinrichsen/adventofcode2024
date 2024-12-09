package adventofcode2024

type tuple struct {
	before, after uint8
}

func Day05(lines []string) (sum uint) {
	digit := func(b byte) uint8 {
		return b - '0'
	}
	ten := func(b byte) uint8 {
		return 10 * digit(b)
	}

	// find empty group separation line
	var section int
	for i := range lines {
		if len(lines[i]) == 0 {
			section = i
			break
		}
	}

	// parse section 1: rules

	rules := make([]tuple, section)
	for i, line := range lines[:section] {
		rules[i].before = ten(line[0]) + digit(line[1])
		rules[i].after = ten(line[3]) + digit(line[4])
	}

	var numbers [100]uint8
	for _, line := range lines[section+1:] {

		// parse section 2: numbers
		n := (len(line) + 1) / 3
		for j, idx := 0, 0; j < n; j++ {
			numbers[j] = ten(line[idx]) + digit(line[idx+1])
			idx += 3
		}

		ordered := true
		for j := range numbers[:n] {
			for k := j + 1; k < n; k++ {
				current := numbers[j]
				succ := numbers[k]
				for _, rule := range rules {
					if current == rule.after && succ == rule.before {
						ordered = false
						goto l1
					}
				}
			}
		}
	l1:
		if ordered {
			middle := numbers[n/2]
			sum += uint(middle)
		}
	}
	return
}
