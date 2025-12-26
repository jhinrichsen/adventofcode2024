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
			lines := linesFromFilename(t, tt.file)
			puzzle := NewDay10(lines)
			got := Day10(puzzle, true)
			if tt.want != got {
				t.Fatalf("want %d but got %d", tt.want, got)
			}
		})
	}
}

func TestDay10Part1Example(t *testing.T) {
	const want = 36
	lines := linesFromFilename(t, exampleFilename(10))
	puzzle := NewDay10(lines)
	got := Day10(puzzle, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay10Part2Example(t *testing.T) {
	const want = 81
	lines := linesFromFilename(t, exampleFilename(10))
	puzzle := NewDay10(lines)
	got := Day10(puzzle, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay10Part1(t *testing.T) {
	const want = 587
	lines := linesFromFilename(t, filename(10))
	puzzle := NewDay10(lines)
	got := Day10(puzzle, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay10Part2(t *testing.T) {
	const want = 1340
	lines := linesFromFilename(t, filename(10))
	puzzle := NewDay10(lines)
	got := Day10(puzzle, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay10Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(10))
	puzzle := NewDay10(lines)
	b.ResetTimer()
	for range b.N {
		_ = Day10(puzzle, true)
	}
}

func BenchmarkDay10Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(10))
	puzzle := NewDay10(lines)
	b.ResetTimer()
	for range b.N {
		_ = Day10(puzzle, false)
	}
}
