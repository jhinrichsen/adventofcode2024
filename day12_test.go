package adventofcode2024

import (
	"testing"
)

func TestDay12Part1Example(t *testing.T) {
	tests := []struct {
		name string
		file string
		want uint
	}{
		{
			name: "4x4 example",
			file: example1Filename(12),
			want: 140,
		},
		{
			name: "5x5 with O and X",
			file: example2Filename(12),
			want: 772,
		},
		{
			name: "10x10 larger example",
			file: example3Filename(12),
			want: 1930,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testWithParser(t, 12, func(uint8) string { return tt.file }, true, NewDay12, Day12, tt.want)
		})
	}
}

func TestDay12Part2Example(t *testing.T) {
	tests := []struct {
		name string
		file string
		want uint
	}{
		{
			name: "4x4 example",
			file: example1Filename(12),
			want: 80,
		},
		{
			name: "5x5 with O and X",
			file: example2Filename(12),
			want: 436,
		},
		{
			name: "10x10 larger example",
			file: example3Filename(12),
			want: 1206,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testWithParser(t, 12, func(uint8) string { return tt.file }, false, NewDay12, Day12, tt.want)
		})
	}
}

func TestDay12Part1(t *testing.T) {
	testWithParser(t, 12, filename, true, NewDay12, Day12, 1361494)
}

func TestDay12Part2(t *testing.T) {
	testWithParser(t, 12, filename, false, NewDay12, Day12, 830516)
}

func BenchmarkDay12Part1(b *testing.B) {
	benchWithParser(b, 12, true, NewDay12, Day12)
}

func BenchmarkDay12Part2(b *testing.B) {
	benchWithParser(b, 12, false, NewDay12, Day12)
}
