package adventofcode2024

import (
	"testing"
)

func TestDay06Part1Example(t *testing.T) {
	const want = 41
	grid, err := bytesFromFilename(exampleFilename(06))
	if err != nil {
		t.Fatal(err)
	}
	got := Day06(grid)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay06Part1(t *testing.T) {
	const want = 4903
	grid, err := bytesFromFilename(filename(06))
	if err != nil {
		t.Fatal(err)
	}
	got := Day06(grid)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay06Part1(b *testing.B) {
	grid, err := bytesFromFilename(filename(06))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for range b.N {
		_ = Day06(grid)
	}
}
