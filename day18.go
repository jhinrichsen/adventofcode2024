package adventofcode2024

import (
	"fmt"
	"strconv"
	"strings"
)

// Day18Puzzle represents the parsed Day 18 puzzle data
type Day18Puzzle struct {
	points     []int // flat indices
	dimX, dimY int
}

// NewDay18 creates a Day18Puzzle from lines
func NewDay18(lines []string, dimX, dimY int) (Day18Puzzle, error) {
	var points []int
	for i, line := range lines {
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			return Day18Puzzle{}, fmt.Errorf("want <number>,<number> but got %q in line %d", line, i+1)
		}
		x, err := strconv.Atoi(parts[0])
		if err != nil {
			return Day18Puzzle{}, fmt.Errorf("cannot parse x coordinate in line %d", i+1)
		}
		y, err := strconv.Atoi(parts[1])
		if err != nil {
			return Day18Puzzle{}, fmt.Errorf("cannot parse y coordinate in line %d", i+1)
		}
		if x >= 0 && x < dimX && y >= 0 && y < dimY {
			points = append(points, y*dimX+x)
		}
	}
	return Day18Puzzle{points: points, dimX: dimX, dimY: dimY}, nil
}

// Day18 solves Day 18 using flat arrays and index arithmetic
func Day18(p Day18Puzzle, part1 bool) (uint, string) {
	if part1 {
		return day18Part1(p), ""
	}
	x, y := day18Part2(p)
	return 0, fmt.Sprintf("%d,%d", x, y)
}

func day18Part1(p Day18Puzzle) uint {
	dimX, dimY := p.dimX, p.dimY
	size := dimX * dimY

	// Flat grid: 0 = open, 1 = corrupted
	grid := make([]byte, size)
	for _, idx := range p.points {
		grid[idx] = 1
	}

	return bfsFlat(grid, dimX, dimY)
}

func day18Part2(p Day18Puzzle) (int, int) {
	// Binary search for first blocking byte
	left, right := 0, len(p.points)-1

	for left < right {
		mid := (left + right) / 2
		if canReachFlat(p, mid+1) {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if left < len(p.points) {
		idx := p.points[left]
		return idx % p.dimX, idx / p.dimX
	}
	return 0, 0
}

func canReachFlat(p Day18Puzzle, limit int) bool {
	dimX, dimY := p.dimX, p.dimY
	size := dimX * dimY

	grid := make([]byte, size)
	for i := range min(limit, len(p.points)) {
		grid[p.points[i]] = 1
	}

	return bfsFlat(grid, dimX, dimY) > 0
}

func bfsFlat(grid []byte, dimX, dimY int) uint {
	size := dimX * dimY
	start, end := 0, size-1

	if grid[start] == 1 || grid[end] == 1 {
		return 0
	}

	// Flat visited
	visited := make([]bool, size)
	visited[start] = true

	// Direction offsets: N, E, S, W
	dirs := [4]int{-dimX, 1, dimX, -1}

	// BFS with flat indices, ring buffer style
	type state struct {
		idx   int
		steps uint
	}
	queue := make([]state, 0, size/4)
	queue = append(queue, state{idx: start, steps: 0})
	head := 0

	for head < len(queue) {
		cur := queue[head]
		head++

		if cur.idx == end {
			return cur.steps
		}

		x, y := cur.idx%dimX, cur.idx/dimX
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

			ni := cur.idx + d
			if !visited[ni] && grid[ni] == 0 {
				visited[ni] = true
				queue = append(queue, state{idx: ni, steps: cur.steps + 1})
			}
		}
	}
	return 0
}
