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
	got := Day01(lines, true)
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
	got := Day01(lines, true)
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
		_ = Day01(lines, true)
	}
}

func TestDay01Part2Example(t *testing.T) {
	const want = 31
	lines, err := linesFromFilename(exampleFilename(01))
	if err != nil {
		t.Fatal(err)
	}
	got := Day01(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay01Part2(t *testing.T) {
	const want = 23741109
	lines, err := linesFromFilename(filename(01))
	if err != nil {
		t.Fatal(err)
	}
	got := Day01(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay01Part2(b *testing.B) {
	lines, err := linesFromFilename(filename(01))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for range b.N {
		_ = Day01(lines, false)
	}
}
