package adventofcode2024

import (
	"testing"
)

func TestDay12Part1Example(t *testing.T) {
	tests := []struct {
		name string
		file string
		want uint
	}{
		{
			name: "4x4 example",
			file: example1Filename(12),
			want: 140,
		},
		{
			name: "5x5 with O and X",
			file: example2Filename(12),
			want: 772,
		},
		{
			name: "10x10 larger example",
			file: example3Filename(12),
			want: 1930,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines := linesFromFilename(t, tt.file)
			puzzle := NewDay12(lines)
			got := Day12(puzzle)
			if got != tt.want {
				t.Errorf("Day12() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay12Part1(t *testing.T) {
	const want = 1361494
	got := Day12(NewDay12(linesFromFilename(t, filename(12))))
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
