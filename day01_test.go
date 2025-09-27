package adventofcode2024

import (
	"slices"
	"testing"
)

func TestDay01Part1Example(t *testing.T) {
	const want = 11
	in := exampleFile(t, 1)
	got := Day01(in, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay01Part1(t *testing.T) {
	const want = 2166959
	in := file(t, 1)
	got := Day01(in, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay01Part1(b *testing.B) {
	in := file(b, 1)
	for range b.N {
		_ = Day01(in, true)
	}
}

func TestDay01Part2Example(t *testing.T) {
	const want = 31
	in := exampleFile(t, 1)
	got := Day01(in, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay01Part2(t *testing.T) {
	const want = 23741109
	in := file(t, 1)
	got := Day01(in, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay01Part2(b *testing.B) {
	in := file(b, 1)
	for range b.N {
		_ = Day01(in, false)
	}
}

func TestTwoUint(t *testing.T) {
	buf := []byte("1234   4321\n6789   9876\n")
	left := []uint{1234, 6789}
	right := []uint{4321, 9876}
	got := twoUints(buf)
	if !slices.Equal(left, got[0]) {
		t.Fatalf("want left=%+v but got %+v", left, got[0])
	}
	if !slices.Equal(right, got[1]) {
		t.Fatalf("want right=%+v but got %+v", right, got[1])
	}
}

func BenchmarkTwoUint(b *testing.B) {
	buf := file(b, 1)
	for range b.N {
		_ = twoUints(buf)
	}
}
