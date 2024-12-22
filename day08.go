package adventofcode2024

import (
	"fmt"
	"image"
)

func Day08(lines []string, part1 bool) (total uint) {
	const (
		empty = '.'
	)
	var (
		dimX = len(lines[0])
		dimY = len(lines)
		off  = func(p image.Point) bool {
			return p.X < 0 || p.X >= dimX || p.Y < 0 || p.Y >= dimY
		}

		uniques = make(map[image.Point]bool, dimX*dimY)
	)

	for y := range dimY {
		for x := range dimX {
			c := lines[y][x]
			if c == empty {
				continue
			}
			// search all pairing freq => O(n²)
			for y2 := range dimY {
				for x2 := range dimX {
					c2 := lines[y2][x2]
					if c2 != c {
						continue
					}
					if x2 == x && y2 == y { // ignore me myself and i
						continue
					}
					// my antinodes
					dx, dy := x2-x, y2-y
					var p image.Point
					if part1 {
						// start from opposite antenna
						p = image.Point{x2, y2}
					} else {
						// start from base antenna
						p = image.Point{x, y}
					}
					for {
						p.X = p.X + dx
						p.Y = p.Y + dy
						if off(p) {
							break
						}
						uniques[p] = true
						if part1 {
							// part 1: frequency stops after one wave
							break
						}
						// part 2: waves keep going
					}
				}
			}
		}
	}
	return uint(len(uniques))
}

func DumpDay08(lines []string, ps map[image.Point]bool) {
	fmt.Println()
	for y := range lines {
		for x := range lines[0] {
			if _, ok := ps[image.Point{x, y}]; ok {
				fmt.Printf("#")
			} else {
				fmt.Printf("%c", lines[y][x])
			}
		}
		fmt.Println()
	}
}
