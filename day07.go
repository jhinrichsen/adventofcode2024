package adventofcode2024

func Day07(lines []string, part1 bool) (sum uint) {
	for _, line := range lines {
		// parse test value

		var (
			want uint

			vals   = [12]uint{}
			idx    uint
			append = func(x uint) {
				vals[idx] = x
				idx++
			}
			j int
			x uint
		)
		for j = range line {
			if line[j] == ':' {
				j++
				j++
				break
			}
			want = 10*want + uint(line[j]-'0')
		}

		// parse equation

		for _, b := range line[j:] {
			if b == ' ' {
				append(x)
				x = 0
				continue
			}
			x = 10*x + uint(b-'0')
		}
		// append final digit (no trailing separator)
		append(x)

		// Helper function to concatenate two numbers
		concat := func(a, b uint) uint {
			multiplier := uint(1)
			temp := b
			for temp > 0 {
				multiplier *= 10
				temp /= 10
			}
			return a*multiplier + b
		}

		// Iterative approach using base-3 counting for operator combinations
		numOps := idx - 1
		maxCombinations := uint(1)
		for i := uint(0); i < numOps; i++ {
			if part1 {
				maxCombinations *= 2 // 2 operators: +, *
			} else {
				maxCombinations *= 3 // 3 operators: +, *, ||
			}
		}

		found := false
		for combo := uint(0); combo < maxCombinations && !found; combo++ {
			result := vals[0]
			temp := combo
			
			for i := uint(0); i < numOps; i++ {
				var op uint
				if part1 {
					op = temp % 2
					temp /= 2
				} else {
					op = temp % 3
					temp /= 3
				}
				
				next := vals[i+1]
				switch op {
				case 0: // addition
					result += next
				case 1: // multiplication
					result *= next
				case 2: // concatenation (Part 2 only)
					result = concat(result, next)
				}
				
				if result > want {
					break // Early pruning
				}
			}
			
			if result == want {
				sum += want
				found = true
			}
		}
	}
	return
}
