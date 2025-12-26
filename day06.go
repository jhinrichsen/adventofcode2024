package adventofcode2024

// Day06Puzzle represents the parsed Day 6 puzzle data
type Day06Puzzle struct {
	grid     []byte
	dimX     int
	dimY     int
	startIdx int
}

// NewDay06 creates a Day06Puzzle from lines
func NewDay06(lines []string) (Day06Puzzle, error) {
	dimY := len(lines)
	dimX := len(lines[0])
	size := dimX * dimY

	grid := make([]byte, size)
	var startIdx int

	for y, line := range lines {
		for x := range line {
			idx := y*dimX + x
			cell := line[x]
			if cell == '^' {
				startIdx = idx
				grid[idx] = '.'
			} else {
				grid[idx] = cell
			}
		}
	}

	return Day06Puzzle{
		grid:     grid,
		dimX:     dimX,
		dimY:     dimY,
		startIdx: startIdx,
	}, nil
}

// Day06 solves Day 6 using flat arrays and index arithmetic
func Day06(p Day06Puzzle, part1 bool) uint {
	if part1 {
		return day06Part1(p)
	}
	return day06Part2(p)
}

func day06Part1(p Day06Puzzle) uint {
	dimX, dimY := p.dimX, p.dimY
	size := dimX * dimY

	// Direction offsets: up, right, down, left
	dirs := [4]int{-dimX, 1, dimX, -1}

	visited := make([]bool, size)
	pos := p.startIdx
	dir := 0 // 0=up, 1=right, 2=down, 3=left

	var count uint

	for {
		x, y := pos%dimX, pos/dimX

		if !visited[pos] {
			visited[pos] = true
			count++
		}

		// Check bounds for next position
		canMove := true
		switch dir {
		case 0: // up
			if y == 0 {
				return count
			}
		case 1: // right
			if x == dimX-1 {
				return count
			}
		case 2: // down
			if y == dimY-1 {
				return count
			}
		case 3: // left
			if x == 0 {
				return count
			}
		}

		nextPos := pos + dirs[dir]
		if p.grid[nextPos] == '#' {
			// Turn right
			dir = (dir + 1) & 3
			canMove = false
		}

		if canMove {
			pos = nextPos
		}
	}
}

func day06Part2(p Day06Puzzle) uint {
	dimX, dimY := p.dimX, p.dimY
	size := dimX * dimY

	// Direction offsets: up, right, down, left
	dirs := [4]int{-dimX, 1, dimX, -1}

	// First, trace original path to find candidate obstruction positions
	originalPath := make([]bool, size)
	{
		pos := p.startIdx
		dir := 0

		for {
			x, y := pos%dimX, pos/dimX
			originalPath[pos] = true

			switch dir {
			case 0:
				if y == 0 {
					goto done
				}
			case 1:
				if x == dimX-1 {
					goto done
				}
			case 2:
				if y == dimY-1 {
					goto done
				}
			case 3:
				if x == 0 {
					goto done
				}
			}

			nextPos := pos + dirs[dir]
			if p.grid[nextPos] == '#' {
				dir = (dir + 1) & 3
			} else {
				pos = nextPos
			}
		}
	done:
	}

	// State visited array for loop detection (reused)
	visited := make([]bool, size*4)

	// Check if adding obstruction at obstIdx causes a loop
	causesLoop := func(obstIdx int) bool {
		clear(visited)

		pos := p.startIdx
		dir := 0

		for {
			x, y := pos%dimX, pos/dimX

			stateIdx := pos*4 + dir
			if visited[stateIdx] {
				return true // Loop detected
			}
			visited[stateIdx] = true

			// Check bounds
			switch dir {
			case 0:
				if y == 0 {
					return false
				}
			case 1:
				if x == dimX-1 {
					return false
				}
			case 2:
				if y == dimY-1 {
					return false
				}
			case 3:
				if x == 0 {
					return false
				}
			}

			nextPos := pos + dirs[dir]
			if p.grid[nextPos] == '#' || nextPos == obstIdx {
				dir = (dir + 1) & 3
			} else {
				pos = nextPos
			}
		}
	}

	var count uint
	for idx := range size {
		if !originalPath[idx] || idx == p.startIdx || p.grid[idx] == '#' {
			continue
		}
		if causesLoop(idx) {
			count++
		}
	}

	return count
}
