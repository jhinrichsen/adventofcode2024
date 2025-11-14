package adventofcode2024

import (
	"strconv"
	"strings"
)

type Day22Puzzle struct {
	secrets []uint
}

func NewDay22(lines []string) Day22Puzzle {
	secrets := make([]uint, 0, len(lines))
	for _, line := range lines {
		if line := strings.TrimSpace(line); line != "" {
			val, _ := strconv.ParseUint(line, 10, 64)
			secrets = append(secrets, uint(val))
		}
	}
	return Day22Puzzle{secrets: secrets}
}

func Day22(puzzle Day22Puzzle, part1 bool) uint {
	if part1 {
		return solveDay22Part1(puzzle)
	}
	return 0
}

func solveDay22Part1(puzzle Day22Puzzle) uint {
	var sum uint
	for _, secret := range puzzle.secrets {
		finalSecret := evolveSecret(secret, 2000)
		sum += finalSecret
	}
	return sum
}

func evolveSecret(secret uint, iterations int) uint {
	for range iterations {
		secret = nextSecret(secret)
	}
	return secret
}

func nextSecret(secret uint) uint {
	// Step 1: multiply by 64, mix, prune
	secret = prune(mix(secret, secret*64))

	// Step 2: divide by 32, mix, prune
	secret = prune(mix(secret, secret/32))

	// Step 3: multiply by 2048, mix, prune
	secret = prune(mix(secret, secret*2048))

	return secret
}

func mix(secret, value uint) uint {
	return secret ^ value
}

func prune(secret uint) uint {
	return secret % 16777216
}
