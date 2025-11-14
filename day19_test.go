package adventofcode2024

import (
	"testing"
)

func TestDay19Part1Example(t *testing.T) {
	const want = 6
	lines := linesFromFilename(t, exampleFilename(19))
	p, err := NewDay19(lines)
	if err != nil {
		t.Fatal(err)
	}
	got, _ := Day19(p, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay19Part2Example(t *testing.T) {
	const want = 16 // 2 + 1 + 4 + 6 + 1 + 2 = 16 total ways
	lines := linesFromFilename(t, exampleFilename(19))
	p, err := NewDay19(lines)
	if err != nil {
		t.Fatal(err)
	}
	got, _ := Day19(p, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay19Part1(t *testing.T) {
	const want = 260
	lines := linesFromFilename(t, filename(19))
	p, err := NewDay19(lines)
	if err != nil {
		t.Fatal(err)
	}
	got, _ := Day19(p, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay19Part2(t *testing.T) {
	const want = 639963796864990
	lines := linesFromFilename(t, filename(19))
	p, err := NewDay19(lines)
	if err != nil {
		t.Fatal(err)
	}
	got, _ := Day19(p, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay19Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(19))
	b.ResetTimer()
	for range b.N {
		p, err := NewDay19(lines)
		if err != nil {
			b.Fatal(err)
		}
		_, _ = Day19(p, true)
	}
}

func BenchmarkDay19Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(19))
	b.ResetTimer()
	for range b.N {
		p, err := NewDay19(lines)
		if err != nil {
			b.Fatal(err)
		}
		_, _ = Day19(p, false)
	}
}