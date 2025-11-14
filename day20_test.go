package adventofcode2024

import (
	"testing"
)

func TestDay20Part1Example(t *testing.T) {
	// Test with lower threshold since example has smaller savings
	lines := linesFromFilename(t, exampleFilename(20))
	puzzle, err := NewDay20(lines)
	if err != nil {
		t.Fatal(err)
	}

	// Test normal path exists
	normalTime := puzzle.findShortestPath(puzzle.start, puzzle.end)
	if normalTime != 84 {
		t.Fatalf("Expected normal time 84, got %d", normalTime)
	}

	// Test cheats with different thresholds
	tests := []struct {
		minSaving uint
		want      uint
	}{
		{2, 44},  // All cheats saving 2+ picoseconds
		{4, 30},  // All cheats saving 4+ picoseconds
		{6, 16},  // All cheats saving 6+ picoseconds
		{8, 14},  // All cheats saving 8+ picoseconds
		{10, 10}, // All cheats saving 10+ picoseconds
		{12, 8},  // All cheats saving 12+ picoseconds
		{20, 5},  // All cheats saving 20+ picoseconds
		{36, 4},  // All cheats saving 36+ picoseconds
		{38, 3},  // All cheats saving 38+ picoseconds
		{40, 2},  // All cheats saving 40+ picoseconds
		{64, 1},  // All cheats saving 64+ picoseconds
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := puzzle.countCheats(tt.minSaving, 2)
			if got != tt.want {
				t.Errorf("countCheats(minSaving=%d, maxCheatDist=2) = %d, want %d", tt.minSaving, got, tt.want)
			}
		})
	}
}

func TestDay20Part2Example(t *testing.T) {
	lines := linesFromFilename(t, exampleFilename(20))
	puzzle, err := NewDay20(lines)
	if err != nil {
		t.Fatal(err)
	}

	// Debug: Print exact savings counts
	savings := puzzle.countCheatsBySavings(20)
	t.Logf("Exact savings counts for Part 2:")
	for saving := uint(50); saving <= 76; saving += 2 {
		if count, exists := savings[saving]; exists && count > 0 {
			t.Logf("  %d picoseconds: %d cheats", saving, count)
		}
	}

	// Test Part 2 with exact savings based on problem description
	tests := []struct {
		exactSaving uint
		want        uint
	}{
		{50, 32},  // 32 cheats that save 50 picoseconds
		{52, 31},  // 31 cheats that save 52 picoseconds
		{54, 29},  // 29 cheats that save 54 picoseconds
		{56, 39},  // 39 cheats that save 56 picoseconds
		{58, 25},  // 25 cheats that save 58 picoseconds
		{60, 23},  // 23 cheats that save 60 picoseconds
		{62, 20},  // 20 cheats that save 62 picoseconds
		{64, 19},  // 19 cheats that save 64 picoseconds
		{66, 12},  // 12 cheats that save 66 picoseconds
		{68, 14},  // 14 cheats that save 68 picoseconds
		{70, 12},  // 12 cheats that save 70 picoseconds
		{72, 22},  // 22 cheats that save 72 picoseconds
		{74, 4},   // 4 cheats that save 74 picoseconds
		{76, 3},   // 3 cheats that save 76 picoseconds
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := savings[tt.exactSaving]
			if got != tt.want {
				t.Errorf("cheats saving exactly %d picoseconds = %d, want %d", tt.exactSaving, got, tt.want)
			}
		})
	}
}

func TestDay20Part1(t *testing.T) {
	const want = 1497
	lines := linesFromFilename(t, filename(20))
	puzzle, err := NewDay20(lines)
	if err != nil {
		t.Fatal(err)
	}
	got := Day20(puzzle, true)
	if got != want {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay20Part2(t *testing.T) {
	const want = 1030809
	lines := linesFromFilename(t, filename(20))
	puzzle, err := NewDay20(lines)
	if err != nil {
		t.Fatal(err)
	}
	got := Day20(puzzle, false)
	if got != want {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay20Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(20))
	b.ResetTimer()
	for range b.N {
		puzzle, err := NewDay20(lines)
		if err != nil {
			b.Fatal(err)
		}
		Day20(puzzle, true)
	}
}

func BenchmarkDay20Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(20))
	b.ResetTimer()
	for range b.N {
		puzzle, err := NewDay20(lines)
		if err != nil {
			b.Fatal(err)
		}
		Day20(puzzle, false)
	}
}
