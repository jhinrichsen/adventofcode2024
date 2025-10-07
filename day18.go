package adventofcode2024

import (
	"fmt"
	"image"
	"strconv"
	"strings"
)

type Day18Puzzle struct {
	points     []image.Point
	dimX, dimY int
}

func NewDay18(lines []string, dimX, dimY int) (Day18Puzzle, error) {
	a := Day18Puzzle{dimX: dimX, dimY: dimY}
	for i, line := range lines {
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			return a, fmt.Errorf("want <number>,<number> but got %q in line %d", line, i+1)
		}
		x, err := strconv.Atoi(parts[0])
		if err != nil {
			return a, fmt.Errorf("cannot parse x coordinate in line %d", i+1)
		}
		y, err := strconv.Atoi(parts[1])
		if err != nil {
			return a, fmt.Errorf("cannot parse y coordinate in line %d", i+1)
		}
		a.points = append(a.points, image.Point{x, y})
	}
	return a, nil
}

func Day18(a Day18Puzzle, part1 bool) (uint, string) {
	if part1 {
		steps := findShortestPath(a)
		return steps, ""
	}
	coord := findBlockingByteString(a)
	return 0, coord
}

func findShortestPath(a Day18Puzzle) uint {
	const corrupted = '#'

	grid := make([][]byte, a.dimY)
	for y := range a.dimY {
		grid[y] = make([]byte, a.dimX)
		for x := range a.dimX {
			grid[y][x] = '.'
		}
	}

	// Apply all corruption from input
	for _, p := range a.points {
		if p.X >= 0 && p.X < a.dimX && p.Y >= 0 && p.Y < a.dimY {
			grid[p.Y][p.X] = corrupted
		}
	}

	return bfsShortestPath(grid, a.dimX, a.dimY)
}

func findBlockingByteString(a Day18Puzzle) string {
	// Binary search to find the first byte that blocks all paths
	left, right := 0, len(a.points)-1

	for left < right {
		mid := (left + right) / 2
		if canReachExit(a, mid+1) {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if left < len(a.points) {
		p := a.points[left]
		return fmt.Sprintf("%d,%d", p.X, p.Y)
	}
	return "0,0"
}

func canReachExit(a Day18Puzzle, corruptionLimit int) bool {
	const corrupted = '#'

	grid := make([][]byte, a.dimY)
	for y := range a.dimY {
		grid[y] = make([]byte, a.dimX)
		for x := range a.dimX {
			grid[y][x] = '.'
		}
	}

	// Apply corruption up to limit
	limit := min(corruptionLimit, len(a.points))
	for i := range limit {
		p := a.points[i]
		if p.X >= 0 && p.X < a.dimX && p.Y >= 0 && p.Y < a.dimY {
			grid[p.Y][p.X] = corrupted
		}
	}

	return bfsShortestPath(grid, a.dimX, a.dimY) > 0
}


func bfsShortestPath(grid [][]byte, dimX, dimY int) uint {
	const corrupted = '#'

	start := image.Point{X: 0, Y: 0}
	end := image.Point{X: dimX - 1, Y: dimY - 1}

	if grid[start.Y][start.X] == corrupted || grid[end.Y][end.X] == corrupted {
		return 0 // blocked start or end
	}

	if start == end {
		return 0
	}

	visited := make([][]bool, dimY)
	for y := range dimY {
		visited[y] = make([]bool, dimX)
	}

	type state struct {
		pos   image.Point
		steps uint
	}

	queue := []state{{pos: start, steps: 0}}
	visited[start.Y][start.X] = true

	directions := []image.Point{
		{X: 0, Y: -1}, // up
		{X: 0, Y: 1},  // down
		{X: -1, Y: 0}, // left
		{X: 1, Y: 0},  // right
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.pos == end {
			return current.steps
		}

		for _, dir := range directions {
			next := image.Point{
				X: current.pos.X + dir.X,
				Y: current.pos.Y + dir.Y,
			}

			if next.X >= 0 && next.X < dimX &&
				next.Y >= 0 && next.Y < dimY &&
				!visited[next.Y][next.X] &&
				grid[next.Y][next.X] != corrupted {

				visited[next.Y][next.X] = true
				queue = append(queue, state{
					pos:   next,
					steps: current.steps + 1,
				})
			}
		}
	}

	return 0 // no path found
}
