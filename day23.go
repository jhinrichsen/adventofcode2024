package adventofcode2024

import (
	"slices"
	"strings"
)

type Day23Puzzle struct {
	connections map[string][]string
	computers   []string
}

func NewDay23(lines []string) Day23Puzzle {
	connections := make(map[string][]string)
	computerSet := make(map[string]bool)

	for _, line := range lines {
		if line := strings.TrimSpace(line); line != "" {
			parts := strings.Split(line, "-")
			if len(parts) == 2 {
				a, b := parts[0], parts[1]
				connections[a] = append(connections[a], b)
				connections[b] = append(connections[b], a)
				computerSet[a] = true
				computerSet[b] = true
			}
		}
	}

	// Convert set to sorted slice for deterministic iteration
	computers := make([]string, 0, len(computerSet))
	for comp := range computerSet {
		computers = append(computers, comp)
	}
	slices.Sort(computers)

	return Day23Puzzle{
		connections: connections,
		computers:   computers,
	}
}

func Day23(puzzle Day23Puzzle, part1 bool) uint {
	if part1 {
		return solveDay23Part1(puzzle)
	}
	return 0
}

func solveDay23Part1(puzzle Day23Puzzle) uint {
	triangles := findTriangles(puzzle)

	// Count triangles with at least one computer starting with 't'
	var count uint
	for triangle := range triangles {
		if hasComputerStartingWithT(triangle) {
			count++
		}
	}

	return count
}

func findTriangles(puzzle Day23Puzzle) map[[3]string]bool {
	triangles := make(map[[3]string]bool)

	// For each computer, check all pairs of its neighbors
	for _, a := range puzzle.computers {
		neighbors := puzzle.connections[a]

		for i := range neighbors {
			for j := i + 1; j < len(neighbors); j++ {
				b := neighbors[i]
				c := neighbors[j]

				// Check if b and c are connected
				if isConnected(puzzle, b, c) {
					// Create canonical triangle (sorted)
					triangle := [3]string{a, b, c}
					slices.Sort(triangle[:])
					triangles[triangle] = true
				}
			}
		}
	}

	return triangles
}

func isConnected(puzzle Day23Puzzle, a, b string) bool {
	neighbors := puzzle.connections[a]
	return slices.Contains(neighbors, b)
}

func hasComputerStartingWithT(triangle [3]string) bool {
	for _, comp := range triangle {
		if len(comp) > 0 && comp[0] == 't' {
			return true
		}
	}
	return false
}
