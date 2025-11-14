package adventofcode2024

import (
	"testing"
)

func TestDay08Part1Example(t *testing.T) {
	testLines(t, 8, exampleFilename, true, Day08, 14)
}

func TestDay08Part2Example(t *testing.T) {
	testLines(t, 8, exampleFilename, false, Day08, 34)
}

func TestDay08Part1(t *testing.T) {
	testLines(t, 8, filename, true, Day08, 291)
}

func TestDay08Part2(t *testing.T) {
	testLines(t, 8, filename, false, Day08, 1015)
}

func BenchmarkDay08Part1(b *testing.B) {
	benchLines(b, 8, true, Day08)
}

func BenchmarkDay08Part2(b *testing.B) {
	benchLines(b, 8, false, Day08)
}
