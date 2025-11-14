package adventofcode2024

import (
	"image"
)

type Day20Puzzle struct {
	grid  [][]byte
	dimY  int
	dimX  int
	start image.Point
	end   image.Point
}

func NewDay20(lines []string) (Day20Puzzle, error) {
	dimY := len(lines)
	grid := make([][]byte, dimY)
	var start, end image.Point

	for y := range grid {
		grid[y] = []byte(lines[y])
		for x, cell := range grid[y] {
			if cell == 'S' {
				start = image.Point{X: x, Y: y}
				grid[y][x] = '.'
			} else if cell == 'E' {
				end = image.Point{X: x, Y: y}
				grid[y][x] = '.'
			}
		}
	}

	return Day20Puzzle{
		grid:  grid,
		dimY:  dimY,
		dimX:  len(lines[0]),
		start: start,
		end:   end,
	}, nil
}

func Day20(puzzle Day20Puzzle, part1 bool) uint {
	if part1 {
		return puzzle.countCheats(100, 2)
	}
	return puzzle.countCheats(100, 20)
}

func (p Day20Puzzle) countCheats(minSaving uint, maxCheatDist int) uint {
	normalTime := p.findShortestPath(p.start, p.end)
	if normalTime == 0 {
		return 0
	}

	var count uint
	distFromStart := p.calculateDistances(p.start)
	distToEnd := p.calculateDistances(p.end)

	for y := range p.dimY {
		for x := range p.dimX {
			if p.grid[y][x] == '#' {
				continue
			}

			pos := image.Point{X: x, Y: y}

			// Try all possible cheats up to maxCheatDist from this position
			for dy := -maxCheatDist; dy <= maxCheatDist; dy++ {
				for dx := -maxCheatDist; dx <= maxCheatDist; dx++ {
					cheatDist := abs(dx) + abs(dy)
					if cheatDist == 0 || cheatDist > maxCheatDist {
						continue
					}

					newY := y + dy
					newX := x + dx
					if newY < 0 || newY >= p.dimY || newX < 0 || newX >= p.dimX {
						continue
					}
					if p.grid[newY][newX] == '#' {
						continue
					}

					endPos := image.Point{X: newX, Y: newY}

					startDist := distFromStart[pos.Y][pos.X]
					endDist := distToEnd[endPos.Y][endPos.X]

					if startDist == ^uint(0) || endDist == ^uint(0) {
						continue
					}

					cheatTime := startDist + uint(cheatDist) + endDist
					if cheatTime < normalTime && (normalTime-cheatTime) >= minSaving {
						count++
					}
				}
			}
		}
	}

	return count
}

func (p Day20Puzzle) countCheatsBySavings(maxCheatDist int) map[uint]uint {
	normalTime := p.findShortestPath(p.start, p.end)
	if normalTime == 0 {
		return nil
	}

	savings := make(map[uint]uint)
	distFromStart := p.calculateDistances(p.start)
	distToEnd := p.calculateDistances(p.end)

	for y := range p.dimY {
		for x := range p.dimX {
			if p.grid[y][x] == '#' {
				continue
			}

			pos := image.Point{X: x, Y: y}

			for dy := -maxCheatDist; dy <= maxCheatDist; dy++ {
				for dx := -maxCheatDist; dx <= maxCheatDist; dx++ {
					cheatDist := abs(dx) + abs(dy)
					if cheatDist == 0 || cheatDist > maxCheatDist {
						continue
					}

					newY := y + dy
					newX := x + dx
					if newY < 0 || newY >= p.dimY || newX < 0 || newX >= p.dimX {
						continue
					}
					if p.grid[newY][newX] == '#' {
						continue
					}

					endPos := image.Point{X: newX, Y: newY}

					startDist := distFromStart[pos.Y][pos.X]
					endDist := distToEnd[endPos.Y][endPos.X]

					if startDist == ^uint(0) || endDist == ^uint(0) {
						continue
					}

					cheatTime := startDist + uint(cheatDist) + endDist
					if cheatTime < normalTime {
						timeSaved := normalTime - cheatTime
						savings[timeSaved]++
					}
				}
			}
		}
	}

	return savings
}

func (p Day20Puzzle) findShortestPath(start, end image.Point) uint {
	distances := p.calculateDistances(start)
	if distances[end.Y][end.X] == ^uint(0) {
		return 0
	}
	return distances[end.Y][end.X]
}

func (p Day20Puzzle) calculateDistances(start image.Point) [][]uint {
	distances := make([][]uint, p.dimY)
	for y := range distances {
		distances[y] = make([]uint, p.dimX)
		for x := range distances[y] {
			distances[y][x] = ^uint(0)
		}
	}

	queue := []image.Point{start}
	distances[start.Y][start.X] = 0
	directions := []image.Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, dir := range directions {
			next := image.Point{X: current.X + dir.X, Y: current.Y + dir.Y}

			if next.X < 0 || next.X >= p.dimX || next.Y < 0 || next.Y >= p.dimY {
				continue
			}
			if p.grid[next.Y][next.X] == '#' {
				continue
			}

			newDist := distances[current.Y][current.X] + 1
			if newDist < distances[next.Y][next.X] {
				distances[next.Y][next.X] = newDist
				queue = append(queue, next)
			}
		}
	}

	return distances
}
