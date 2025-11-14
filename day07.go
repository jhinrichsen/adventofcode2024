package adventofcode2024

type Day07Equation struct {
	target uint
	values [12]uint
	count  uint
}

type Day07Puzzle []Day07Equation

func NewDay07(lines []string) (Day07Puzzle, error) {
	puzzle := make(Day07Puzzle, len(lines))

	for i, line := range lines {
		var (
			want uint
			vals [12]uint
			idx  uint
			j    int
			x    uint
		)

		// parse test value
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
				vals[idx] = x
				idx++
				x = 0
				continue
			}
			x = 10*x + uint(b-'0')
		}
		// append final digit (no trailing separator)
		vals[idx] = x
		idx++

		puzzle[i] = Day07Equation{
			target: want,
			values: vals,
			count:  idx,
		}
	}

	return puzzle, nil
}

func Day07(puzzle Day07Puzzle, part1 bool) (sum uint) {
	for _, eq := range puzzle {
		want := eq.target
		vals := eq.values
		idx := eq.count

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

		// Optimized iterative approach with better pruning
		numOps := idx - 1

		// Use stack-based approach to avoid full enumeration
		type state struct {
			pos    uint
			result uint
		}

		stack := [1024]state{{pos: 0, result: vals[0]}}
		stackSize := 1

		found := false
		for stackSize > 0 && !found {
			stackSize--
			current := stack[stackSize]

			if current.pos == numOps {
				if current.result == want {
					sum += want
					found = true
				}
				continue
			}

			if current.result > want {
				continue // Prune this branch
			}

			next := vals[current.pos+1]
			nextPos := current.pos + 1

			// Try addition
			newResult := current.result + next
			if newResult <= want && stackSize < 1024 {
				stack[stackSize] = state{pos: nextPos, result: newResult}
				stackSize++
			}

			// Try multiplication
			newResult = current.result * next
			if newResult <= want && stackSize < 1024 {
				stack[stackSize] = state{pos: nextPos, result: newResult}
				stackSize++
			}

			// Try concatenation (Part 2 only)
			if !part1 {
				newResult = concat(current.result, next)
				if newResult <= want && stackSize < 1024 {
					stack[stackSize] = state{pos: nextPos, result: newResult}
					stackSize++
				}
			}
		}
	}
	return
}
