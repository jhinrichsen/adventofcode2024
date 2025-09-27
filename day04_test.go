package adventofcode2024

import (
	"testing"
)

func TestDay04Part1Example(t *testing.T) {
	const want = 18
	lines := linesFromFilename(t, exampleFilename(4))
	puzzle := NewDay04(lines)
	got := Day04(puzzle, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay04Part1(t *testing.T) {
	const want = 2685
	lines := linesFromFilename(t, filename(4))
	puzzle := NewDay04(lines)
	got := Day04(puzzle, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay04Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(4))
	for range b.N {
		puzzle := NewDay04(lines)
		_ = Day04(puzzle, true)
	}
}

func TestDay04Part2Example(t *testing.T) {
	const want = 9
	lines := linesFromFilename(t, exampleFilename(4))
	puzzle := NewDay04(lines)
	got := Day04(puzzle, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay04Part2(t *testing.T) {
	const want = 2048
	lines := linesFromFilename(t, filename(4))
	puzzle := NewDay04(lines)
	got := Day04(puzzle, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay04Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(4))
	for range b.N {
		puzzle := NewDay04(lines)
		_ = Day04(puzzle, false)
	}
}
