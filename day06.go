package adventofcode2024

import (
	"image"
)

// Day06Puzzle represents the parsed Day 6 puzzle data
type Day06Puzzle struct {
	grid [][]byte
	dimX int
	dimY int
}

// NewDay06 creates a Day06Puzzle from lines
func NewDay06(lines []string) Day06Puzzle {
	dimY := len(lines)
	grid := make([][]byte, dimY)
	for y := range grid {
		grid[y] = []byte(lines[y])
	}
	return Day06Puzzle{
		grid: grid,
		dimY: dimY,
		dimX: len(lines[0]),
	}
}

// Day06 solves Day 6 using the puzzle struct
func Day06(puzzle Day06Puzzle, part1 bool) uint {
	const (
		block = '#'
		guard = '^'
	)
	var (
		dimX  = puzzle.dimX
		dimY  = puzzle.dimY
		start = func() image.Point {
			for y := range dimY {
				for x := range dimX {
					if puzzle.grid[y][x] == guard {
						return image.Point{x, y}
					}
				}
			}
			return image.Point{-1, -1}
		}()
		off = func(p image.Point) bool {
			return p.X < 0 || p.X >= dimX || p.Y < 0 || p.Y >= dimY
		}
		blocked = func(p image.Point, extraBlock *image.Point) bool {
			if extraBlock != nil && p.X == extraBlock.X && p.Y == extraBlock.Y {
				return true
			}
			return puzzle.grid[p.Y][p.X] == block
		}
	)

	if part1 {
		// Part 1: Count distinct positions visited
		dir := image.Point{0, -1}
		pos := start
		visited := make(map[image.Point]bool, dimX*dimY)

		for !off(pos) {
			p2 := pos.Add(dir)
			if off(p2) || !blocked(p2, nil) {
				pos = p2
				visited[pos] = true
			} else {
				dir = image.Point{-dir.Y, dir.X} // Y goes down
			}
		}
		return uint(len(visited)) - 1 // do not count the last step off the grid
	}

	// Part 2: Count positions where adding obstruction causes loop
	originalPath := make(map[image.Point]bool, dimX*dimY)
	{
		dir := image.Point{0, -1}
		pos := start
		for !off(pos) {
			originalPath[pos] = true
			p2 := pos.Add(dir)
			if off(p2) || !blocked(p2, nil) {
				pos = p2
			} else {
				dir = image.Point{-dir.Y, dir.X}
			}
		}
	}

	dirToInt := func(d image.Point) int {
		switch d {
		case image.Point{0, -1}:
			return 0 // up
		case image.Point{1, 0}:
			return 1 // right
		case image.Point{0, 1}:
			return 2 // down
		case image.Point{-1, 0}:
			return 3 // left
		}
		return -1
	}

	visited := make([]bool, dimX*dimY*4)

	simulate := func(obst image.Point) bool {
		for i := range visited {
			visited[i] = false
		}

		dir := image.Point{0, -1}
		pos := start

		for !off(pos) {
			stateIdx := pos.Y*dimX*4 + pos.X*4 + dirToInt(dir)
			if visited[stateIdx] {
				return true // Loop detected
			}
			visited[stateIdx] = true

			p2 := pos.Add(dir)
			if off(p2) || !blocked(p2, &obst) {
				pos = p2
			} else {
				dir = image.Point{-dir.Y, dir.X} // Y goes down
			}
		}
		return false // Guard left the grid
	}

	count := uint(0)
	for pos := range originalPath {
		if pos == start || puzzle.grid[pos.Y][pos.X] == block {
			continue
		}
		if simulate(pos) {
			count++
		}
	}
	return count
}
