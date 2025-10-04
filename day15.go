package adventofcode2024

import "fmt"

func Day15(input []byte, part1 bool) (uint, error) {
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

	for i := 0; i < len(input)-1; i++ {
		if input[i] == robot {
			robotPos = i
		}

		if input[i] == newline {
			if dimX == 0 {
				dimX = i
			}
			dimY++

			if input[i+1] == newline {
				moves = i + 2
				break
			}
		}
	}


	if robotPos == -1 {
		return 0, fmt.Errorf("no robot found")
	}

	commandCount := 0
	for i := moves; i < len(input); i++ {
		direction := input[i]
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

		robotY = robotPos / (dimX + 1)
		robotX = robotPos % (dimX + 1)

		targetX = robotX + dx
		targetY = robotY + dy

		targetPos = targetY*(dimX+1) + targetX

		targetChar = input[targetPos]

		if targetChar == wall {
			// Wall blocks movement - do nothing
		} else if targetChar == empty {
			input[robotPos] = empty
			input[targetPos] = robot
			robotPos = targetPos
		} else if targetChar == box {
			checkX := targetX
			checkY := targetY

			for {
				checkX += dx
				checkY += dy

				if checkX < 0 || checkX >= dimX || checkY < 0 || checkY >= dimY {
					goto blockedMove
				}

				checkPos := checkY*(dimX+1) + checkX
				checkChar := input[checkPos]

				if checkChar == wall {
					goto blockedMove
				} else if checkChar == empty {
					break
				}
			}

			emptyPos := checkY*(dimX+1) + checkX
			for checkX != targetX || checkY != targetY {
				prevX := checkX - dx
				prevY := checkY - dy
				prevPos := prevY*(dimX+1) + prevX

				input[emptyPos] = input[prevPos]
				emptyPos = prevPos

				checkX = prevX
				checkY = prevY
			}

			input[robotPos] = empty
			input[targetPos] = robot
			robotPos = targetPos
		}

	blockedMove:
		if day15TestHook != nil {
			day15TestHook(input[:moves-1])
		}
	}


	var total uint
	for i := 0; i < moves-1; i++ {
		if input[i] == box {
			y := i / (dimX + 1)
			x := i % (dimX + 1)
			total += uint(100*y + x)
		}
	}

	return total, nil
}
