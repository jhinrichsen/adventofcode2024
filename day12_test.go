package adventofcode2024

import (
	"os"
	"testing"
)

func TestDay12Part1Example(t *testing.T) {
	tests := []struct {
		name string
		file string
		want uint
	}{
		{"4x4 example", example1Filename(12), 140},
		{"5x5 with O and X", example2Filename(12), 772},
		{"10x10 larger example", example3Filename(12), 1930},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf, err := os.ReadFile(tt.file)
			if err != nil {
				t.Fatal(err)
			}
			got := Day12(buf, true)
			if got != tt.want {
				t.Errorf("Day12() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay12Part2Example(t *testing.T) {
	tests := []struct {
		name string
		file string
		want uint
	}{
		{"4x4 example", example1Filename(12), 80},
		{"5x5 with O and X", example2Filename(12), 436},
		{"10x10 larger example", example3Filename(12), 1206},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf, err := os.ReadFile(tt.file)
			if err != nil {
				t.Fatal(err)
			}
			got := Day12(buf, false)
			if got != tt.want {
				t.Errorf("Day12() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay12Part1(t *testing.T) {
	const want = 1361494
	got := Day12(file(t, 12), true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay12Part2(t *testing.T) {
	const want = 830516
	got := Day12(file(t, 12), false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay12Part1(b *testing.B) {
	buf := file(b, 12)
	b.ResetTimer()
	for range b.N {
		_ = Day12(buf, true)
	}
}

func BenchmarkDay12Part2(b *testing.B) {
	buf := file(b, 12)
	b.ResetTimer()
	for range b.N {
		_ = Day12(buf, false)
	}
}
