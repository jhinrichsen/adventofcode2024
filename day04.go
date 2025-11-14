package adventofcode2024

import (
	"image"
)

// Day04Puzzle represents the parsed Day 4 puzzle data
type Day04Puzzle struct {
	grid [][]byte
	dimX int
	dimY int
}

// NewDay04 creates a Day04Puzzle from lines
func NewDay04(lines []string) (Day04Puzzle, error) {
	dimY := len(lines)
	grid := make([][]byte, dimY)
	for y := range grid {
		grid[y] = []byte(lines[y])
	}
	return Day04Puzzle{
		grid: grid,
		dimY: dimY,
		dimX: len(lines[0]),
	}, nil
}

// Day04 solves Day 4 using the puzzle struct (both parts)
func Day04(puzzle Day04Puzzle, part1 bool) uint {
	var (
		r = image.Rect(0, 0, len(puzzle.grid[0]), len(puzzle.grid))

		N  = image.Point{+0, -1}
		NE = image.Point{+1, -1}
		E  = image.Point{+1, 0}
		SE = image.Point{+1, +1}
		S  = image.Point{+0, +1}
		SW = image.Point{-1, +1}
		W  = image.Point{-1, 0}
		NW = image.Point{-1, -1}
	)

	var n uint
	if part1 {
		const (
			magic = "XMAS"
			l     = len(magic)
		)
		for y := range r.Max.Y {
			for x := range r.Max.X {
				// fast predicates first
				if puzzle.grid[y][x] != magic[0] {
					continue
				}
				p0 := image.Point{x, y}
				for _, dp := range []image.Point{N, NE, E, SE, S, SW, W, NW} {
					// check if end of word is still inside the grid
					p3 := p0.Add(dp.Mul(l - 1))
					if !p3.In(r) {
						continue
					}

					found := true
					// now go for the magic word itself, X already checked
					for i := 1; i < l; i++ {
						pi := p0.Add(dp.Mul(i))
						if puzzle.grid[pi.Y][pi.X] != magic[i] {
							found = false
							break
						}
					}
					if found {
						n++
					}
				}
			}
		}
		return n
	}

	// part 2
	const (
		magic = "MAS"
		l     = len(magic)
	)
	// any 'A' must be found within 1 off border
	for y := 1; y < r.Max.Y-1; y++ {
		for x := 1; x < r.Max.X-1; x++ {
			if puzzle.grid[y][x] != magic[1] {
				continue
			}
			p1 := image.Point{x, y}
			has := func(p image.Point, dir image.Point, b byte) bool {
				p2 := p.Add(dir)
				return puzzle.grid[p2.Y][p2.X] == b
			}
			hasSE := func() bool {
				return has(p1, NW, magic[0]) && has(p1, SE, magic[2])
			}
			hasNW := func() bool {
				return has(p1, SE, magic[0]) && has(p1, NW, magic[2])
			}
			hasNE := func() bool {
				return has(p1, SW, magic[0]) && has(p1, NE, magic[2])
			}
			hasSW := func() bool {
				return has(p1, NE, magic[0]) && has(p1, SW, magic[2])
			}
			if (hasSE() || hasNW()) && (hasNE() || hasSW()) {
				n++
			}
		}
	}
	return n
}
