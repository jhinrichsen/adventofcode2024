package adventofcode2024

import (
	"testing"
)

func TestDay14Part1Example1(t *testing.T) {
	const want = 0
	p, err := NewDay14([]string{"p=2,4 v=2,-3"}, 11, 7)
	if err != nil {
		t.Fatal(err)
	}
	got := Day14(p, 5, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay14Part1Example(t *testing.T) {
	const want = 12
	p, err := NewDay14(linesFromFilename(t, exampleFilename(14)), 11, 7)
	if err != nil {
		t.Fatal(err)
	}
	got := Day14(p, 100, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay14Part1(t *testing.T) {
	const want = 230461440
	p, err := NewDay14(linesFromFilename(t, filename(14)), 101, 103)
	if err != nil {
		t.Fatal(err)
	}
	got := Day14(p, 100, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay14Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(14))
	for range b.N {
		p, _ := NewDay14(lines, 101, 103)
		_ = Day14(p, 100, true)
	}
}
