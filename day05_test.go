package adventofcode2024

import (
	"testing"
)

func TestDay05Part1Example(t *testing.T) {
	testLines(t, 5, exampleFilename, true, Day05, 143)
}

func TestDay05Part2Example(t *testing.T) {
	testLines(t, 5, exampleFilename, false, Day05, 123)
}

func TestDay05Part1(t *testing.T) {
	testLines(t, 5, filename, true, Day05, 5108)
}

func TestDay05Part2(t *testing.T) {
	testLines(t, 5, filename, false, Day05, 7380)
}

func BenchmarkDay05Part1(b *testing.B) {
	benchLines(b, 5, true, Day05)
}

func BenchmarkDay05Part2(b *testing.B) {
	benchLines(b, 5, false, Day05)
}
