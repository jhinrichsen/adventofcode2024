package adventofcode2024

import (
	"testing"
)

func TestDay21Part1Example(t *testing.T) {
	const want = 126384
	got := Day21(NewDay21(linesFromFilename(t, exampleFilename(21))), true)
	if got != want {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay21Part1Example029A(t *testing.T) {
	// From website: 029A produces 68-character sequence with complexity 1972
	puzzle := NewDay21([]string{"029A"})
	gotSeqLength := uint(len(puzzle.findShortestSequence("029A")))
	gotComplexity := gotSeqLength * puzzle.extractNumericValue("029A")

	wantSeqLength := uint(68)
	wantComplexity := uint(1972) // 68 * 29 = 1972

	if gotSeqLength != wantSeqLength {
		t.Errorf("sequence length: want %d, got %d", wantSeqLength, gotSeqLength)
	}
	if gotComplexity != wantComplexity {
		t.Errorf("complexity: want %d, got %d", wantComplexity, gotComplexity)
	}
}

func TestDay21Part1ExampleOthers(t *testing.T) {
	// From website: total is 126384, and 029A is 1972
	// So the other 4 codes must sum to: 126384 - 1972 = 124412
	puzzle := NewDay21([]string{"980A", "179A", "456A", "379A"})

	var totalComplexity uint
	for _, code := range puzzle.codes {
		totalComplexity += uint(len(puzzle.findShortestSequence(code))) * puzzle.extractNumericValue(code)
	}

	wantTotal := uint(124412) // 126384 - 1972 = 124412
	if totalComplexity != wantTotal {
		t.Errorf("sum of other 4 complexities: want %d, got %d", wantTotal, totalComplexity)
	}
}

func TestDay21Part1(t *testing.T) {
	const want = 161468
	got := Day21(NewDay21(linesFromFilename(t, filename(21))), true)
	if got != want {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay21Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(21))
	for range b.N {
		Day21(NewDay21(lines), true)
	}
}
