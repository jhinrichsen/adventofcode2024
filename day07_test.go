package adventofcode2024

import (
	"testing"
)

func TestDay07Part1Example(t *testing.T) {
	const want = 3749
	lines := linesFromFilename(t, exampleFilename(07))
	puzzle := NewDay07(lines)
	got := Day07(puzzle, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay07Part2Example(t *testing.T) {
	const want = 11387
	lines := linesFromFilename(t, exampleFilename(07))
	puzzle := NewDay07(lines)
	got := Day07(puzzle, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay07Part1(t *testing.T) {
	const want = 20281182715321
	lines := linesFromFilename(t, filename(07))
	puzzle := NewDay07(lines)
	got := Day07(puzzle, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay07Part2(t *testing.T) {
	const want = 159490400628354
	lines := linesFromFilename(t, filename(07))
	puzzle := NewDay07(lines)
	got := Day07(puzzle, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay07Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(07))
	for range b.N {
		puzzle := NewDay07(lines)
		_ = Day07(puzzle, true)
	}
}

func BenchmarkDay07Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(07))
	for range b.N {
		puzzle := NewDay07(lines)
		_ = Day07(puzzle, false)
	}
}
