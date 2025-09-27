package adventofcode2024

import (
	"testing"
)

func TestDay08Part1Example(t *testing.T) {
	const want = 14
	lines := linesFromFilename(t, exampleFilename(8))
	got := Day08(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay08Part2Example(t *testing.T) {
	const want = 34
	lines := linesFromFilename(t, exampleFilename(8))
	got := Day08(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay08Part1(t *testing.T) {
	const want = 291
	lines := linesFromFilename(t, filename(8))
	got := Day08(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay08Part2(t *testing.T) {
	const want = 1015
	lines := linesFromFilename(t, filename(8))
	got := Day08(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay08Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(8))
	for range b.N {
		_ = Day08(lines, true)
	}
}

func BenchmarkDay08Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(8))
	for range b.N {
		_ = Day08(lines, false)
	}
}
