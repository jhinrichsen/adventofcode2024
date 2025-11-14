package adventofcode2024

import (
	"fmt"
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

			got := Day13(puzzle, true)
			if got != tt.want {
				t.Errorf("Day13() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay13Part1Example(t *testing.T) {
	testWithParserNoErr(t, 13, exampleFilename, true, NewDay13, Day13, 480)
}

func TestDay13Part1(t *testing.T) {
	testWithParserNoErr(t, 13, filename, true, NewDay13, Day13, 25751)
}

func TestDay13Part2Example(t *testing.T) {
	testWithParserNoErr(t, 13, exampleFilename, false, NewDay13, Day13, 875318608908)
}

func TestDay13Part2(t *testing.T) {
	testWithParserNoErr(t, 13, filename, false, NewDay13, Day13, 108528956728655)
}

func BenchmarkDay13Part1(b *testing.B) {
	benchWithParserNoErr(b, 13, true, NewDay13, Day13)
}

func BenchmarkDay13Part2(b *testing.B) {
	benchWithParserNoErr(b, 13, false, NewDay13, Day13)
}

func BenchmarkCramerVsBareiss(b *testing.B) {
	// Use a typical machine from day13 for benchmarking
	machine := ClawMachine{
		ButtonA: Point{94, 34},
		ButtonB: Point{22, 67},
		Prize:   Point{8400, 5400},
	}

	eq1 := Eq{machine.ButtonA.X, machine.ButtonB.X, machine.Prize.X}
	eq2 := Eq{machine.ButtonA.Y, machine.ButtonB.Y, machine.Prize.Y}
	eqs := []Eq{eq1, eq2}

	b.Run("Cramer", func(b *testing.B) {
		for range b.N {
			Cramer(eq1, eq2)
		}
	})

	b.Run("Bareiss", func(b *testing.B) {
		for range b.N {
			Bareiss(eqs)
		}
	})
}

func TestCramerEquivalentToBareiss(t *testing.T) {
	// Test that Cramer and Bareiss give identical results for day13 machines
	lines := linesFromFilename(t, exampleFilename(13))
	puzzle := NewDay13(lines)

	for i, machine := range puzzle.Machines {
		t.Run(fmt.Sprintf("machine_%d", i+1), func(t *testing.T) {
			// Test both part1 and part2 scenarios
			for _, part1 := range []bool{true, false} {
				// Part 2 default: Add offset
				prizeX := machine.Prize.X + 10000000000000
				prizeY := machine.Prize.Y + 10000000000000

				if part1 {
					prizeX = machine.Prize.X
					prizeY = machine.Prize.Y
				}

				// Cramer solution
				eq1 := Eq{machine.ButtonA.X, machine.ButtonB.X, prizeX}
				eq2 := Eq{machine.ButtonA.Y, machine.ButtonB.Y, prizeY}
				cramX, cramY, cramOk := Cramer(eq1, eq2)

				// Bareiss solution
				eqs := []Eq{eq1, eq2}
				barSol, barOk := Bareiss(eqs)

				// Both should have same success/failure
				if cramOk != barOk {
					t.Errorf("part1=%t: Cramer ok=%t, Bareiss ok=%t, should be equal",
						part1, cramOk, barOk)
					continue
				}

				// If both succeeded, solutions should match
				if cramOk && barOk {
					if len(barSol) != 2 {
						t.Errorf("part1=%t: Bareiss returned %d solutions, want 2",
							part1, len(barSol))
						continue
					}
					if cramX != barSol[0] || cramY != barSol[1] {
						t.Errorf("part1=%t: Cramer=(%d,%d), Bareiss=(%d,%d), should be equal",
							part1, cramX, cramY, barSol[0], barSol[1])
					}
				}
			}
		})
	}
}
