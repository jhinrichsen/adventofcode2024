package adventofcode2024

import (
	"image"
)

func Day06(grid [][]byte) uint {
	const (
		block = '#'
		guard = '^'
	)
	var (
		dimX = len(grid[0])
		dimY = len(grid)
		dir  = image.Point{0, -1}
		pos  = func() image.Point {
			for y := range dimY {
				for x := range dimY {
					if grid[y][x] == guard {
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
			return grid[p.Y][p.X] == block
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
