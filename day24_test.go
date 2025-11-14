package adventofcode2024

import "testing"

func TestDay24Part1Example1(t *testing.T) {
	const want = "4"
	got := Day24(NewDay24(linesFromFilename(t, example1Filename(24))), true)
	if got != want {
		t.Fatalf("want %q but got %q", want, got)
	}
}

func TestDay24Part1Example(t *testing.T) {
	const want = "2024"
	got := Day24(NewDay24(linesFromFilename(t, exampleFilename(24))), true)
	if got != want {
		t.Fatalf("want %q but got %q", want, got)
	}
}

func TestDay24Part1(t *testing.T) {
	const want = "59336987801432"
	got := Day24(NewDay24(linesFromFilename(t, filename(24))), true)
	if got != want {
		t.Fatalf("want %q but got %q", want, got)
	}
}

func BenchmarkDay24Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(24))
	for range b.N {
		Day24(NewDay24(lines), true)
	}
}
