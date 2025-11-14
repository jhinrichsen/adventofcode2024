package adventofcode2024

import (
	"image"
	"strconv"
	"strings"
)

type Day21Puzzle struct {
	codes []string
}

func NewDay21(lines []string) Day21Puzzle {
	codes := make([]string, 0, len(lines))
	for _, line := range lines {
		if line := strings.TrimSpace(line); line != "" {
			codes = append(codes, line)
		}
	}
	return Day21Puzzle{codes: codes}
}

func Day21(puzzle Day21Puzzle, part1 bool) uint {
	if part1 {
		return puzzle.calculateComplexity()
	}
	return 0
}

func (p Day21Puzzle) calculateComplexity() uint {
	var totalComplexity uint

	for _, code := range p.codes {
		sequence := p.findShortestSequence(code)
		numericValue := p.extractNumericValue(code)
		complexity := uint(len(sequence)) * numericValue
		totalComplexity += complexity
	}

	return totalComplexity
}

func (p Day21Puzzle) findShortestSequence(code string) string {
	// First robot (directional) controls second robot (directional) controls third robot (numeric)
	// We work backwards: numeric -> directional -> directional

	// Step 1: Get sequence needed for numeric keypad
	numericSeq := p.getNumericKeypadSequence(code)

	// Step 2: Get sequence needed for first directional keypad to input numericSeq
	dirSeq1 := p.getDirectionalKeypadSequence(numericSeq)

	// Step 3: Get sequence needed for second directional keypad to input dirSeq1
	dirSeq2 := p.getDirectionalKeypadSequence(dirSeq1)

	return dirSeq2
}

func (p Day21Puzzle) getNumericKeypadSequence(code string) string {
	numericKeypad := map[byte]image.Point{
		'7': {0, 0}, '8': {1, 0}, '9': {2, 0},
		'4': {0, 1}, '5': {1, 1}, '6': {2, 1},
		'1': {0, 2}, '2': {1, 2}, '3': {2, 2},
		'0': {1, 3}, 'A': {2, 3},
	}

	var result strings.Builder
	currentPos := numericKeypad['A'] // Start at A

	for i := range code {
		targetPos := numericKeypad[code[i]]
		path := p.findPathNumeric(currentPos, targetPos)
		result.WriteString(path)
		result.WriteByte('A') // Press the button
		currentPos = targetPos
	}

	return result.String()
}

func (p Day21Puzzle) getDirectionalKeypadSequence(sequence string) string {
	directionalKeypad := map[byte]image.Point{
		'^': {1, 0}, 'A': {2, 0},
		'<': {0, 1}, 'v': {1, 1}, '>': {2, 1},
	}

	var result strings.Builder
	currentPos := directionalKeypad['A'] // Start at A

	for i := range sequence {
		targetPos := directionalKeypad[sequence[i]]
		path := p.findPathDirectional(currentPos, targetPos)
		result.WriteString(path)
		result.WriteByte('A') // Press the button
		currentPos = targetPos
	}

	return result.String()
}

func (p Day21Puzzle) findPathNumeric(from, to image.Point) string {
	// Avoid the empty space at (0, 3)
	dx := to.X - from.X
	dy := to.Y - from.Y

	var moves strings.Builder

	// Strategy: Try to avoid going through (0, 3)
	// If we're at (0, 3) or going to (0, 3), be careful about order

	if from.X == 0 && to.Y == 3 {
		// Moving from left column to bottom row - go right first
		for range dx {
			if dx > 0 {
				moves.WriteByte('>')
			}
		}
		for range dy {
			if dy > 0 {
				moves.WriteByte('v')
			} else if dy < 0 {
				moves.WriteByte('^')
			}
		}
	} else if from.Y == 3 && to.X == 0 {
		// Moving from bottom row to left column - go up first
		for range -dy {
			if dy < 0 {
				moves.WriteByte('^')
			}
		}
		for range -dx {
			if dx < 0 {
				moves.WriteByte('<')
			}
		}
	} else {
		// Normal case - prefer left/right first, then up/down
		for range -dx {
			if dx < 0 {
				moves.WriteByte('<')
			}
		}
		for range dx {
			if dx > 0 {
				moves.WriteByte('>')
			}
		}
		for range -dy {
			if dy < 0 {
				moves.WriteByte('^')
			}
		}
		for range dy {
			if dy > 0 {
				moves.WriteByte('v')
			}
		}
	}

	return moves.String()
}

func (p Day21Puzzle) findPathDirectional(from, to image.Point) string {
	// Avoid the empty space at (0, 0)
	dx := to.X - from.X
	dy := to.Y - from.Y

	var moves strings.Builder

	// Strategy: Try to avoid going through (0, 0)
	// If we're at (0, 0) or going to (0, 0), be careful about order

	if from.X == 0 && to.Y == 0 {
		// Moving from left to top row - go right first
		for range dx {
			if dx > 0 {
				moves.WriteByte('>')
			}
		}
		for range -dy {
			if dy < 0 {
				moves.WriteByte('^')
			}
		}
	} else if from.Y == 0 && to.X == 0 {
		// Moving from top row to left - go down first
		for range dy {
			if dy > 0 {
				moves.WriteByte('v')
			}
		}
		for range -dx {
			if dx < 0 {
				moves.WriteByte('<')
			}
		}
	} else {
		// Normal case - prefer left/right first, then up/down
		for range -dx {
			if dx < 0 {
				moves.WriteByte('<')
			}
		}
		for range dx {
			if dx > 0 {
				moves.WriteByte('>')
			}
		}
		for range -dy {
			if dy < 0 {
				moves.WriteByte('^')
			}
		}
		for range dy {
			if dy > 0 {
				moves.WriteByte('v')
			}
		}
	}

	return moves.String()
}

func (p Day21Puzzle) extractNumericValue(code string) uint {
	// Extract numeric part (remove 'A' suffix)
	numStr := strings.TrimSuffix(code, "A")
	val, _ := strconv.Atoi(numStr)
	return uint(val)
}