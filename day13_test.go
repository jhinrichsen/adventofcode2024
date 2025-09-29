package adventofcode2024

import (
	"testing"
)

func TestDay13Part1Examples(t *testing.T) {
	tests := []struct {
		name string
		file string
		want int
	}{
		{
			name: "Machine 1 - Solvable (80A + 40B = 280 tokens)",
			file: example1Filename(13),
			want: 280,
		},
		{
			name: "Machine 2 - Unsolvable",
			file: example2Filename(13),
			want: 0,
		},
		{
			name: "Machine 3 - Solvable (38A + 86B = 200 tokens)",
			file: example3Filename(13),
			want: 200,
		},
		{
			name: "Machine 4 - Unsolvable",
			file: example4Filename(13),
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines := linesFromFilename(t, tt.file)
			puzzle := NewDay13(lines)

			if len(puzzle.Machines) != 1 {
				t.Fatalf("Expected 1 machine, got %d", len(puzzle.Machines))
			}

			got := Day13(puzzle)
			if got != tt.want {
				t.Errorf("Day13() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay13Part1Example(t *testing.T) {
	const want = 280 + 0 + 200 + 0
	lines := linesFromFilename(t, exampleFilename(13))
	puzzle := NewDay13(lines)
	got := Day13(puzzle)
	if got != want {
		t.Errorf("Day13() = %v, want %v", got, want)
	}
}

func TestDay13Part1(t *testing.T) {
	const want = 25751
	puzzle := NewDay13(linesFromFilename(t, filename(13)))
	got := Day13(puzzle)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay13Part1(b *testing.B) {
	puzzle := NewDay13(linesFromFilename(b, filename(13)))
	for range b.N {
		_ = Day13(puzzle)
	}
}
