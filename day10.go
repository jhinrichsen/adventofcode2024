package adventofcode2024

import (
	"container/list"
	"image"
)

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

// Day10 solves Day 10 using the puzzle struct
func Day10(puzzle Day10Puzzle, part1 bool) uint {
	// Internal types for this algorithm
	type trail [10]image.Point
	type distinctTrail struct {
		start, end image.Point
	}

	var (
		dimX  = len(puzzle.grid[0])
		dimY  = len(puzzle.grid)
		board = image.Rect(0, 0, dimX, dimY)
	)

	// list all trail starts

	var all = list.New()
	for y := 0; y < dimY; y++ {
		for x := 0; x < dimX; x++ {
			if puzzle.grid[y][x] == '0' {
				var t trail
				t[0] = image.Point{x, y}
				all.PushBack(t)
			}
		}
	}

	for i := 1; i <= 9; i++ {
		for j := all.Len(); j > 0; j-- {
			e := all.Front()
			all.Remove(e)
			if trail, ok := e.Value.(trail); ok {
				for _, delta := range []image.Point{
					{0, -1}, // N
					{1, 0},  // E
					{0, 1},  // S
					{-1, 0}, // W
				} {
					p := trail[i-1]
					p = p.Add(delta)
					if !p.In(board) {
						continue
					}
					want := byte(i + '0')
					got := puzzle.grid[p.Y][p.X]
					if want != got {
						continue
					}
					trail[i] = p
					all.PushBack(trail)
				}
			}
		}
	}

	// part 1: count distinct (start, end) pairs
	if part1 {
		m := make(map[distinctTrail]bool, all.Len())
		for e := all.Front(); e != nil; e = e.Next() {
			if t, ok := e.Value.(trail); ok {
				d := distinctTrail{t[0], t[9]}
				m[d] = true
			}
		}
		return uint(len(m))
	}

	// part 2: count all trails
	return uint(all.Len())
}
