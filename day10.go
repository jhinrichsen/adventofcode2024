package adventofcode2024

import (
	"container/list"
	"image"
)

type trail [10]image.Point
type distinctTrail struct {
	start, end image.Point
}

func Day10(grid [][]byte) uint {
	var (
		dimX  = len(grid[0])
		dimY  = len(grid)
		board = image.Rect(0, 0, dimX, dimY)
	)

	// list all trail starts

	var all = list.New()
	for y := 0; y < dimY; y++ {
		for x := 0; x < dimX; x++ {
			if grid[y][x] == '0' {
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
					image.Point{0, -1}, // N
					image.Point{1, 0},  // E
					image.Point{0, 1},  // S
					image.Point{-1, 0}, // W
				} {
					p := trail[i-1]
					p = p.Add(delta)
					if !p.In(board) {
						continue
					}
					want := byte(i + '0')
					got := grid[p.Y][p.X]
					if want != got {
						continue
					}
					trail[i] = p
					all.PushBack(trail)
				}
			}
		}
	}

	// multiple trails from 0 -> 9 count as 1 for the same 0 and 9

	m := make(map[distinctTrail]bool, all.Len())
	for e := all.Front(); e != nil; e = e.Next() {
		if t, ok := e.Value.(trail); ok {
			d := distinctTrail{t[0], t[9]}
			m[d] = true
		}
	}
	return uint(len(m))
}
