package adventofcode2024

import (
	"testing"
)

func TestDay07Part1Example(t *testing.T) {
	const want = 3749
	lines, err := linesFromFilename(exampleFilename(07))
	if err != nil {
		t.Fatal(err)
	}
	got := Day07(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay07Part1(t *testing.T) {
	const want = 20281182715321
	lines, err := linesFromFilename(filename(07))
	if err != nil {
		t.Fatal(err)
	}
	got := Day07(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay07Part1(b *testing.B) {
	lines, err := linesFromFilename(filename(07))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for range b.N {
		_ = Day07(lines)
	}
}
