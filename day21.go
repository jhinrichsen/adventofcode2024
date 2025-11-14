package adventofcode2024

import (
	"image"
	"strconv"
	"strings"
)

type Day21Puzzle struct {
	codes []string
}

func NewDay21(lines []string) (Day21Puzzle, error) {
	codes := make([]string, 0, len(lines))
	for _, line := range lines {
		if line := strings.TrimSpace(line); line != "" {
			codes = append(codes, line)
		}
	}
	return Day21Puzzle{codes: codes}, nil
}

func Day21(puzzle Day21Puzzle, part1 bool) uint {
	robots := 2
	if !part1 {
		robots = 25
	}

	var totalComplexity uint
	for _, code := range puzzle.codes {
		length := solveCode(code, robots)
		numericValue := extractNumericValue(code)
		complexity := length * numericValue
		totalComplexity += complexity
	}

	return totalComplexity
}

// Cache for directional keypad transitions at different depths
type transitionKey struct {
	from  byte
	to    byte
	depth int
}

var transitionCache = make(map[transitionKey]uint)

func solveCode(code string, depth int) uint {
	numericKeypad := map[byte]image.Point{
		'7': {0, 0}, '8': {1, 0}, '9': {2, 0},
		'4': {0, 1}, '5': {1, 1}, '6': {2, 1},
		'1': {0, 2}, '2': {1, 2}, '3': {2, 2},
		'0': {1, 3}, 'A': {2, 3},
	}

	var totalLength uint
	currentKey := byte('A')

	for i := range code {
		targetKey := code[i]
		fromPos := numericKeypad[currentKey]
		toPos := numericKeypad[targetKey]
		path := findPathNumeric(fromPos, toPos) + "A"

		// Count cost of typing this path through 'depth' directional keypads
		if depth == 0 {
			totalLength += uint(len(path))
		} else {
			subKey := byte('A')
			for j := range path {
				subTarget := path[j]
				totalLength += solveDirectionalTransition(subKey, subTarget, depth)
				subKey = subTarget
			}
		}
		currentKey = targetKey
	}

	return totalLength
}

func solveDirectionalTransition(from, to byte, depth int) uint {
	key := transitionKey{from, to, depth}
	if cached, ok := transitionCache[key]; ok {
		return cached
	}

	directionalKeypad := map[byte]image.Point{
		'^': {1, 0}, 'A': {2, 0},
		'<': {0, 1}, 'v': {1, 1}, '>': {2, 1},
	}

	fromPos := directionalKeypad[from]
	toPos := directionalKeypad[to]
	path := findPathDirectional(fromPos, toPos) + "A"

	var length uint
	if depth == 1 {
		// Base case: just count the path length
		length = uint(len(path))
	} else {
		// Recursive case: count cost of typing path at next depth level
		currentKey := byte('A')
		for i := range path {
			targetKey := path[i]
			length += solveDirectionalTransition(currentKey, targetKey, depth-1)
			currentKey = targetKey
		}
	}

	transitionCache[key] = length
	return length
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