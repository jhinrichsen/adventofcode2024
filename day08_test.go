package adventofcode2024

import (
	"testing"
)

func TestDay08Part1Example(t *testing.T) {
	const want = 14
	lines, err := linesFromFilename(exampleFilename(8))
	if err != nil {
		t.Fatal(err)
	}
	got := Day08(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay08Part2Example(t *testing.T) {
	const want = 0
	lines, err := linesFromFilename(exampleFilename(8))
	if err != nil {
		t.Fatal(err)
	}
	got := Day08(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay08Part1(t *testing.T) {
	const want = 291
	lines, err := linesFromFilename(filename(8))
	if err != nil {
		t.Fatal(err)
	}
	got := Day08(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay08Part2(t *testing.T) {
	const want = 0
	lines, err := linesFromFilename(filename(8))
	if err != nil {
		t.Fatal(err)
	}
	got := Day08(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay08Part1(b *testing.B) {
	lines, err := linesFromFilename(filename(8))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for range b.N {
		_ = Day08(lines)
	}
}

func BenchmarkDay08Part2(b *testing.B) {
	lines, err := linesFromFilename(filename(8))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for range b.N {
		_ = Day08(lines)
	}
}
