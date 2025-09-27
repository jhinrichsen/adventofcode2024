package adventofcode2024

import (
	"testing"
)

func TestDay00Part1Example(t *testing.T) {
	const want = 0
	lines, err := linesFromFilename(exampleFilename(0))
	if err != nil {
		t.Fatal(err)
	}
	got := Day00(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay00Part2Example(t *testing.T) {
	const want = 0
	lines, err := linesFromFilename(exampleFilename(0))
	if err != nil {
		t.Fatal(err)
	}
	got := Day00(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay00Part1(t *testing.T) {
	const want = 0
	lines, err := linesFromFilename(filename(0))
	if err != nil {
		t.Fatal(err)
	}
	got := Day00(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay00Part2(t *testing.T) {
	const want = 0
	lines, err := linesFromFilename(filename(0))
	if err != nil {
		t.Fatal(err)
	}
	got := Day00(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay00Part1(b *testing.B) {
	lines, err := linesFromFilename(filename(0))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for range b.N {
		_ = Day00(lines)
	}
}

func BenchmarkDay00Part2(b *testing.B) {
	lines, err := linesFromFilename(filename(0))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for range b.N {
		_ = Day00(lines)
	}
}
