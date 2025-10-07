package adventofcode2024

import "fmt"

var day15TestHook func(gridBytes []byte)

// widen transforms the warehouse map according to Part 2 rules:
// # becomes ##, O becomes [], . becomes .., @ becomes @.
// Accepts input with no trailing newline, one trailing newline (grid only),
// or two trailing newlines (complete puzzle with moves section).
// The into parameter must be large enough to hold the widened result.
func widen(from []byte, into []byte) int {
	n := 0
	for i, b := range from {
		// Stop at double newline (end of grid)
		if b == '\n' && i+1 < len(from) && from[i+1] == '\n' {
			into[n] = b
			n++
			return n
		}

		switch b {
		case '#':
			into[n] = '#'
			into[n+1] = '#'
			n += 2
		case 'O':
			into[n] = '['
			into[n+1] = ']'
			n += 2
		case '.':
			into[n] = '.'
			into[n+1] = '.'
			n += 2
		case '@':
			into[n] = '@'
			into[n+1] = '.'
			n += 2
		default:
			// Keep newlines and other characters as-is
			into[n] = b
			n++
		}
	}
	return n
}

func Day15(puzzle []byte, part1 bool) (uint, error) {
	const (
		robot   = '@'
		empty   = '.'
		wall    = '#'
		box     = 'O'
		newline = '\n'
	)
	var dimX, dimY int
	var moves int
	robotPos := -1

	for i := 0; i < len(puzzle)-1; i++ {
		if puzzle[i] == robot {
			robotPos = i
		}

		if puzzle[i] == newline {
			if dimX == 0 {
				dimX = i
			}
			dimY++

			if puzzle[i+1] == newline {
				moves = i + 2
				break
			}
		}
	}

	if robotPos == -1 {
		return 0, fmt.Errorf("no robot found")
	}

	// Copy the grid to avoid modifying the original
	gridSize := moves - 1
	var grid []byte
	var gridDimX int
	if part1 {
		grid = make([]byte, gridSize)
		copy(grid, puzzle[:gridSize])
		gridDimX = dimX
	} else {
		grid = make([]byte, gridSize*2)
		actualSize := widen(puzzle[:gridSize], grid)
		grid = grid[:actualSize]
		gridDimX = dimX * 2
		// Find robot position in widened grid
		robotPos = -1
		for i, b := range grid {
			if b == robot {
				robotPos = i
				break
			}
		}
		if robotPos == -1 {
			return 0, fmt.Errorf("no robot found in widened grid")
		}
	}

	commandCount := 0
	for i := moves; i < len(puzzle); i++ {
		direction := puzzle[i]
		if direction == newline {
			continue
		}
		commandCount++
		var dx, dy int
		var targetPos int
		var targetChar byte
		var robotY, robotX int
		var targetX, targetY int

		switch direction {
		case '<':
			dx, dy = -1, 0
		case '>':
			dx, dy = 1, 0
		case '^':
			dx, dy = 0, -1
		case 'v':
			dx, dy = 0, 1
		default:
			return 0, fmt.Errorf("invalid direction: %c", direction)
		}

		robotY = robotPos / (gridDimX + 1)
		robotX = robotPos % (gridDimX + 1)

		targetX = robotX + dx
		targetY = robotY + dy

		targetPos = targetY*(gridDimX+1) + targetX

		targetChar = grid[targetPos]

		if targetChar == wall {
			// Wall blocks movement - do nothing
		} else if targetChar == empty {
			grid[robotPos] = empty
			grid[targetPos] = robot
			robotPos = targetPos
		} else if targetChar == box || (!part1 && (targetChar == '[' || targetChar == ']')) {
			if dx == 0 { // Vertical movement
				if part1 {
					// Narrow boxes vertical movement
					if pushNarrow(grid, robotX, robotY, dx, dy, gridDimX, dimY, targetX, targetY) {
						grid[robotPos] = empty
						grid[targetPos] = robot
						robotPos = targetPos
					}
				} else {
					// Wide boxes vertical movement
					if pushWide(grid, robotX, robotY, dy, gridDimX, dimY) {
						grid[robotPos] = empty
						grid[targetPos] = robot
						robotPos = targetPos
					}
				}
			} else { // Horizontal movement
				checkX := targetX
				checkY := targetY

				for {
					checkX += dx
					checkY += dy

					if checkX < 0 || checkX >= gridDimX || checkY < 0 || checkY >= dimY {
						goto blockedMove
					}

					checkPos := checkY*(gridDimX+1) + checkX
					checkChar := grid[checkPos]

					if checkChar == wall {
						goto blockedMove
					} else if checkChar == empty {
						break
					}
				}

				emptyPos := checkY*(gridDimX+1) + checkX
				for checkX != targetX || checkY != targetY {
					prevX := checkX - dx
					prevY := checkY - dy
					prevPos := prevY*(gridDimX+1) + prevX

					grid[emptyPos] = grid[prevPos]
					emptyPos = prevPos

					checkX = prevX
					checkY = prevY
				}

				grid[robotPos] = empty
				grid[targetPos] = robot
				robotPos = targetPos
			}
		}

	blockedMove:
		if day15TestHook != nil {
			day15TestHook(grid[:moves-1])
		}
	}

	var total uint
	gridLimit := len(grid)
	for i := 0; i < gridLimit; i++ {
		if (part1 && grid[i] == box) || (!part1 && grid[i] == '[') {
			y := i / (gridDimX + 1)
			x := i % (gridDimX + 1)
			if part1 {
				total += uint(100*y + x)
			} else {
				// For Part 2, use coordinates in the widened grid
				total += uint(100*y + x)
			}
		}
	}

	return total, nil
}

// pushNarrow handles vertical movement of narrow boxes in Part 1
func pushNarrow(input []byte, robotX, robotY, dx, dy, gridDimX, dimY, targetX, targetY int) bool {
	const (
		empty = '.'
		wall  = '#'
	)

	checkX := targetX
	checkY := targetY

	for {
		checkX += dx
		checkY += dy

		if checkX < 0 || checkX >= gridDimX || checkY < 0 || checkY >= dimY {
			return false
		}

		checkPos := checkY*(gridDimX+1) + checkX
		checkChar := input[checkPos]

		if checkChar == wall {
			return false
		} else if checkChar == empty {
			break
		}
	}

	emptyPos := checkY*(gridDimX+1) + checkX
	for checkX != targetX || checkY != targetY {
		prevX := checkX - dx
		prevY := checkY - dy
		prevPos := prevY*(gridDimX+1) + prevX

		input[emptyPos] = input[prevPos]
		emptyPos = prevPos

		checkX = prevX
		checkY = prevY
	}

	return true
}

// pushWide handles vertical movement of wide boxes in Part 2
func pushWide(grid []byte, robotX, robotY, dy, gridDimX, dimY int) bool {
	const (
		empty    = '.'
		wall     = '#'
		boxLeft  = '['
		boxRight = ']'
	)

	// Collect all boxes that need to move with their original characters
	type boxMove struct {
		pos      int
		original byte
	}

	var boxesToMove []boxMove
	var toCheck []int
	visited := make(map[int]bool)

	targetY := robotY + dy
	targetPos := targetY*(gridDimX+1) + robotX
	targetChar := grid[targetPos]

	// Start with the box the robot is pushing
	if targetChar == boxLeft {
		toCheck = append(toCheck, targetPos, targetPos+1) // Both parts of the box
	} else if targetChar == boxRight {
		toCheck = append(toCheck, targetPos-1, targetPos) // Both parts of the box
	}

	// Find all connected boxes that need to move
	for len(toCheck) > 0 {
		pos := toCheck[0]
		toCheck = toCheck[1:]

		if visited[pos] || pos < 0 || pos >= len(grid) {
			continue
		}
		visited[pos] = true

		char := grid[pos]
		if char != boxLeft && char != boxRight {
			continue
		}

		boxesToMove = append(boxesToMove, boxMove{pos: pos, original: char})

		// Check the position this box would move to
		nextY := (pos / (gridDimX + 1)) + dy
		nextX := pos % (gridDimX + 1)

		if nextY < 0 || nextY >= dimY {
			return false // Out of bounds
		}

		nextPos := nextY*(gridDimX+1) + nextX
		if nextPos < 0 || nextPos >= len(grid) {
			return false
		}

		nextChar := grid[nextPos]

		if nextChar == wall {
			return false // Blocked by wall
		} else if nextChar == boxLeft && !visited[nextPos] {
			toCheck = append(toCheck, nextPos, nextPos+1)
		} else if nextChar == boxRight && !visited[nextPos] {
			toCheck = append(toCheck, nextPos-1, nextPos)
		}
	}

	// Move all boxes
	// First clear all box positions
	for _, box := range boxesToMove {
		grid[box.pos] = empty
	}

	// Then place boxes in new positions
	for _, box := range boxesToMove {
		oldY := box.pos / (gridDimX + 1)
		oldX := box.pos % (gridDimX + 1)
		newY := oldY + dy
		newPos := newY*(gridDimX+1) + oldX

		if newPos >= 0 && newPos < len(grid) {
			grid[newPos] = box.original
		}
	}

	return true
}
