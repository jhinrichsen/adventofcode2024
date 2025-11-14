package adventofcode2024

import (
	"strconv"
	"strings"
)

type Day22Puzzle struct {
	secrets []uint
}

func NewDay22(lines []string) (Day22Puzzle, error) {
	secrets := make([]uint, 0, len(lines))
	for _, line := range lines {
		if line := strings.TrimSpace(line); line != "" {
			val, err := strconv.ParseUint(line, 10, 64)
			if err != nil {
				return Day22Puzzle{}, err
			}
			secrets = append(secrets, uint(val))
		}
	}
	return Day22Puzzle{secrets: secrets}, nil
}

func Day22(puzzle Day22Puzzle, part1 bool) uint {
	if part1 {
		return solveDay22Part1(puzzle)
	}
	return solveDay22Part2(puzzle)
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

func solveDay22Part2(puzzle Day22Puzzle) uint {
	// Map from sequence (as [4]int8) to total bananas across all buyers
	sequenceTotals := make(map[[4]int8]uint)

	for _, initialSecret := range puzzle.secrets {
		// Generate prices and changes for this buyer
		prices := make([]uint, 2001)
		prices[0] = initialSecret % 10

		secret := initialSecret
		for i := 1; i <= 2000; i++ {
			secret = nextSecret(secret)
			prices[i] = secret % 10
		}

		// Track which sequences we've seen for this buyer (first occurrence only)
		seen := make(map[[4]int8]bool)

		// Build sequences of 4 consecutive changes
		for i := 4; i < len(prices); i++ {
			sequence := [4]int8{
				int8(prices[i-3]) - int8(prices[i-4]),
				int8(prices[i-2]) - int8(prices[i-3]),
				int8(prices[i-1]) - int8(prices[i-2]),
				int8(prices[i]) - int8(prices[i-1]),
			}

			// Only count first occurrence for this buyer
			if !seen[sequence] {
				seen[sequence] = true
				sequenceTotals[sequence] += prices[i]
			}
		}
	}

	// Find the maximum total
	var maxBananas uint
	for _, total := range sequenceTotals {
		if total > maxBananas {
			maxBananas = total
		}
	}

	return maxBananas
}
