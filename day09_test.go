package adventofcode2024

import (
	"os"
	"testing"
)

func TestDay09Part1Example(t *testing.T) {
	const (
		buf  = "2333133121414131402"
		want = 1928
	)
	got := Day09([]byte(buf))
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay09Part1(t *testing.T) {
	const want = 6337921897505
	buf, err := os.ReadFile(filename(9))
	if err != nil {
		t.Fatal(err)
	}
	got := Day09(buf)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay09Part1(b *testing.B) {
	buf, err := os.ReadFile(filename(9))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for range b.N {
		_ = Day09(buf)
	}
}
