package adventofcode2024

import (
	"fmt"
	"slices"
	"testing"
)

func TestDay11Part1Examples(t *testing.T) {
	var (
		t0 = []uint64{125, 17}
		t1 = []uint64{253000, 1, 7}
		t2 = []uint64{253, 0, 2024, 14168}
		t3 = []uint64{512072, 1, 20, 24, 28676032}
		t4 = []uint64{512, 72, 2024, 2, 0, 2, 4, 2867, 6032}
		t5 = []uint64{1036288, 7, 2, 20, 24, 4048, 1, 4048, 8096, 28, 67, 60, 32}
		t6 = []uint64{2097446912, 14168, 4048, 2, 0, 2, 4, 40, 48, 2024, 40, 48, 80, 96, 2, 8, 6, 7, 6, 0, 3, 2}
		tt = []struct {
			from, into []uint64
		}{
			{t0, t1},
			{t1, t2},
			{t2, t3},
			{t3, t4},
			{t4, t5},
			{t5, t6},
		}
	)
	for i := range tt {
		id := fmt.Sprintf("%+v", tt[i].from)
		t.Run(id, func(t *testing.T) {
			want := tt[i].into
			got := blink(tt[i].from)
			if !slices.Equal(want, got) {
				t.Fatalf("want %v but got %v", want, got)
			}
		})
	}
}

func TestDay11Part2Example(t *testing.T) {
	const want = 0
}

func TestDay11Part1(t *testing.T) {
	const want = 185517
	got := Day11([]uint64{64554, 35, 906, 6, 6960985, 5755, 975820, 0})
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay11Part2(t *testing.T) {
	const want = 0
	got := Day11([]uint64{64554, 35, 906, 6, 6960985, 5755, 975820, 0})
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay11Part1(b *testing.B) {
	for range b.N {
		_ = Day11([]uint64{64554, 35, 906, 6, 6960985, 5755, 975820, 0})
	}
}

func BenchmarkDay11Part2(b *testing.B) {
	for range b.N {
		_ = Day11([]uint64{64554, 35, 906, 6, 6960985, 5755, 975820, 0})
	}
}
