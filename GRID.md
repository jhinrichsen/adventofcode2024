# Grid.go Rewrite Opportunities for AoC 2024

Analysis of AoC 2024 puzzles that could be refactored to use the new `grid.go` iterator-based approach instead of manual direction vectors.

## Summary

| Day | Puzzle | Grid Type | Current Approach | Rewrite Benefit |
|-----|--------|-----------|------------------|-----------------|
| 4 | Ceres Search | C8 | Manual 8 directions | **None** (benchmarked) |
| 6 | Guard Gallivant | C4 + direction | Manual directions | Medium |
| 10 | Hoof It | C4 | Flat array + slices | **Done** - 77-81% faster |
| 12 | Garden Groups | C4 | Union-find + C4Indices | **Done** - 340x fewer allocs |
| 14 | Restroom Redoubt | C4/C8 | Manual neighbors | Medium |
| 15 | Warehouse Woes | Movement | Index arithmetic | Low |
| 16 | Reindeer Maze | C4 + direction | Manual directions | Medium |
| 18 | RAM Run | C4 | Flat array + BFS | **Done** - 38% faster |
| 20 | Race Condition | C4 | Flat array + BFS | **Done** - 45% faster |
| 21 | Keypad Conundrum | Small grids | (not implemented) | Low |

## High Priority Rewrites

### Day 4: Ceres Search - BENCHMARKED

**Benchmark Results** (see `day04_bench_test.go`):

| Implementation | Part 1 | Part 2 | Allocs |
|----------------|--------|--------|--------|
| Original (image.Point.Mul) | 84µs | 19µs | 0 |
| C8Points iterator | - | 480µs | 10k |
| Simplified (direct arithmetic) | 61µs | 22µs | 0 |

**Finding:** C8Points iterator is **25x slower** for Part 2 due to closure overhead.

**Current:** Manual 8-direction vectors with `image.Point.Mul()`
```go
for _, dp := range []image.Point{N, NE, E, SE, S, SW, W, NW} {
    p3 := p0.Add(dp.Mul(l - 1))  // Mul() is slow
    ...
}
```

**Recommended:** Avoid `image.Point.Mul()`, use direct arithmetic
```go
endX, endY := x+dp.X*3, y+dp.Y*3
if puzzle.grid[y+dp.Y][x+dp.X] == 'M' && ...
```

**Conclusion:** Grid.go iterators don't help Day 4. The speedup comes from avoiding `Point.Mul()`, not from iterator patterns. This puzzle needs directional rays, not neighbor enumeration.

---

### Day 10: Hoof It - BENCHMARKED

**Benchmark Results** (see `day10_bench_test.go`):

| Implementation | Part 1 | Part 2 | Allocs | vs Original |
|----------------|--------|--------|--------|-------------|
| Original (container/list) | 650µs | 592µs | 13k | baseline |
| Slice-based | 226µs | 168µs | 112 | **3-3.5x faster** |
| Flat array | 152µs | 104µs | 111 | **4-6x faster** |
| Grid C4Points | 597µs | 542µs | 6k | ~same |
| Grid C4Indices | 270µs | 213µs | 6k | 2-3x faster |

**Finding:** The bottleneck is `container/list`, not direction iteration. Grid.go doesn't help because building the neighbor lookup table costs 6k allocations.

**Current:** Uses `container/list` with manual direction iteration
```go
var all = list.New()
for _, delta := range []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
    p := trail[i-1].Add(delta)
    if !p.In(board) { continue }
    all.PushBack(trail)
}
```

**Recommended:** Replace `container/list` with slices, use flat array
```go
var current, next []trailState
for _, ts := range current {
    for _, d := range dirs {
        ni := ts.posIdx + d
        if flat[ni] == height { next = append(next, ...) }
    }
}
current = next
```

**Conclusion:** Grid.go iterators don't help Day 10. The 4-6x speedup comes from using slices instead of `container/list`.

---

### Day 12: Garden Groups - BENCHMARKED

**Benchmark Results** (see `day12_bench_test.go`):

| Implementation | Part 1 | Part 2 | P1 Allocs | vs Original |
|----------------|--------|--------|-----------|-------------|
| Original (DFS) | 480µs | 4.87ms | 1.7k | baseline |
| Flat (DFS) | 286µs | 2.22ms | 1.5k | **1.7x faster** |
| GridUnionFind (iterator) | 474µs | 2.68ms | **5** | **~same, 340x fewer allocs** |

**Finding:** Grid.go C4Indices works well when used correctly:
- **Wrong:** Materialize iterator into slice → 60k allocations, 3x slower
- **Right:** Direct iteration with union-find → **5 allocations**, matches Original speed

**Correct iterator usage (scan + union-find):**
```go
// Single scan: union same-type neighbors
for idx, neighbors := range g.C4Indices() {
    for ni := range neighbors {
        if flat[idx] == flat[ni] { union(idx, ni) }
    }
}
// Second scan: accumulate per-region stats
for idx, neighbors := range g.C4Indices() {
    root := find(idx)
    regionArea[root]++
    sameType := 0
    for ni := range neighbors { if flat[ni] == flat[idx] { sameType++ } }
    regionPerimeter[root] += uint(4 - sameType)
}
```

**Conclusion:** Grid.go iterator is efficient (~20% overhead) when:
1. Used for full grid scans (not random access during DFS)
2. Not materialized into slices
3. Algorithm adapts to scan pattern (union-find instead of DFS)

Flat DFS is still fastest for this problem, but GridUnionFind demonstrates the iterator's potential.

---

### Day 18: RAM Run - OPTIMIZED

**Benchmark Results:**

| Implementation | Part 1 | Part 2 | Allocs | vs Original |
|----------------|--------|--------|--------|-------------|
| Original ([][]byte, image.Point) | 173µs | 755µs | 1,697 / 12,690 | baseline |
| Flat array + index arithmetic | 144µs | 466µs | 1,043 / 3,506 | **17-38% faster** |
| Grid C4Indices (pre-computed) | 499µs | 746µs | 16,163 / 18,625 | **3x slower** |

**Finding:** Grid.go C4Indices is 3x slower for BFS pathfinding due to neighbor table allocation overhead (16k allocs). Flat arrays with inline bounds checking are optimal.

**Optimizations applied:**
1. Flat `[]byte` grid instead of `[][]byte`
2. Flat `[]bool` visited instead of `[][]bool`
3. Integer indices instead of `image.Point`
4. Pre-computed direction offsets as `[4]int{-dimX, 1, dimX, -1}`
5. Ring buffer queue (head index instead of slice shifting)

**Conclusion:** Grid.go iterators don't help BFS pathfinding. The optimization comes from flat arrays and avoiding image.Point overhead.

---

### Day 20: Race Condition - OPTIMIZED

**Benchmark Results:**

| Implementation | Part 1 | Part 2 | Allocs | vs Original |
|----------------|--------|--------|--------|-------------|
| Original ([][]byte, image.Point) | 1.54ms | 46.0ms | 28,894 | baseline |
| Flat array + index arithmetic | 0.87ms | 24.9ms | 11 | **44-46% faster, 99.96% fewer allocs** |

**Finding:** The massive allocation reduction (28,894 → 11) comes from:
- Flat `[]uint` distances instead of `[][]uint` (eliminates row allocations)
- Ring buffer queue instead of slice shifting

**Optimizations applied:**
1. Flat `[]byte` grid instead of `[][]byte`
2. Flat `[]uint` distances instead of `[][]uint`
3. Integer indices instead of `image.Point`
4. Pre-computed direction offsets as `[4]int{-dimX, 1, dimX, -1}`
5. Ring buffer queue (head index instead of slice shifting)
6. Optimized manhattan distance iteration (skip invalid dx range)

**Conclusion:** Same pattern as Day 18 - flat arrays with index arithmetic beat grid.go iterators for BFS-based algorithms.

---

## Medium Priority Rewrites

### Day 6: Guard Gallivant

**Current:** Manual direction with rotation
```go
dir := image.Point{0, -1}
// ...
dir = image.Point{-dir.Y, dir.X} // rotate right
```

**Note:** This puzzle needs directional state (facing), not just neighbor enumeration. The grid.go iterators help less here, but could be extended with a `Directions4` constant.

---

### Day 14: Restroom Redoubt (C4/C8 for clustering)

**Current:** Manual neighbor checks for clustering score
```go
if occupied[[2]int{x-1, y}] { score++ }
if occupied[[2]int{x+1, y}] { score++ }
if occupied[[2]int{x, y-1}] { score++ }
if occupied[[2]int{x, y+1}] { score++ }
```

**Note:** This uses sparse occupancy (map), not dense grid iteration. Could benefit from a standalone `Neighbors4(p)` or `Neighbors8(p)` function rather than full grid iteration.

---

### Day 16: Reindeer Maze

**Current:** Dijkstra with manual directions
```go
dirs := []image.Point{{X: 0, Y: -1}, {X: 1, Y: 0}, {X: 0, Y: 1}, {X: -1, Y: 0}}
newPos := current.state.pos.Add(dirs[current.state.dir])
```

**Note:** Needs direction index for state tracking. Could use shared direction constants from grid.go.

---

## Proposed grid.go Extensions

Based on 2024 puzzle patterns, consider adding:

```go
// Direction constants
var (
    N  = image.Point{0, -1}
    NE = image.Point{1, -1}
    E  = image.Point{1, 0}
    SE = image.Point{1, 1}
    S  = image.Point{0, 1}
    SW = image.Point{-1, 1}
    W  = image.Point{-1, 0}
    NW = image.Point{-1, -1}

    Dirs4 = []image.Point{N, E, S, W}
    Dirs8 = []image.Point{N, NE, E, SE, S, SW, W, NW}
)

// Neighbors4 returns 4-connected neighbors of p within bounds
func (g Grid) Neighbors4(p image.Point) iter.Seq[image.Point]

// Neighbors8 returns 8-connected neighbors of p within bounds
func (g Grid) Neighbors8(p image.Point) iter.Seq[image.Point]

// RotateRight rotates direction 90 degrees clockwise
func RotateRight(dir image.Point) image.Point {
    return image.Point{-dir.Y, dir.X}
}

// RotateLeft rotates direction 90 degrees counter-clockwise
func RotateLeft(dir image.Point) image.Point {
    return image.Point{dir.Y, -dir.X}
}
```

## Non-Grid Puzzles (2024)

These puzzles do not benefit from grid.go:

- Day 1: List comparison
- Day 2: Sequence validation
- Day 3: String parsing
- Day 5: Topological ordering
- Day 7: Operator combinations
- Day 9: Disk compaction (1D)
- Day 11: Stone transformation
- Day 13: Linear algebra (2D but not grid traversal)
- Day 17: VM simulation
- Day 19: String composition
- Day 22: PRNG sequences
- Day 23: Graph cliques
- Day 24: Logic gate simulation
- Day 25: Column height matching
