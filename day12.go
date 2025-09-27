package adventofcode2024

import "image"

type Day12Puzzle struct {
	grid [][]byte
	dimY int
	dimX int
}

func NewDay12(lines []string) Day12Puzzle {
	dimY := len(lines)
	if dimY == 0 {
		return Day12Puzzle{}
	}
	dimX := len(lines[0])

	grid := make([][]byte, dimY)
	for y := range grid {
		grid[y] = []byte(lines[y])
	}

	return Day12Puzzle{
		grid: grid,
		dimY: dimY,
		dimX: dimX,
	}
}

func Day12(puzzle Day12Puzzle) uint {
	var totalPrice uint
	visited := make([][]bool, puzzle.dimY)
	for y := range visited {
		visited[y] = make([]bool, puzzle.dimX)
	}

	for y := range puzzle.dimY {
		for x := range puzzle.dimX {
			if !visited[y][x] {
				area, perimeter := exploreRegion(puzzle, visited, y, x, puzzle.grid[y][x])
				price := area * perimeter
				totalPrice += price
			}
		}
	}

	return totalPrice
}

func exploreRegion(puzzle Day12Puzzle, visited [][]bool, startY, startX int, plantType byte) (area, perimeter uint) {
	if visited[startY][startX] {
		return 0, 0
	}

	// Use a stack for iterative DFS
	stack := []image.Point{{X: startX, Y: startY}}
	directions := []image.Point{{X: 0, Y: -1}, {X: 0, Y: 1}, {X: -1, Y: 0}, {X: 1, Y: 0}}

	for len(stack) > 0 {
		// Pop from stack
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		y, x := current.Y, current.X

		// Skip if already visited or out of bounds
		if y < 0 || y >= puzzle.dimY || x < 0 || x >= puzzle.dimX {
			continue
		}
		if puzzle.grid[y][x] != plantType {
			continue
		}
		if visited[y][x] {
			continue
		}

		// Mark as visited and count area
		visited[y][x] = true
		area++

		// Calculate perimeter for this cell
		for _, dir := range directions {
			next := current.Add(dir)
			newY, newX := next.Y, next.X
			if newY < 0 || newY >= puzzle.dimY || newX < 0 || newX >= puzzle.dimX {
				perimeter++ // Edge of grid
			} else if puzzle.grid[newY][newX] != plantType {
				perimeter++ // Different plant type
			} else if !visited[newY][newX] {
				// Same plant type, not visited - add to stack
				stack = append(stack, next)
			}
		}
	}

	return area, perimeter
}