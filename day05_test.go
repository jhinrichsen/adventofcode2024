package adventofcode2024

import (
	"testing"
)

func TestDay05Part1Example(t *testing.T) {
	const want = 143
	lines, err := linesFromFilename(exampleFilename(05))
	if err != nil {
		t.Fatal(err)
	}
	got := Day05(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay05Part2Example(t *testing.T) {
	const want = 123
	lines, err := linesFromFilename(exampleFilename(05))
	if err != nil {
		t.Fatal(err)
	}
	got := Day05(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay05Part1(t *testing.T) {
	const want = 5108
	lines, err := linesFromFilename(filename(05))
	if err != nil {
		t.Fatal(err)
	}
	got := Day05(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay05Part2(t *testing.T) {
	const want = 7380
	lines, err := linesFromFilename(filename(05))
	if err != nil {
		t.Fatal(err)
	}
	got := Day05(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay05Part1(b *testing.B) {
	lines, err := linesFromFilename(filename(05))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for range b.N {
		_ = Day05(lines, true)
	}
}

func BenchmarkDay05Part2(b *testing.B) {
	lines, err := linesFromFilename(filename(05))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for range b.N {
		_ = Day05(lines, false)
	}
}
