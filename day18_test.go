package adventofcode2024

import (
	"testing"
)

func TestDay18Part1Example(t *testing.T) {
	const want = 22
	lines := linesFromFilename(t, exampleFilename(18))
	p, err := NewDay18(lines, 7, 7) // Grid is 0-6, so 7x7
	if err != nil {
		t.Fatal(err)
	}
	got, _ := Day18(p, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay18Part2Example(t *testing.T) {
	const want = "5,1"
	lines := linesFromFilename(t, exampleFilename(18))
	p, err := NewDay18(lines, 7, 7)
	if err != nil {
		t.Fatal(err)
	}
	_, got := Day18(p, false)
	if want != got {
		t.Fatalf("want %s but got %s", want, got)
	}
}

func TestDay18Part1(t *testing.T) {
	const want = 380
	lines := linesFromFilename(t, filename(18))
	// Only use first 1024 lines
	if len(lines) > 1024 {
		lines = lines[:1024]
	}
	p, err := NewDay18(lines, 71, 71)
	if err != nil {
		t.Fatal(err)
	}
	got, _ := Day18(p, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay18Part2(t *testing.T) {
	const want = "26,50"
	lines := linesFromFilename(t, filename(18))
	p, err := NewDay18(lines, 71, 71)
	if err != nil {
		t.Fatal(err)
	}
	_, got := Day18(p, false)
	if want != got {
		t.Fatalf("want %s but got %s", want, got)
	}
}

func BenchmarkDay18Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(18))
	for range b.N {
		p, _ := NewDay18(lines[:1024], 71, 71)
		_, _ = Day18(p, true)
	}
}

func BenchmarkDay18Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(18))
	for range b.N {
		p, _ := NewDay18(lines, 71, 71)
		_, _ = Day18(p, false)
	}
}
