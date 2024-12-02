package adventofcode2023

import (
	"testing"
)

func TestDay02Examples(t *testing.T) {
	var tt = []struct {
		input string
		want  uint
		part1 bool
	}{
		{"7 6 4 2 1", 1, true},
		{"1 2 7 8 9", 0, true},
		{"9 7 6 2 1", 0, true},
		{"1 3 2 4 5", 0, true},
		{"8 6 4 4 1", 0, true},
		{"1 3 6 7 9", 1, true},

		{"7 6 4 2 1", 1, false},
		{"1 2 7 8 9", 0, false},
		{"9 7 6 2 1", 0, false},
		{"1 3 2 4 5", 1, false},
		{"8 6 4 4 1", 1, false},
		{"1 3 6 7 9", 1, false},
	}
	for i := range tt {
		t.Run(tt[i].input, func(t *testing.T) {
			want := tt[i].want
			got := Day02([]string{tt[i].input}, tt[i].part1)
			if want != got {
				t.Fatalf("part1: %t, want %v but got %v", tt[i].part1, want, got)
			}
		})
	}
}

func TestDay02Part1Example(t *testing.T) {
	const want = 2
	lines, err := linesFromFilename(exampleFilename(02))
	if err != nil {
		t.Fatal(err)
	}
	got := Day02(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay02Part2Example(t *testing.T) {
	const want = 4
	lines, err := linesFromFilename(exampleFilename(02))
	if err != nil {
		t.Fatal(err)
	}
	got := Day02(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay02Part1(t *testing.T) {
	const want = 526
	lines, err := linesFromFilename(filename(02))
	if err != nil {
		t.Fatal(err)
	}
	got := Day02(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay02Part2(t *testing.T) {
	const want = 566
	lines, err := linesFromFilename(filename(02))
	if err != nil {
		t.Fatal(err)
	}
	got := Day02(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay02Part1(b *testing.B) {
	lines, err := linesFromFilename(filename(02))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for range b.N {
		_ = Day02(lines, true)
	}
}

func BenchmarkDay02Part2(b *testing.B) {
	lines, err := linesFromFilename(filename(02))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for range b.N {
		_ = Day02(lines, false)
	}
}
