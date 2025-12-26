package adventofcode2024

import (
	"strings"
)

type Day19Puzzle struct {
	trie    []day19Node // Trie for patterns
	designs []string
}

// Trie node with 5 children (w=0, u=1, b=2, r=3, g=4)
type day19Node struct {
	children [5]int // index into trie, 0 = no child
	isEnd    bool
}

func day19CharIndex(c byte) int {
	switch c {
	case 'w':
		return 0
	case 'u':
		return 1
	case 'b':
		return 2
	case 'r':
		return 3
	case 'g':
		return 4
	}
	return -1
}

func NewDay19(lines []string) (Day19Puzzle, error) {
	var p Day19Puzzle

	// Initialize trie with root node
	p.trie = make([]day19Node, 1, 1024)

	// First line contains available patterns (comma-separated)
	if len(lines) > 0 {
		patterns := strings.Split(strings.TrimSpace(lines[0]), ", ")
		for _, pattern := range patterns {
			p.addPattern(pattern)
		}
	}

	// Skip empty line (line 1), designs start from line 2
	for i := 2; i < len(lines); i++ {
		if line := strings.TrimSpace(lines[i]); line != "" {
			p.designs = append(p.designs, line)
		}
	}

	return p, nil
}

func (p *Day19Puzzle) addPattern(pattern string) {
	node := 0
	for i := range pattern {
		idx := day19CharIndex(pattern[i])
		if idx < 0 {
			return
		}
		if p.trie[node].children[idx] == 0 {
			p.trie[node].children[idx] = len(p.trie)
			p.trie = append(p.trie, day19Node{})
		}
		node = p.trie[node].children[idx]
	}
	p.trie[node].isEnd = true
}

func Day19(puzzle Day19Puzzle, part1 bool) (uint, string) {
	if part1 {
		var count uint
		for _, design := range puzzle.designs {
			if canMakeDesignTrie(design, puzzle.trie) {
				count++
			}
		}
		return count, ""
	}
	// Part 2
	var total uint
	for _, design := range puzzle.designs {
		total += countWaysTrie(design, puzzle.trie)
	}
	return total, ""
}

func canMakeDesignTrie(design string, trie []day19Node) bool {
	n := len(design)
	dp := make([]bool, n+1)
	dp[0] = true

	for i := 0; i < n; i++ {
		if !dp[i] {
			continue
		}
		// Walk trie from position i
		node := 0
		for j := i; j < n; j++ {
			idx := day19CharIndex(design[j])
			if idx < 0 || trie[node].children[idx] == 0 {
				break
			}
			node = trie[node].children[idx]
			if trie[node].isEnd {
				dp[j+1] = true
			}
		}
	}

	return dp[n]
}

func countWaysTrie(design string, trie []day19Node) uint {
	n := len(design)
	dp := make([]uint, n+1)
	dp[0] = 1

	for i := 0; i < n; i++ {
		if dp[i] == 0 {
			continue
		}
		// Walk trie from position i
		node := 0
		for j := i; j < n; j++ {
			idx := day19CharIndex(design[j])
			if idx < 0 || trie[node].children[idx] == 0 {
				break
			}
			node = trie[node].children[idx]
			if trie[node].isEnd {
				dp[j+1] += dp[i]
			}
		}
	}

	return dp[n]
}
