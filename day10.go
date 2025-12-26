package adventofcode2024

// Day10Puzzle represents the parsed Day 10 puzzle data
type Day10Puzzle struct {
	grid [][]byte
	dimX int
	dimY int
}

// NewDay10 creates a Day10Puzzle from lines
func NewDay10(lines []string) Day10Puzzle {
	dimY := len(lines)
	grid := make([][]byte, dimY)
	for y := range grid {
		grid[y] = []byte(lines[y])
	}
	return Day10Puzzle{
		grid: grid,
		dimY: dimY,
		dimX: len(lines[0]),
	}
}

// Day10 solves Day 10 using flat array and slices instead of container/list
func Day10(puzzle Day10Puzzle, part1 bool) uint {
	dimX, dimY := puzzle.dimX, puzzle.dimY

	// Flatten grid for cache-friendly access
	flat := make([]byte, dimX*dimY)
	for y := range dimY {
		copy(flat[y*dimX:], puzzle.grid[y])
	}

	type trailState struct {
		startIdx int
		posIdx   int
	}

	// Find all trailheads (height 0)
	var current []trailState
	for idx, b := range flat {
		if b == '0' {
			current = append(current, trailState{startIdx: idx, posIdx: idx})
		}
	}

	// Direction offsets in flat array: N, E, S, W
	dirs := [4]int{-dimX, 1, dimX, -1}

	// Expand trails step by step from height 1 to 9
	for height := byte('1'); height <= '9'; height++ {
		var next []trailState
		for _, ts := range current {
			x, y := ts.posIdx%dimX, ts.posIdx/dimX
			for di, d := range dirs {
				// Bounds check based on direction
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
				ni := ts.posIdx + d
				if flat[ni] == height {
					next = append(next, trailState{startIdx: ts.startIdx, posIdx: ni})
				}
			}
		}
		current = next
	}

	// Part 2: count all complete trails
	if !part1 {
		return uint(len(current))
	}

	// Part 1: count distinct (start, end) pairs
	type pair struct{ start, end int }
	seen := make(map[pair]bool, len(current))
	for _, ts := range current {
		seen[pair{ts.startIdx, ts.posIdx}] = true
	}
	return uint(len(seen))
}
