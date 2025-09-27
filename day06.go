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
func Day06(puzzle Day06Puzzle) uint {
	const (
		block = '#'
		guard = '^'
	)
	var (
		dimX = len(puzzle.grid[0])
		dimY = len(puzzle.grid)
		dir  = image.Point{0, -1}
		pos  = func() image.Point {
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
		blocked = func(p image.Point) bool {
			return puzzle.grid[p.Y][p.X] == block
		}
		visited = make(map[image.Point]bool, dimX*dimY)
	)

	for !off(pos) {
		p2 := pos.Add(dir)
		if off(p2) || !blocked(p2) {
			pos = p2
			visited[pos] = true
		} else {
			dir = image.Point{-dir.Y, dir.X} // Y goes down
		}
	}
	return uint(len(visited)) - 1 // do not count the last step off the grid
}

