package adventofcode2023

import (
	"testing"
)

func TestDay02Part1Examples(t *testing.T) {
	var tt = []struct {
		input string
		want  uint
	}{
		{"7 6 4 2 1", 1},
		{"1 2 7 8 9", 0},
		{"9 7 6 2 1", 0},
		{"1 3 2 4 5", 0},
		{"8 6 4 4 1", 0},
		{"1 3 6 7 9", 1},
	}
	for i := range tt {
		t.Run(tt[i].input, func(t *testing.T) {
			want := tt[i].want
			got := Day02([]string{tt[i].input})
			if want != got {
				t.Fatalf("want %v but got %v", want, got)
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
	got := Day02(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay02Part2Example(t *testing.T) {
	const want = 0
	lines, err := linesFromFilename(exampleFilename(02))
	if err != nil {
		t.Fatal(err)
	}
	got := Day02(lines)
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
	got := Day02(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay02Part2(t *testing.T) {
	const want = 0
	lines, err := linesFromFilename(filename(02))
	if err != nil {
		t.Fatal(err)
	}
	got := Day02(lines)
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
		_ = Day02(lines)
	}
}

func BenchmarkDay02Part2(b *testing.B) {
	lines, err := linesFromFilename(filename(02))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for range b.N {
		_ = Day02(lines)
	}
}
