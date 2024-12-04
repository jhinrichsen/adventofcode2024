package adventofcode2024

import (
	"testing"
)

func TestDay04Part1Example(t *testing.T) {
	const want = 18
	grid, err := bytesFromFilename(exampleFilename(4))
	if err != nil {
		t.Fatal(err)
	}
	got := Day04(grid)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay04Part1(t *testing.T) {
	const want = 2685
	grid, err := bytesFromFilename(filename(4))
	if err != nil {
		t.Fatal(err)
	}
	got := Day04(grid)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay04Part1(b *testing.B) {
	for range b.N {
		grid, _ := bytesFromFilename(filename(4))
		_ = Day04(grid)
	}
}
