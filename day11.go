package adventofcode2024

import "slices"

func NewDay11(input string) []uint64 {
	stones := make([]uint64, 0, len(input)/2) // worst case: single digits separated by one space
	var n uint64

	for _, c := range input {
		if c >= '0' && c <= '9' {
			n = 10*n + uint64(c-'0')
		} else {
			stones = append(stones, n)
			n = 0
		}
	}

	return stones
}

func Day11(stones []uint64) uint {
	const blinks = 25
	for range blinks {
		stones = blink(stones)
	}
	return uint(len(stones))
}

func blink(stones []uint64) []uint64 {
	for i := len(stones) - 1; i >= 0; i-- {
		stone := stones[i]
		// "If the stone is engraved with the number 0, it is replaced by a stone engraved with the number 1."
		if stone == 0 {
			stones[i] = 1
		} else if a, b, ok := split(stone); ok {
			stones[i] = a
			stones = slices.Insert(stones, i+1, b)
		} else {
			stones[i] *= 2024
		}
	}
	return stones
}

func evenNumberOfDigits(n uint64) bool {
	return digits_branchless(n)%2 == 0
}

func countDigits(n uint64) int {
	if n == 0 {
		return 1
	}
	count := 0
	for n > 0 {
		n /= 10
		count++
	}
	return count
}

// Branchless digit counting using bit manipulation
// Each comparison is O(1) and branchless, totaling ~11 operations
func digits_branchless(n uint64) int {
	d := 1

	// Branchless comparisons using bit manipulation
	d += int((n-10)>>63) ^ 1           // >= 10
	d += int((n-100)>>63) ^ 1          // >= 100
	d += int((n-1000)>>63) ^ 1         // >= 1000
	d += int((n-10000)>>63) ^ 1        // >= 10000
	d += int((n-100000)>>63) ^ 1       // >= 100000
	d += int((n-1000000)>>63) ^ 1      // >= 1000000
	d += int((n-10000000)>>63) ^ 1     // >= 10000000
	d += int((n-100000000)>>63) ^ 1    // >= 100000000
	d += int((n-1000000000)>>63) ^ 1   // >= 1000000000
	d += int((n-10000000000)>>63) ^ 1  // >= 10000000000
	d += int((n-100000000000)>>63) ^ 1 // >= 100000000000

	return d
}

// Split even-digit number into high and low halves using precomputed powers
func split(n uint64) (high, low uint64, ok bool) {
	if !evenNumberOfDigits(n) {
		return 0, 0, false
	}

	d := digits_branchless(n)
	half := d / 2

	// Precomputed powers of 10 up to 10^19
	powers := [20]uint64{
		1, 10, 100, 1_000, 10_000, 100_000, 1_000_000, 10_000_000, 100_000_000,
		1_000_000_000, 10_000_000_000, 100_000_000_000, 1_000_000_000_000,
		10_000_000_000_000, 100_000_000_000_000, 1_000_000_000_000_000,
		10_000_000_000_000_000, 100_000_000_000_000_000,
		1_000_000_000_000_000_000, 10_000_000_000_000_000_000,
	}

	pow := powers[half]
	high = n / pow
	low = n % pow
	ok = true
	return
}
