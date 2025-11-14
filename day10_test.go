package adventofcode2024

import (
	"testing"
)

func TestDay10Part1Examples(t *testing.T) {
	tests := []struct {
		name string
		file string
		want uint
	}{
		{"example1", example1Filename(10), 2},
		{"example2", example2Filename(10), 4},
		{"example3", example3Filename(10), 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testWithParserNoErr(t, 10, func(uint8) string { return tt.file }, true, NewDay10, Day10, tt.want)
		})
	}
}

func TestDay10Part1Example(t *testing.T) {
	testWithParserNoErr(t, 10, exampleFilename, true, NewDay10, Day10, 36)
}

func TestDay10Part2Example(t *testing.T) {
	testWithParserNoErr(t, 10, exampleFilename, false, NewDay10, Day10, 81)
}

func TestDay10Part1(t *testing.T) {
	testWithParserNoErr(t, 10, filename, true, NewDay10, Day10, 587)
}

func TestDay10Part2(t *testing.T) {
	testWithParserNoErr(t, 10, filename, false, NewDay10, Day10, 1340)
}

func BenchmarkDay10Part1(b *testing.B) {
	benchWithParserNoErr(b, 10, true, NewDay10, Day10)
}

func BenchmarkDay10Part2(b *testing.B) {
	benchWithParserNoErr(b, 10, false, NewDay10, Day10)
}
