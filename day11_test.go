package adventofcode2024

import (
	"os"
	"testing"
)

func TestDay11Part1Example(t *testing.T) {
	const want = 55312
	got := Day11([]uint64{125, 17}, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay11Part1(t *testing.T) {
	const want = 175006
	data, err := os.ReadFile("testdata/day11.txt")
	if err != nil {
		t.Fatal(err)
	}
	stones := NewDay11(string(data))
	got := Day11(stones, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay11Part2Example(t *testing.T) {
	const want = 65601038650482
	got := Day11([]uint64{125, 17}, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay11Part2(t *testing.T) {
	const want = 207961583799296
	data, err := os.ReadFile("testdata/day11.txt")
	if err != nil {
		t.Fatal(err)
	}
	stones := NewDay11(string(data))
	got := Day11(stones, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay11Part1(b *testing.B) {
	data, err := os.ReadFile("testdata/day11.txt")
	if err != nil {
		b.Fatal(err)
	}
	for range b.N {
		stones := NewDay11(string(data))
		_ = Day11(stones, true)
	}
}

func BenchmarkDay11Part2(b *testing.B) {
	data, err := os.ReadFile("testdata/day11.txt")
	if err != nil {
		b.Fatal(err)
	}
	for range b.N {
		stones := NewDay11(string(data))
		_ = Day11(stones, false)
	}
}
