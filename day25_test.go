package adventofcode2024

import "testing"

func TestDay25Part1Example(t *testing.T) {
	lines := linesFromFilename(t, exampleFilename(25))
	puzzle, err := NewDay25(lines)
	if err != nil {
		t.Fatal(err)
	}
	got := Day25(puzzle)
	want := "3"
	if got != want {
		t.Fatalf("want %q but got %q", want, got)
	}
}

func TestDay25Part1(t *testing.T) {
	lines := linesFromFilename(t, filename(25))
	puzzle, err := NewDay25(lines)
	if err != nil {
		t.Fatal(err)
	}
	got := Day25(puzzle)
	t.Logf("Day 25 Part 1: %s", got)
}

func BenchmarkDay25Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(25))
	for range b.N {
		puzzle, err := NewDay25(lines)
		if err != nil {
			b.Fatal(err)
		}
		Day25(puzzle)
	}
}
