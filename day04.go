package adventofcode2024

import "image"

func Day04(grid [][]byte) (n uint) {
	const (
		magic = "XMAS"
		l     = len(magic)
	)
	var (
		dimX = len(grid[0])
		dimY = len(grid)
		r    = image.Rect(0, 0, dimX, dimY)
	)

	for y := range dimY {
		for x := range dimX {
			// fast predicates first
			if grid[y][x] != magic[0] {
				continue
			}
			p0 := image.Point{x, y}
			for _, dp := range []image.Point{
				{+0, -1}, // N
				{+1, -1}, // NE
				{+1, 0},  // E
				{+1, +1}, // SE
				{+0, +1}, // S
				{-1, +1}, // SW
				{-1, 0},  // W
				{-1, -1}, // NW
			} {
				// check if end of word is still inside the grid
				p3 := p0.Add(dp.Mul(l - 1))
				if !p3.In(r) {
					continue
				}

				found := true
				// now go for the magic word itself, X already checked
				for i := 1; i < l; i++ {
					pi := p0.Add(dp.Mul(i))
					if grid[pi.Y][pi.X] != magic[i] {
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
	return
}
