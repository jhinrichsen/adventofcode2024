package adventofcode2024

import (
	"testing"
)

func TestDay05Part1Example(t *testing.T) {
	const want = 143
	lines := linesFromFilename(t, exampleFilename(05))
	got := Day05(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay05Part2Example(t *testing.T) {
	const want = 123
	lines := linesFromFilename(t, exampleFilename(05))
	got := Day05(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay05Part1(t *testing.T) {
	const want = 5108
	lines := linesFromFilename(t, filename(05))
	got := Day05(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay05Part2(t *testing.T) {
	const want = 7380
	lines := linesFromFilename(t, filename(05))
	got := Day05(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay05Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(05))
	for range b.N {
		_ = Day05(lines, true)
	}
}

func BenchmarkDay05Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(05))
	for range b.N {
		_ = Day05(lines, false)
	}
}
