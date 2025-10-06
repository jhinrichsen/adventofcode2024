package adventofcode2024

import (
	"fmt"
	"math"
	"testing"
)

// If register C contains 9, the program 2,6 would set register B to 1.
// If register A contains 10, the program 5,0,5,1,5,4 would output 0,1,2.
// If register A contains 2024, the program 0,1,5,4,3,0 would output 4,2,5,6,7,7,7,7,3,1,0 and leave 0 in register A.
// If register B contains 29, the program 1,7 would set register B to 26.
// If register B contains 2024 and register C contains 43690, the program 4,0 would set register B to 44354.
func TestDay17Part1Examples(t *testing.T) {
	const undefined = math.MaxUint
	tests := []struct {
		inA, inB, inC uint
		commands      string
		a, b, c       uint
		output        string
	}{
		{8, 0, 0, "0,2", 2, undefined, undefined, ""},  // adv example 1
		{64, 5, 0, "0,5", 2, undefined, undefined, ""}, // adv example 2
		{0, 1, 2, "4,0", undefined, 3, undefined, ""},  // 1 xor 2 == 3
		{0, 0, 9, "2,6", undefined, 1, undefined, ""},
		{10, 0, 0, "5,0,5,1,5,4", undefined, undefined, undefined, "0,1,2"},
		{2024, 0, 0, "0,1,5,4,3,0", 0, undefined, undefined, "4,2,5,6,7,7,7,7,3,1,0"},
		{0, 29, 0, "1,7", undefined, 26, undefined, ""},
		{0, 2024, 43690, "4,0", undefined, 44354, undefined, ""},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			a, b, c, output := Day17([]string{
				fmt.Sprintf("Register A: %d", tt.inA),
				fmt.Sprintf("Register B: %d", tt.inB),
				fmt.Sprintf("Register C: %d", tt.inC),
				"",
				fmt.Sprintf("Program: %s", tt.commands),
			}, true)
			if tt.a != undefined && tt.a != a || tt.b != undefined && tt.b != b || tt.c != undefined && tt.c != c {
				t.Fatalf("want a=%d, b=%d, c=%d but got a=%d, b=%d, c=%d", tt.a, tt.b, tt.c, a, b, c)
			}
			if tt.output != output {
				t.Fatalf("want output=%q but got %q", tt.output, output)
			}
		})
	}
}

func TestDay17Part1Example(t *testing.T) {
	const want = "4,6,3,5,6,3,5,2,1,0"
	lines := linesFromFilename(t, exampleFilename(17))
	_, _, _, got := Day17(lines, true)
	if want != got {
		t.Fatalf("want %s but got %s", want, got)
	}
}

func TestDay17Part2Example(t *testing.T) {
	const want = ""
	lines := linesFromFilename(t, exampleFilename(17))
	_, _, _, got := Day17(lines, false)
	if want != got {
		t.Fatalf("want %s but got %s", want, got)
	}
}

func TestDay17Part1(t *testing.T) {
	const want = "7,3,5,7,5,7,4,3,0"
	lines := linesFromFilename(t, filename(17))
	_, _, _, got := Day17(lines, true)
	if want != got {
		t.Fatalf("want %s but got %s", want, got)
	}
}

func TestDay17Part2(t *testing.T) {
	const want = ""
	lines := linesFromFilename(t, filename(17))
	_, _, _, got := Day17(lines, false)
	if want != got {
		t.Fatalf("want %s but got %s", want, got)
	}
}

func BenchmarkDay17Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(17))
	for range b.N {
		_, _, _, _ = Day17(lines, true)
	}
}

func BenchmarkDay17Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(17))
	for range b.N {
		_, _, _, _ = Day17(lines, false)
	}
}
