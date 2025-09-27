package adventofcode2024

import (
	"testing"
)

func TestDay06Part1Example(t *testing.T) {
	const want = 41
	lines := linesFromFilename(t, exampleFilename(6))
	puzzle := NewDay06(lines)
	got := Day06(puzzle)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay06Part1(t *testing.T) {
	const want = 4903
	lines := linesFromFilename(t, filename(6))
	puzzle := NewDay06(lines)
	got := Day06(puzzle)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay06Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(6))
	for range b.N {
		puzzle := NewDay06(lines)
		_ = Day06(puzzle)
	}
}
