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
	// Part 2: Find swapped wires in the adder circuit
	swapped := findSwappedWires(puzzle)
	slices.Sort(swapped)
	return strings.Join(swapped, ",")
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

func findSwappedWires(puzzle Day24Puzzle) []string {
	swappedSet := make(map[string]bool)

	// Find the highest z wire (final carry bit)
	maxZ := ""
	for i := range puzzle.gates {
		if strings.HasPrefix(puzzle.gates[i].output, "z") {
			if maxZ == "" || puzzle.gates[i].output > maxZ {
				maxZ = puzzle.gates[i].output
			}
		}
	}

	// For a ripple-carry adder, check rules:
	// 1. All z outputs (except last) must be from XOR gates
	// 2. XOR gates with x,y inputs should not output to z (except z00)
	// 3. AND gates (except x00,y00) must feed into OR gates
	// 4. XOR gates (except x,y inputs) must output to z wires
	// 5. OR gates should not output to z wires (except possibly final carry)
	// 6. XOR with x,y inputs must feed into both XOR and AND gates

	for i := range puzzle.gates {
		gate := &puzzle.gates[i]
		out := gate.output
		op := gate.op
		in1 := gate.input1
		in2 := gate.input2

		// Rule 1: z outputs (except last) must be from XOR
		if strings.HasPrefix(out, "z") && out != maxZ {
			if op != "XOR" {
				swappedSet[out] = true
			}
		}

		// Rule 2: XOR with x,y inputs should not output to z (except z00)
		if op == "XOR" {
			isXYInput := (strings.HasPrefix(in1, "x") || strings.HasPrefix(in1, "y")) &&
				(strings.HasPrefix(in2, "x") || strings.HasPrefix(in2, "y"))
			if isXYInput && strings.HasPrefix(out, "z") && out != "z00" {
				swappedSet[out] = true
			}
		}

		// Rule 3: AND outputs (except x00 AND y00) must feed into OR
		if op == "AND" && !(in1 == "x00" && in2 == "y00") && !(in1 == "y00" && in2 == "x00") {
			// Check if this output feeds into an OR gate
			feedsOR := false
			for j := range puzzle.gates {
				if puzzle.gates[j].op == "OR" {
					if puzzle.gates[j].input1 == out || puzzle.gates[j].input2 == out {
						feedsOR = true
						break
					}
				}
			}
			if !feedsOR {
				swappedSet[out] = true
			}
		}

		// Rule 4: XOR gates without x,y direct inputs must output to z
		if op == "XOR" {
			isXYInput := (strings.HasPrefix(in1, "x") || strings.HasPrefix(in1, "y")) &&
				(strings.HasPrefix(in2, "x") || strings.HasPrefix(in2, "y"))
			if !isXYInput && !strings.HasPrefix(out, "z") {
				swappedSet[out] = true
			}
		}

		// Rule 5: OR gates should not output to z wires (except possibly the final carry)
		if op == "OR" && strings.HasPrefix(out, "z") && out != maxZ {
			swappedSet[out] = true
		}

		// Rule 6: XOR with x,y inputs (except z00) must feed into both XOR and AND
		if op == "XOR" && out != "z00" {
			isXYInput := (strings.HasPrefix(in1, "x") || strings.HasPrefix(in1, "y")) &&
				(strings.HasPrefix(in2, "x") || strings.HasPrefix(in2, "y"))
			if isXYInput {
				feedsXOR := false
				feedsAND := false
				for j := range puzzle.gates {
					if puzzle.gates[j].input1 == out || puzzle.gates[j].input2 == out {
						if puzzle.gates[j].op == "XOR" {
							feedsXOR = true
						}
						if puzzle.gates[j].op == "AND" {
							feedsAND = true
						}
					}
				}
				if !feedsXOR || !feedsAND {
					swappedSet[out] = true
				}
			}
		}
	}

	// Convert set to list
	var result []string
	for wire := range swappedSet {
		result = append(result, wire)
	}
	return result
}
