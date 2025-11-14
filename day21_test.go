package adventofcode2024

import (
	"testing"
)

func TestDay21Part1Example(t *testing.T) {
	testWithParser(t, 21, exampleFilename, true, NewDay21, Day21, 126384)
}

func TestDay21Part1ExampleIndividual(t *testing.T) {
	tests := []struct {
		code           string
		wantSeqLength  uint
		wantComplexity uint
	}{
		{"029A", 68, 1972},   // 68 * 29 = 1972
		{"980A", 60, 58800},  // 60 * 980 = 58800
		{"179A", 68, 12172},  // 68 * 179 = 12172
		{"456A", 64, 29184},  // 64 * 456 = 29184
		{"379A", 64, 24256},  // 64 * 379 = 24256
	}

	for _, tt := range tests {
		t.Run(tt.code, func(t *testing.T) {
			gotSeqLength := solveCode(tt.code, 2)
			gotComplexity := gotSeqLength * extractNumericValue(tt.code)

			if gotSeqLength != tt.wantSeqLength {
				t.Errorf("sequence length: want %d, got %d", tt.wantSeqLength, gotSeqLength)
			}
			if gotComplexity != tt.wantComplexity {
				t.Errorf("complexity: want %d, got %d", tt.wantComplexity, gotComplexity)
			}
		})
	}
}

func TestDay21Part1(t *testing.T) {
	testWithParser(t, 21, filename, true, NewDay21, Day21, 157892)
}

func TestDay21Part2(t *testing.T) {
	testWithParser(t, 21, filename, false, NewDay21, Day21, 197015606336332)
}

func BenchmarkDay21Part1(b *testing.B) {
	benchWithParser(b, 21, true, NewDay21, Day21)
}

func BenchmarkDay21Part2(b *testing.B) {
	benchWithParser(b, 21, false, NewDay21, Day21)
}
