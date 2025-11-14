package adventofcode2024

import "testing"

func TestDay23Part1Example(t *testing.T) {
	const want = 7
	got := Day23(NewDay23(linesFromFilename(t, exampleFilename(23))), true)
	if got != want {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay23Part1(t *testing.T) {
	const want = 1230
	got := Day23(NewDay23(linesFromFilename(t, filename(23))), true)
	if got != want {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay23Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(23))
	for range b.N {
		Day23(NewDay23(lines), true)
	}
}
