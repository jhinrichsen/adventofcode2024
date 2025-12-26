package adventofcode2024

// Day20Puzzle represents the parsed Day 20 puzzle data
type Day20Puzzle struct {
	grid     []byte
	dimX     int
	dimY     int
	startIdx int
	endIdx   int
}

// NewDay20 creates a Day20Puzzle from lines
func NewDay20(lines []string) (Day20Puzzle, error) {
	dimY := len(lines)
	dimX := len(lines[0])
	size := dimX * dimY

	grid := make([]byte, size)
	var startIdx, endIdx int

	for y, line := range lines {
		for x := range line {
			idx := y*dimX + x
			cell := line[x]
			if cell == 'S' {
				startIdx = idx
				grid[idx] = '.'
			} else if cell == 'E' {
				endIdx = idx
				grid[idx] = '.'
			} else {
				grid[idx] = cell
			}
		}
	}

	return Day20Puzzle{
		grid:     grid,
		dimX:     dimX,
		dimY:     dimY,
		startIdx: startIdx,
		endIdx:   endIdx,
	}, nil
}

// Day20 solves Day 20 using flat arrays and index arithmetic
func Day20(p Day20Puzzle, part1 bool) uint {
	if part1 {
		return countCheats(p, 100, 2)
	}
	return countCheats(p, 100, 20)
}

func countCheats(p Day20Puzzle, minSaving uint, maxCheatDist int) uint {
	dimX, dimY := p.dimX, p.dimY
	size := dimX * dimY

	// Calculate distances from start and end
	distFromStart := bfsDistances(p.grid, dimX, dimY, p.startIdx)
	distToEnd := bfsDistances(p.grid, dimX, dimY, p.endIdx)

	normalTime := distFromStart[p.endIdx]
	if normalTime == ^uint(0) {
		return 0
	}

	var count uint

	// Iterate all non-wall cells
	for idx := range size {
		if p.grid[idx] == '#' {
			continue
		}

		x, y := idx%dimX, idx/dimX
		startDist := distFromStart[idx]
		if startDist == ^uint(0) {
			continue
		}

		// Try all possible cheats within manhattan distance
		for dy := -maxCheatDist; dy <= maxCheatDist; dy++ {
			newY := y + dy
			if newY < 0 || newY >= dimY {
				continue
			}

			absdy := dy
			if absdy < 0 {
				absdy = -absdy
			}
			maxDx := maxCheatDist - absdy

			for dx := -maxDx; dx <= maxDx; dx++ {
				newX := x + dx
				if newX < 0 || newX >= dimX {
					continue
				}

				cheatDist := absdy
				if dx < 0 {
					cheatDist += -dx
				} else {
					cheatDist += dx
				}

				if cheatDist == 0 {
					continue
				}

				endIdx := newY*dimX + newX
				if p.grid[endIdx] == '#' {
					continue
				}

				endDist := distToEnd[endIdx]
				if endDist == ^uint(0) {
					continue
				}

				cheatTime := startDist + uint(cheatDist) + endDist
				if cheatTime < normalTime && (normalTime-cheatTime) >= minSaving {
					count++
				}
			}
		}
	}

	return count
}

func bfsDistances(grid []byte, dimX, dimY, startIdx int) []uint {
	size := dimX * dimY

	// Flat distance array
	dist := make([]uint, size)
	for i := range dist {
		dist[i] = ^uint(0)
	}
	dist[startIdx] = 0

	// Direction offsets: N, E, S, W
	dirs := [4]int{-dimX, 1, dimX, -1}

	// Ring buffer queue
	queue := make([]int, 0, size/4)
	queue = append(queue, startIdx)
	head := 0

	for head < len(queue) {
		cur := queue[head]
		head++

		curDist := dist[cur]
		x, y := cur%dimX, cur/dimX

		for di, d := range dirs {
			// Bounds check
			switch di {
			case 0: // N
				if y == 0 {
					continue
				}
			case 1: // E
				if x == dimX-1 {
					continue
				}
			case 2: // S
				if y == dimY-1 {
					continue
				}
			case 3: // W
				if x == 0 {
					continue
				}
			}

			ni := cur + d
			if grid[ni] == '#' {
				continue
			}

			newDist := curDist + 1
			if newDist < dist[ni] {
				dist[ni] = newDist
				queue = append(queue, ni)
			}
		}
	}

	return dist
}

// countCheatsBySavings returns a map of savings to count (for testing)
func countCheatsBySavings(p Day20Puzzle, maxCheatDist int) map[uint]uint {
	dimX, dimY := p.dimX, p.dimY
	size := dimX * dimY

	distFromStart := bfsDistances(p.grid, dimX, dimY, p.startIdx)
	distToEnd := bfsDistances(p.grid, dimX, dimY, p.endIdx)

	normalTime := distFromStart[p.endIdx]
	if normalTime == ^uint(0) {
		return nil
	}

	savings := make(map[uint]uint)

	for idx := range size {
		if p.grid[idx] == '#' {
			continue
		}

		x, y := idx%dimX, idx/dimX
		startDist := distFromStart[idx]
		if startDist == ^uint(0) {
			continue
		}

		for dy := -maxCheatDist; dy <= maxCheatDist; dy++ {
			newY := y + dy
			if newY < 0 || newY >= dimY {
				continue
			}

			absdy := dy
			if absdy < 0 {
				absdy = -absdy
			}
			maxDx := maxCheatDist - absdy

			for dx := -maxDx; dx <= maxDx; dx++ {
				newX := x + dx
				if newX < 0 || newX >= dimX {
					continue
				}

				cheatDist := absdy
				if dx < 0 {
					cheatDist += -dx
				} else {
					cheatDist += dx
				}

				if cheatDist == 0 {
					continue
				}

				endIdx := newY*dimX + newX
				if p.grid[endIdx] == '#' {
					continue
				}

				endDist := distToEnd[endIdx]
				if endDist == ^uint(0) {
					continue
				}

				cheatTime := startDist + uint(cheatDist) + endDist
				if cheatTime < normalTime {
					timeSaved := normalTime - cheatTime
					savings[timeSaved]++
				}
			}
		}
	}

	return savings
}
