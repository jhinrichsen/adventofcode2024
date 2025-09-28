package adventofcode2024

import (
	"testing"
)

func TestDay09Part1Example(t *testing.T) {
	const (
		buf  = "2333133121414131402"
		want = 1928
	)
	got := Day09([]byte(buf), true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay09Part2Example(t *testing.T) {
	const (
		buf  = "2333133121414131402"
		want = 2858
	)
	got := Day09([]byte(buf), false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay09Part1(t *testing.T) {
	const want = 6337921897505
	buf := file(t, 9)
	got := Day09(buf, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay09Part2(t *testing.T) {
	const want = 6362722604045
	buf := file(t, 9)
	got := Day09(buf, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay09Part1(b *testing.B) {
	buf := file(b, 9)
	for range b.N {
		_ = Day09(buf, true)
	}
}

func BenchmarkDay09Part2(b *testing.B) {
	buf := file(b, 9)
	for range b.N {
		_ = Day09(buf, false)
	}
}
