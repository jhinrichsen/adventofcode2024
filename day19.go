package adventofcode2024

import (
	"strings"
)

type Day19Puzzle struct {
	patterns []string
	designs  []string
}

func NewDay19(lines []string) Day19Puzzle {
	var p Day19Puzzle

	// First line contains available patterns (comma-separated)
	if len(lines) > 0 {
		p.patterns = strings.Split(strings.TrimSpace(lines[0]), ", ")
	}

	// Skip empty line (line 1), designs start from line 2
	for i := 2; i < len(lines); i++ {
		if line := strings.TrimSpace(lines[i]); line != "" {
			p.designs = append(p.designs, line)
		}
	}

	return p
}

func Day19(puzzle Day19Puzzle, part1 bool) (uint, string) {
	if part1 {
		count := countPossibleDesigns(puzzle.patterns, puzzle.designs)
		return count, ""
	}
	// Part 2: Count total number of ways to make all designs
	totalWays := countAllWaysToMakeDesigns(puzzle.patterns, puzzle.designs)
	return totalWays, ""
}

func countPossibleDesigns(patterns []string, designs []string) uint {
	var count uint

	for _, design := range designs {
		if canMakeDesign(design, patterns) {
			count++
		}
	}

	return count
}

func canMakeDesign(design string, patterns []string) bool {
	// Use dynamic programming - iterative bottom-up approach
	n := len(design)
	dp := make([]bool, n+1)
	dp[0] = true // Empty string can always be made

	for i := 1; i <= n; i++ {
		for _, pattern := range patterns {
			patternLen := len(pattern)
			if i >= patternLen && design[i-patternLen:i] == pattern {
				if dp[i-patternLen] {
					dp[i] = true
					break
				}
			}
		}
	}

	return dp[n]
}

func countAllWaysToMakeDesigns(patterns []string, designs []string) uint {
	var total uint

	for _, design := range designs {
		ways := countWaysToMakeDesign(design, patterns)
		total += ways
	}

	return total
}

func countWaysToMakeDesign(design string, patterns []string) uint {
	// Use dynamic programming - count number of ways to make each prefix
	n := len(design)
	dp := make([]uint, n+1)
	dp[0] = 1 // One way to make empty string

	for i := 1; i <= n; i++ {
		for _, pattern := range patterns {
			patternLen := len(pattern)
			if i >= patternLen && design[i-patternLen:i] == pattern {
				dp[i] += dp[i-patternLen]
			}
		}
	}

	return dp[n]
}