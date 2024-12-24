package adventofcode2024

import (
	"strconv"
	"testing"
)

func TestDay10Part1Examples(t *testing.T) {
	var tt = []struct {
		want   uint
		sample string
	}{
		{2, `...0...
...1...
...2...
6543456
7.....7
8.....8
9.....9
`},
		{4, `..90..9
...1.98
...2..7
6543456
765.987
876....
987....
`},
		{3, `10..9..
2...8..
3...7..
4567654
...8..3
...9..2
.....01
`},
	}

	for i := range tt {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			want := tt[i].want
			got := Day10(gridFromBytes([]byte(tt[i].sample)), true)
			if want != got {
				t.Fatalf("want %d but got %d\n", want, got)
			}
		})
	}
}

func TestDay10Part1Example(t *testing.T) {
	const want = 36
	lines, err := gridFromFilename(exampleFilename(10))
	if err != nil {
		t.Fatal(err)
	}
	got := Day10(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay10Part2Example(t *testing.T) {
	const want = 81
	lines, err := gridFromFilename(exampleFilename(10))
	if err != nil {
		t.Fatal(err)
	}
	got := Day10(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay10Part1(t *testing.T) {
	const want = 587
	lines, err := gridFromFilename(filename(10))
	if err != nil {
		t.Fatal(err)
	}
	got := Day10(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay10Part2(t *testing.T) {
	const want = 1340
	lines, err := gridFromFilename(filename(10))
	if err != nil {
		t.Fatal(err)
	}
	got := Day10(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay10Part1(b *testing.B) {
	lines, err := gridFromFilename(filename(10))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for range b.N {
		_ = Day10(lines, true)
	}
}

func BenchmarkDay10Part2(b *testing.B) {
	lines, err := gridFromFilename(filename(10))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for range b.N {
		_ = Day10(lines, false)
	}
}
