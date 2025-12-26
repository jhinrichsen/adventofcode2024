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
	// Encode 4 changes (each -9 to 9) as single index
	// (c0+9)*19^3 + (c1+9)*19^2 + (c2+9)*19 + (c3+9)
	// Max: 19^4 = 130321 possible sequences
	const seqSize = 19 * 19 * 19 * 19 // 130321

	sequenceTotals := make([]uint, seqSize)
	seen := make([]bool, seqSize)

	for _, initialSecret := range puzzle.secrets {
		// Clear seen for this buyer
		clear(seen)

		// Generate first 4 prices to bootstrap
		secret := initialSecret
		p0 := secret % 10
		secret = nextSecret(secret)
		p1 := secret % 10
		secret = nextSecret(secret)
		p2 := secret % 10
		secret = nextSecret(secret)
		p3 := secret % 10

		// Process remaining prices
		for range 1997 { // 2000 - 3 already done
			secret = nextSecret(secret)
			p4 := secret % 10

			// Encode sequence of 4 changes as index
			// Changes are in [-9, 9], offset by 9 to get [0, 18]
			idx := (int(p1)-int(p0)+9)*6859 + // 19^3
				(int(p2)-int(p1)+9)*361 + // 19^2
				(int(p3)-int(p2)+9)*19 +
				(int(p4) - int(p3) + 9)

			if !seen[idx] {
				seen[idx] = true
				sequenceTotals[idx] += p4
			}

			// Shift window
			p0, p1, p2, p3 = p1, p2, p3, p4
		}
	}

	// Find maximum
	var maxBananas uint
	for _, total := range sequenceTotals {
		if total > maxBananas {
			maxBananas = total
		}
	}

	return maxBananas
}
