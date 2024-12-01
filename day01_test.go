package adventofcode2023

import (
	"testing"
)

func TestDay01Part1Example(t *testing.T) {
	const want = 11
	lines, err := linesFromFilename(exampleFilename(01))
	if err != nil {
		t.Fatal(err)
	}
	got := Day01(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay01Part1(t *testing.T) {
	const want = 2166959
	lines, err := linesFromFilename(filename(01))
	if err != nil {
		t.Fatal(err)
	}
	got := Day01(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay01Part1(b *testing.B) {
	lines, err := linesFromFilename(filename(01))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for range b.N {
		_ = Day01(lines)
	}
}

func TestDay01Part2(t *testing.T) {
	const want = 0
	lines, err := linesFromFilename(filename(01))
	if err != nil {
		t.Fatal(err)
	}
	got := Day01(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
