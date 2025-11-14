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
		return calculateComplexity(puzzle)
	}
	return 0
}

func calculateComplexity(puzzle Day21Puzzle) uint {
	var totalComplexity uint

	for _, code := range puzzle.codes {
		sequence := findShortestSequence(code)
		numericValue := extractNumericValue(code)
		complexity := uint(len(sequence)) * numericValue
		totalComplexity += complexity
	}

	return totalComplexity
}

func findShortestSequence(code string) string {
	// First robot (directional) controls second robot (directional) controls third robot (numeric)
	// We work backwards: numeric -> directional -> directional

	// Step 1: Get sequence needed for numeric keypad
	numericSeq := getNumericKeypadSequence(code)

	// Step 2: Get sequence needed for first directional keypad to input numericSeq
	dirSeq1 := getDirectionalKeypadSequence(numericSeq)

	// Step 3: Get sequence needed for second directional keypad to input dirSeq1
	dirSeq2 := getDirectionalKeypadSequence(dirSeq1)

	return dirSeq2
}

func getNumericKeypadSequence(code string) string {
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
		path := findPathNumeric(currentPos, targetPos)
		result.WriteString(path)
		result.WriteByte('A') // Press the button
		currentPos = targetPos
	}

	return result.String()
}

func getDirectionalKeypadSequence(sequence string) string {
	directionalKeypad := map[byte]image.Point{
		'^': {1, 0}, 'A': {2, 0},
		'<': {0, 1}, 'v': {1, 1}, '>': {2, 1},
	}

	var result strings.Builder
	currentPos := directionalKeypad['A'] // Start at A

	for i := range sequence {
		targetPos := directionalKeypad[sequence[i]]
		path := findPathDirectional(currentPos, targetPos)
		result.WriteString(path)
		result.WriteByte('A') // Press the button
		currentPos = targetPos
	}

	return result.String()
}

func findPathNumeric(from, to image.Point) string {
	// Avoid the empty space at (0, 3)
	// Optimal move order: < v ^ > (minimize expensive moves on directional keypad)
	dx := to.X - from.X
	dy := to.Y - from.Y

	var moves strings.Builder

	// Check if path would go through gap at (0, 3)
	goingThroughGap := (from.X == 0 && to.Y == 3) || (from.Y == 3 && to.X == 0)

	if goingThroughGap {
		// Avoid gap: if starting from left column going to bottom row, go horizontal first
		if from.X == 0 && to.Y == 3 {
			// Right then down/up
			for range dx {
				moves.WriteByte('>')
			}
			for range dy {
				moves.WriteByte('v')
			}
		} else {
			// Up then left
			for range -dy {
				moves.WriteByte('^')
			}
			for range -dx {
				moves.WriteByte('<')
			}
		}
	} else {
		// Optimal order: < v ^ >
		for range -dx {
			moves.WriteByte('<')
		}
		for range dy {
			moves.WriteByte('v')
		}
		for range -dy {
			moves.WriteByte('^')
		}
		for range dx {
			moves.WriteByte('>')
		}
	}

	return moves.String()
}

func findPathDirectional(from, to image.Point) string {
	// Avoid the empty space at (0, 0)
	// Optimal move order: < v ^ > (minimize distance from A button)
	dx := to.X - from.X
	dy := to.Y - from.Y

	var moves strings.Builder

	// Check if path would go through gap at (0, 0)
	goingThroughGap := (from.X == 0 && to.Y == 0) || (from.Y == 0 && to.X == 0)

	if goingThroughGap {
		// Avoid gap: prioritize moves that don't go through (0,0)
		if from.X == 0 && to.Y == 0 {
			// Right then up
			for range dx {
				moves.WriteByte('>')
			}
			for range -dy {
				moves.WriteByte('^')
			}
		} else {
			// Down then left
			for range dy {
				moves.WriteByte('v')
			}
			for range -dx {
				moves.WriteByte('<')
			}
		}
	} else {
		// Optimal order: < v ^ >
		for range -dx {
			moves.WriteByte('<')
		}
		for range dy {
			moves.WriteByte('v')
		}
		for range -dy {
			moves.WriteByte('^')
		}
		for range dx {
			moves.WriteByte('>')
		}
	}

	return moves.String()
}

func extractNumericValue(code string) uint {
	// Extract numeric part (remove 'A' suffix)
	numStr := strings.TrimSuffix(code, "A")
	val, _ := strconv.Atoi(numStr)
	return uint(val)
}