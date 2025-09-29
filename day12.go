package adventofcode2024

import "image"

type Day12Puzzle struct {
	grid [][]byte
	dimY int
	dimX int
}

func NewDay12(lines []string) Day12Puzzle {
	// Filter out empty lines
	var validLines []string
	for _, line := range lines {
		if len(line) > 0 {
			validLines = append(validLines, line)
		}
	}
	
	dimY := len(validLines)
	if dimY == 0 {
		return Day12Puzzle{}
	}
	dimX := len(validLines[0])

	grid := make([][]byte, dimY)
	for y := range grid {
		grid[y] = []byte(validLines[y])
	}

	return Day12Puzzle{
		grid: grid,
		dimY: dimY,
		dimX: dimX,
	}
}

func Day12(puzzle Day12Puzzle, part1 bool) uint {
	var totalPrice uint
	visited := make([][]bool, puzzle.dimY)
	for y := range visited {
		visited[y] = make([]bool, puzzle.dimX)
	}

	for y := range puzzle.dimY {
		for x := range puzzle.dimX {
			if !visited[y][x] {
				if part1 {
					area, perimeter := exploreRegion(puzzle, visited, y, x, puzzle.grid[y][x])
					price := area * perimeter
					totalPrice += price
				} else {
					area, sides := exploreRegionWithSides(puzzle, visited, y, x, puzzle.grid[y][x])
					price := area * sides
					totalPrice += price
				}
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

		// Skip if already visited or out of bounds or wrong plant type
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

		// Calculate perimeter for this cell and add neighbors to stack
		for _, dir := range directions {
			newY, newX := y+dir.Y, x+dir.X
			if newY < 0 || newY >= puzzle.dimY || newX < 0 || newX >= puzzle.dimX {
				perimeter++ // Edge of grid
			} else if puzzle.grid[newY][newX] != plantType {
				perimeter++ // Different plant type
			} else if !visited[newY][newX] {
				// Same plant type, not visited - add to stack
				stack = append(stack, image.Point{X: newX, Y: newY})
			}
		}
	}

	return area, perimeter
}

func exploreRegionWithSides(puzzle Day12Puzzle, visited [][]bool, startY, startX int, plantType byte) (area, sides uint) {
	if visited[startY][startX] {
		return 0, 0
	}

	// First, collect all cells in the region
	var regionCells []image.Point
	stack := []image.Point{{X: startX, Y: startY}}
	directions := []image.Point{{X: 0, Y: -1}, {X: 0, Y: 1}, {X: -1, Y: 0}, {X: 1, Y: 0}}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		y, x := current.Y, current.X

		if y < 0 || y >= puzzle.dimY || x < 0 || x >= puzzle.dimX {
			continue
		}
		if puzzle.grid[y][x] != plantType {
			continue
		}
		if visited[y][x] {
			continue
		}

		visited[y][x] = true
		area++
		regionCells = append(regionCells, current)

		for _, dir := range directions {
			newY, newX := y+dir.Y, x+dir.X
			if newY >= 0 && newY < puzzle.dimY && newX >= 0 && newX < puzzle.dimX &&
				puzzle.grid[newY][newX] == plantType && !visited[newY][newX] {
				stack = append(stack, image.Point{X: newX, Y: newY})
			}
		}
	}

	// Count corners (which equals number of sides)
	sides = countCorners(puzzle, regionCells, plantType)
	return area, sides
}

func countCorners(puzzle Day12Puzzle, regionCells []image.Point, plantType byte) uint {
	// Create a set of region cells for fast lookup
	regionSet := make(map[image.Point]bool)
	for _, cell := range regionCells {
		regionSet[cell] = true
	}

	var corners uint
	
	// For each cell in the region, count corners
	for _, cell := range regionCells {
		y, x := cell.Y, cell.X
		
		// Check all 4 possible corner positions around this cell
		// A corner exists when we have a specific pattern of region/non-region cells
		
		// Top-left corner
		up := regionSet[image.Point{X: x, Y: y - 1}]
		left := regionSet[image.Point{X: x - 1, Y: y}]
		upLeft := regionSet[image.Point{X: x - 1, Y: y - 1}]
		
		// External corner: neither up nor left are in region
		if !up && !left {
			corners++
		}
		// Internal corner: both up and left are in region, but diagonal is not
		if up && left && !upLeft {
			corners++
		}
		
		// Top-right corner
		right := regionSet[image.Point{X: x + 1, Y: y}]
		upRight := regionSet[image.Point{X: x + 1, Y: y - 1}]
		
		if !up && !right {
			corners++
		}
		if up && right && !upRight {
			corners++
		}
		
		// Bottom-left corner
		down := regionSet[image.Point{X: x, Y: y + 1}]
		downLeft := regionSet[image.Point{X: x - 1, Y: y + 1}]
		
		if !down && !left {
			corners++
		}
		if down && left && !downLeft {
			corners++
		}
		
		// Bottom-right corner
		downRight := regionSet[image.Point{X: x + 1, Y: y + 1}]
		
		if !down && !right {
			corners++
		}
		if down && right && !downRight {
			corners++
		}
	}
	
	return corners
}