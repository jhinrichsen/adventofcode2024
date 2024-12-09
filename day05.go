package adventofcode2024

type tuple struct {
	before, after uint8
}

func Day05(lines []string) (sum uint) {
	const N = 1200

	var (
		nRules  int
		rules   [N]tuple
		numbers [N]uint8

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
