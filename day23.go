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
	// Part 2 returns string, not uint - return 0 and use Day23Part2Password() for actual result
	return 0
}

// Day23Part2Password returns the password for the LAN party (largest clique, comma-separated)
func Day23Part2Password(puzzle Day23Puzzle) string {
	maxClique := findMaximumClique(puzzle)
	slices.Sort(maxClique)
	return strings.Join(maxClique, ",")
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

// findMaximumClique finds the largest clique using Bron-Kerbosch algorithm
func findMaximumClique(puzzle Day23Puzzle) []string {
	var maxClique []string

	// Convert computers slice to set for easier manipulation
	allComputers := make(map[string]bool)
	for _, comp := range puzzle.computers {
		allComputers[comp] = true
	}

	// Bron-Kerbosch algorithm
	bronKerbosch(
		puzzle,
		make(map[string]bool),  // R: current clique
		allComputers,            // P: candidates
		make(map[string]bool),  // X: already processed
		&maxClique,
	)

	return maxClique
}

// bronKerbosch implements the Bron-Kerbosch algorithm for finding maximal cliques
func bronKerbosch(
	puzzle Day23Puzzle,
	R map[string]bool,  // current clique
	P map[string]bool,  // candidates
	X map[string]bool,  // already processed
	maxClique *[]string,
) {
	if len(P) == 0 && len(X) == 0 {
		// Found a maximal clique
		if len(R) > len(*maxClique) {
			*maxClique = make([]string, 0, len(R))
			for node := range R {
				*maxClique = append(*maxClique, node)
			}
		}
		return
	}

	// Make a copy of P to iterate over (since we modify P in the loop)
	PCopy := make([]string, 0, len(P))
	for node := range P {
		PCopy = append(PCopy, node)
	}

	for _, v := range PCopy {
		// R ∪ {v}
		newR := make(map[string]bool, len(R)+1)
		for node := range R {
			newR[node] = true
		}
		newR[v] = true

		// P ∩ N(v)
		neighbors := puzzle.connections[v]
		newP := make(map[string]bool)
		for _, neighbor := range neighbors {
			if P[neighbor] {
				newP[neighbor] = true
			}
		}

		// X ∩ N(v)
		newX := make(map[string]bool)
		for _, neighbor := range neighbors {
			if X[neighbor] {
				newX[neighbor] = true
			}
		}

		bronKerbosch(puzzle, newR, newP, newX, maxClique)

		// P := P \ {v}
		delete(P, v)

		// X := X ∪ {v}
		X[v] = true
	}
}

