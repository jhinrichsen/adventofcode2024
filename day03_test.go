package adventofcode2024

import (
	"os"
	"testing"
)

func TestDay03Part1Example(t *testing.T) {
	const (
		line = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
		want = 161
	)
	got := Day03(line, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay03Part2Example(t *testing.T) {
	const (
		line = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5)"
		want = 48
	)
	got := Day03(line, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay03Part1(t *testing.T) {
	const want = 184576302
	program, err := os.ReadFile(filename(03))
	if err != nil {
		t.Fatal(err)
	}
	got := Day03(string(program), true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay03Part2(t *testing.T) {
	// const want = 31862452 // too low
	// const want = 184576302 // too high
	const want = 118173507 // too high
	program, err := os.ReadFile(filename(03))
	if err != nil {
		t.Fatal(err)
	}
	got := Day03(string(program), false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay03Part1(b *testing.B) {
	program, err := os.ReadFile(filename(03))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for range b.N {
		_ = Day03(string(program), true)
	}
}

func BenchmarkDay03Part2(b *testing.B) {
	program, err := os.ReadFile(filename(03))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for range b.N {
		_ = Day03(string(program), false)
	}
}
