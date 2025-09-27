package adventofcode2024


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

func Day11(stones []uint64, part1 bool) uint64 {
	blinks := 75
	if part1 {
		blinks = 25
	}

	stoneCounts := make(map[uint64]uint64)
	for _, stone := range stones {
		stoneCounts[stone]++
	}

	for i := 0; i < blinks; i++ {
		newCounts := make(map[uint64]uint64)
		for stone, count := range stoneCounts {
			if stone == 0 {
				newCounts[1] += count
			} else if high, low, ok := split(stone); ok {
				newCounts[high] += count
				newCounts[low] += count
			} else {
				newCounts[stone*2024] += count
			}
		}
		stoneCounts = newCounts
	}

	var total uint64
	for _, count := range stoneCounts {
		total += count
	}

	return total
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
// Each comparison is O(1) and branchless, covering uint64 range up to 10^19
func digits_branchless(n uint64) int {
	d := 1

	// Branchless comparisons using bit manipulation
	d += int((n-1e1)>>63) ^ 1  // >= 10^1
	d += int((n-1e2)>>63) ^ 1  // >= 10^2
	d += int((n-1e3)>>63) ^ 1  // >= 10^3
	d += int((n-1e4)>>63) ^ 1  // >= 10^4
	d += int((n-1e5)>>63) ^ 1  // >= 10^5
	d += int((n-1e6)>>63) ^ 1  // >= 10^6
	d += int((n-1e7)>>63) ^ 1  // >= 10^7
	d += int((n-1e8)>>63) ^ 1  // >= 10^8
	d += int((n-1e9)>>63) ^ 1  // >= 10^9
	d += int((n-1e10)>>63) ^ 1 // >= 10^10
	d += int((n-1e11)>>63) ^ 1 // >= 10^11
	d += int((n-1e12)>>63) ^ 1 // >= 10^12
	d += int((n-1e13)>>63) ^ 1 // >= 10^13
	d += int((n-1e14)>>63) ^ 1 // >= 10^14
	d += int((n-1e15)>>63) ^ 1 // >= 10^15
	d += int((n-1e16)>>63) ^ 1 // >= 10^16
	d += int((n-1e17)>>63) ^ 1 // >= 10^17
	d += int((n-1e18)>>63) ^ 1 // >= 10^18
	// Note: 10^19 > uint64 max (18446744073709551615), so we stop at 10^18

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
		1e0, 1e1, 1e2, 1e3, 1e4, 1e5, 1e6, 1e7, 1e8, 1e9,
		1e10, 1e11, 1e12, 1e13, 1e14, 1e15, 1e16, 1e17, 1e18, 1e19,
	}

	pow := powers[half]
	high = n / pow
	low = n % pow
	ok = true
	return
}
