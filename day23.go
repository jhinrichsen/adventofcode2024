package adventofcode2024

import (
	"slices"
	"strconv"
	"strings"
)

type Day23Puzzle struct {
	adj       []uint64     // adjacency bitset: adj[i*11 + j/64] bit j%64
	nodeCount int          // number of unique nodes
	idToName  []string     // id -> "xx" name
	nameToID  map[string]int
	tNodes    []int        // nodes starting with 't'
}

const day23MaxNodes = 676 // 26*26
const day23Words = (day23MaxNodes + 63) / 64 // 11 words per node

func NewDay23(lines []string) (Day23Puzzle, error) {
	nameToID := make(map[string]int)
	var idToName []string
	var edges [][2]int
	var tNodes []int

	for _, line := range lines {
		if line = strings.TrimSpace(line); line == "" {
			continue
		}
		// Parse "xx-yy"
		a, b := line[:2], line[3:5]

		idA, ok := nameToID[a]
		if !ok {
			idA = len(idToName)
			nameToID[a] = idA
			idToName = append(idToName, a)
			if a[0] == 't' {
				tNodes = append(tNodes, idA)
			}
		}

		idB, ok := nameToID[b]
		if !ok {
			idB = len(idToName)
			nameToID[b] = idB
			idToName = append(idToName, b)
			if b[0] == 't' {
				tNodes = append(tNodes, idB)
			}
		}

		edges = append(edges, [2]int{idA, idB})
	}

	n := len(idToName)
	// Adjacency bitset: n nodes, each with day23Words uint64s
	adj := make([]uint64, n*day23Words)

	for _, e := range edges {
		a, b := e[0], e[1]
		// Set bit b in node a's adjacency
		adj[a*day23Words+b/64] |= 1 << (b % 64)
		// Set bit a in node b's adjacency
		adj[b*day23Words+a/64] |= 1 << (a % 64)
	}

	return Day23Puzzle{
		adj:       adj,
		nodeCount: n,
		idToName:  idToName,
		nameToID:  nameToID,
		tNodes:    tNodes,
	}, nil
}

func (p *Day23Puzzle) connected(a, b int) bool {
	return (p.adj[a*day23Words+b/64] & (1 << (b % 64))) != 0
}

func Day23(puzzle Day23Puzzle, part1 bool) string {
	if part1 {
		return strconv.Itoa(countTrianglesWithT(puzzle))
	}
	return findMaxClique(puzzle)
}

func countTrianglesWithT(p Day23Puzzle) int {
	count := 0
	n := p.nodeCount

	// For each node starting with 't', find triangles
	for _, t := range p.tNodes {
		// Get neighbors of t
		for a := 0; a < n; a++ {
			if a == t || !p.connected(t, a) {
				continue
			}
			for b := a + 1; b < n; b++ {
				if b == t || !p.connected(t, b) {
					continue
				}
				if p.connected(a, b) {
					// Triangle t-a-b, but only count if t is smallest 't' node
					// to avoid duplicates when multiple 't' nodes in triangle
					hasSmaller := false
					for _, other := range p.tNodes {
						if other < t && (other == a || other == b) {
							hasSmaller = true
							break
						}
					}
					if !hasSmaller {
						count++
					}
				}
			}
		}
	}
	return count
}

func findMaxClique(p Day23Puzzle) string {
	n := p.nodeCount

	// Bron-Kerbosch with pivoting, iterative with stack
	type state struct {
		R []int   // current clique
		P []int   // candidates
		X []int   // excluded
	}

	// Initial state: R={}, P=all nodes, X={}
	allNodes := make([]int, n)
	for i := range allNodes {
		allNodes[i] = i
	}

	stack := []state{{R: nil, P: allNodes, X: nil}}
	var maxClique []int

	for len(stack) > 0 {
		s := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if len(s.P) == 0 && len(s.X) == 0 {
			if len(s.R) > len(maxClique) {
				maxClique = append([]int(nil), s.R...)
			}
			continue
		}

		if len(s.P) == 0 {
			continue
		}

		// Choose pivot from P ∪ X with most connections to P
		pivot := -1
		maxConn := -1
		for _, u := range s.P {
			conn := 0
			for _, v := range s.P {
				if p.connected(u, v) {
					conn++
				}
			}
			if conn > maxConn {
				maxConn = conn
				pivot = u
			}
		}
		for _, u := range s.X {
			conn := 0
			for _, v := range s.P {
				if p.connected(u, v) {
					conn++
				}
			}
			if conn > maxConn {
				maxConn = conn
				pivot = u
			}
		}

		// Process P \ N(pivot)
		for i := len(s.P) - 1; i >= 0; i-- {
			v := s.P[i]
			if pivot >= 0 && p.connected(pivot, v) {
				continue
			}

			// newR = R ∪ {v}
			newR := append(append([]int(nil), s.R...), v)

			// newP = P ∩ N(v)
			var newP []int
			for _, u := range s.P {
				if u != v && p.connected(v, u) {
					newP = append(newP, u)
				}
			}

			// newX = X ∩ N(v)
			var newX []int
			for _, u := range s.X {
				if p.connected(v, u) {
					newX = append(newX, u)
				}
			}

			stack = append(stack, state{R: newR, P: newP, X: newX})

			// P = P \ {v}, X = X ∪ {v}
			s.P = append(s.P[:i], s.P[i+1:]...)
			s.X = append(s.X, v)
		}
	}

	// Convert to names and sort
	names := make([]string, len(maxClique))
	for i, id := range maxClique {
		names[i] = p.idToName[id]
	}
	slices.Sort(names)
	return strings.Join(names, ",")
}
