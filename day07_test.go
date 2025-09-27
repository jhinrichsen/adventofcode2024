package adventofcode2024

import (
	"testing"
)

func TestDay07Part1Example(t *testing.T) {
	const want = 3749
	lines := linesFromFilename(t, exampleFilename(07))
	got := Day07(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay07Part1(t *testing.T) {
	const want = 20281182715321
	lines := linesFromFilename(t, filename(07))
	got := Day07(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay07Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(07))
	for range b.N {
		_ = Day07(lines)
	}
}
