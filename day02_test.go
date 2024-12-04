package adventofcode2024

import (
	"fmt"
	"testing"
)

func TestDay02Examples(t *testing.T) {
	var tt = []struct {
		input []uint
		want  uint
		part1 bool
	}{
		{[]uint{7, 6, 4, 2, 1}, 1, true},
		{[]uint{1, 2, 7, 8, 9}, 0, true},
		{[]uint{9, 7, 6, 2, 1}, 0, true},
		{[]uint{1, 3, 2, 4, 5}, 0, true},
		{[]uint{8, 6, 4, 4, 1}, 0, true},
		{[]uint{1, 3, 6, 7, 9}, 1, true},

		{[]uint{7, 6, 4, 2, 1}, 1, false},
		{[]uint{1, 2, 7, 8, 9}, 0, false},
		{[]uint{9, 7, 6, 2, 1}, 0, false},
		{[]uint{1, 3, 2, 4, 5}, 1, false},
		{[]uint{8, 6, 4, 4, 1}, 1, false},
		{[]uint{1, 3, 6, 7, 9}, 1, false},
	}
	for i := range tt {
		t.Run(fmt.Sprintf("%v", tt[i].input), func(t *testing.T) {
			want := tt[i].want
			got := Day02([][]uint{tt[i].input}, tt[i].part1)
			if want != got {
				t.Fatalf("part1: %t, want %v but got %v", tt[i].part1, want, got)
			}
		})
	}
}

func TestDay02Part1Example(t *testing.T) {
	const want = 2
	got := Day02(mustNumbersFromFilename(exampleFilename(02)), true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay02Part2Example(t *testing.T) {
	const want = 4
	got := Day02(mustNumbersFromFilename(exampleFilename(02)), false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay02Part1(t *testing.T) {
	const want = 526
	got := Day02(mustNumbersFromFilename(filename(02)), true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay02Part2(t *testing.T) {
	const want = 566
	got := Day02(mustNumbersFromFilename(filename(02)), false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay02Part1(b *testing.B) {
	for range b.N {
		_ = Day02(mustNumbersFromFilename(filename(02)), true)
	}
}

func BenchmarkDay02Part2(b *testing.B) {
	for range b.N {
		_ = Day02(mustNumbersFromFilename(filename(02)), false)
	}
}
