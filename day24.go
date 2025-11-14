package adventofcode2024

import (
	"slices"
	"strconv"
	"strings"
)

type Day24Puzzle struct {
	wires map[string]int // wire name -> value (0, 1, or -1 for uncomputed)
	gates []Gate
}

type Gate struct {
	input1 string
	input2 string
	op     string // "AND", "OR", "XOR"
	output string
}

func NewDay24(lines []string) Day24Puzzle {
	wires := make(map[string]int)
	var gates []Gate

	parsingWires := true
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			parsingWires = false
			continue
		}

		if parsingWires {
			// Parse initial wire values: "x00: 1"
			parts := strings.Split(line, ": ")
			if len(parts) == 2 {
				wire := parts[0]
				value, _ := strconv.Atoi(parts[1])
				wires[wire] = value
			}
		} else {
			// Parse gates: "x00 AND y00 -> z00"
			parts := strings.Split(line, " -> ")
			if len(parts) == 2 {
				gateParts := strings.Fields(parts[0])
				if len(gateParts) == 3 {
					gates = append(gates, Gate{
						input1: gateParts[0],
						input2: gateParts[2],
						op:     gateParts[1],
						output: parts[1],
					})
				}
			}
		}
	}

	return Day24Puzzle{wires: wires, gates: gates}
}

func Day24(puzzle Day24Puzzle, part1 bool) string {
	if part1 {
		result := simulateCircuit(puzzle)
		return strconv.FormatUint(result, 10)
	}
	return "0"
}

func simulateCircuit(puzzle Day24Puzzle) uint64 {
	// Initialize all wires mentioned in gates as uncomputed (-1)
	// except undefined wires which default to 0
	for _, gate := range puzzle.gates {
		if _, exists := puzzle.wires[gate.output]; !exists {
			puzzle.wires[gate.output] = -1
		}
		// Don't initialize input wires - they might be undefined/missing
	}

	// Simulate until no more changes
	changed := true
	for changed {
		changed = false
		for _, gate := range puzzle.gates {
			// Skip if output already computed
			if puzzle.wires[gate.output] != -1 {
				continue
			}

			// Get input values (default to 0 if not found)
			val1, ok1 := puzzle.wires[gate.input1]
			if !ok1 {
				val1 = 0
			}
			val2, ok2 := puzzle.wires[gate.input2]
			if !ok2 {
				val2 = 0
			}

			// Skip if inputs not ready
			if val1 == -1 || val2 == -1 {
				continue
			}

			// Compute output
			var result int
			switch gate.op {
			case "AND":
				if val1 == 1 && val2 == 1 {
					result = 1
				} else {
					result = 0
				}
			case "OR":
				if val1 == 1 || val2 == 1 {
					result = 1
				} else {
					result = 0
				}
			case "XOR":
				if val1 != val2 {
					result = 1
				} else {
					result = 0
				}
			}

			puzzle.wires[gate.output] = result
			changed = true
		}
	}

	// Extract z-wires and build number (z00 is LSB)
	var zWires []string
	for wire := range puzzle.wires {
		if strings.HasPrefix(wire, "z") {
			zWires = append(zWires, wire)
		}
	}
	slices.Sort(zWires)

	var result uint64
	for i, wire := range zWires {
		if puzzle.wires[wire] == 1 {
			result |= (1 << i)
		}
	}

	return result
}
