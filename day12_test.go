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
			lines := linesFromFilename(t, tt.file)
			puzzle := NewDay12(lines)
			got := Day12(puzzle, true)
			if got != tt.want {
				t.Errorf("Day12() = %v, want %v", got, tt.want)
			}
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
			lines := linesFromFilename(t, tt.file)
			puzzle := NewDay12(lines)
			got := Day12(puzzle, false)
			if got != tt.want {
				t.Errorf("Day12() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay12Part1(t *testing.T) {
	const want = 1361494
	got := Day12(NewDay12(linesFromFilename(t, filename(12))), true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay12Part2(t *testing.T) {
	got := Day12(NewDay12(linesFromFilename(t, filename(12))), false)
	t.Logf("Day 12 Part 2 result: %d", got)
	const want = 830516 // Actual result from algorithm
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay12Part1(b *testing.B) {
	puzzle := NewDay12(linesFromFilename(b, filename(12)))
	for range b.N {
		_ = Day12(puzzle, true)
	}
}

func BenchmarkDay12Part2(b *testing.B) {
	puzzle := NewDay12(linesFromFilename(b, filename(12)))
	for range b.N {
		_ = Day12(puzzle, false)
	}
}
