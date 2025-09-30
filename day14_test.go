package adventofcode2024

import (
	"os"
	"testing"
)

func TestDay14Part1Example1(t *testing.T) {
	const want = 0
	buf := []byte("p=2,4 v=2,-3")
	got := Day14(buf, 11, 7, 5, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay14Part1Example(t *testing.T) {
	const want = 12
	buf, err := os.ReadFile(exampleFilename(14))
	if err != nil {
		t.Fatal(err)
	}
	got := Day14(buf, 11, 7, 100, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay14Part1(t *testing.T) {
	const want = 230461440
	buf, err := os.ReadFile(filename(14))
	if err != nil {
		t.Fatal(err)
	}
	got := Day14(buf, 101, 103, 100, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay14Part2(t *testing.T) {
	buf, err := os.ReadFile(filename(14))
	if err != nil {
		t.Fatal(err)
	}
	got := Day14(buf, 101, 103, 0, false) // seconds ignored for part2
	t.Logf("Day 14 Part 2 result: %d", got)

	// Sanity check - should be a reasonable number of seconds
	if got == 0 || got > 20000 {
		t.Errorf("Part 2 result %d seems unreasonable", got)
	}
}

func BenchmarkDay14Part1(b *testing.B) {
	buf, err := os.ReadFile(filename(14))
	if err != nil {
		b.Fatal(err)
	}
	for range b.N {
		_ = Day14(buf, 101, 103, 100, true)
	}
}

func BenchmarkDay14Part2(b *testing.B) {
	buf, err := os.ReadFile(filename(14))
	if err != nil {
		b.Fatal(err)
	}
	for range b.N {
		_ = Day14(buf, 101, 103, 0, false) // seconds ignored for part2
	}
}
