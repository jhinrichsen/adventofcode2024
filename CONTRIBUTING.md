# Contributing Guidelines - Advent of Code 2024

## Architecture Principles

### File Access and Input Handling

**Rule: Puzzles must never access files directly**

- ✅ Tests handle all file I/O using `input_test.go` functions
- ✅ Puzzles only accept parsed data as parameters
- ❌ No `os.ReadFile`, `os.Open`, or any file operations in puzzle code

**Input Pattern:**
```go
// In test file
lines := linesFromFilename(t, "testdata/dayXX_example.txt")
puzzle := NewDayXX(lines)
result := DayXX(puzzle)
```

**Available input functions (all require `testing.TB`):**
- `linesFromFilename(tb, filename) []string` - for line-based input
- `file(tb, day) []byte` - for direct byte access
- `exampleFile(tb, day) []byte` - for example file byte access

### Function Signatures

**Pure Functions Only:**
- ✅ `func DayXX(puzzle DayXXPuzzle) uint` - pure function, immutable input
- ❌ `func (p *DayXXPuzzle) DayXX() uint` - method with mutable state

**Parser Functions:**
- ✅ `func NewDayXX(lines []string) DayXXPuzzle` - return by value
- ✅ `func NewDayXX(input []byte) DayXXPuzzle` - if custom parsing needed
- Use appropriate input type based on parsing requirements

### Data Types

**Return Types for Counts/Sums/Totals:**
- ✅ `uint` for all positive numeric results (counts, sums, distances, prices)
- Push `uint` contract as far up the call chain as possible
- Area, perimeter, prices are inherently non-negative

**Grid Coordinates:**
- ✅ Use `x`/`y` coordinate system throughout
- ✅ `dimX` (width), `dimY` (height) for dimensions
- ✅ `grid[y][x]` indexing (row first, column second)
- ✅ Function parameters: `startY, startX int`

**Character Handling:**
- ✅ `byte` for ASCII-only input (A-Z, 0-9, symbols)
- ❌ `rune` - unnecessary UTF-8 overhead for AoC problems
- ✅ `[]byte(string)` for direct string→byte conversion

### Code Organization

**File Structure:**
```
dayXX.go        - puzzle logic only
dayXX_test.go   - tests with file I/O
testdata/
  dayXX_example1.txt
  dayXX_example2.txt
  dayXX.txt
```

**Never Use Recursion:**
- ✅ Use iterative algorithms with explicit stacks/queues
- ✅ `stack := []image.Point{{X: startX, Y: startY}}`
- Performance and stack safety are priorities

### Testing Patterns

**Table-Driven Tests:**
```go
func TestDayXXPart1Example(t *testing.T) {
    tests := []struct {
        name string
        file string
        want uint
    }{
        {"example1", "testdata/dayXX_example1.txt", 42},
        {"example2", "testdata/dayXX_example2.txt", 123},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            lines := linesFromFilename(t, tt.file)
            puzzle := NewDayXX(lines)
            got := DayXX(puzzle)
            if got != tt.want {
                t.Errorf("DayXX() = %v, want %v", got, tt.want)
            }
        })
    }
}
```

**Separate example files, not inline strings:**
- ✅ `testdata/dayXX_example1.txt`
- ❌ Multiline strings in test code

### Data Structures

**Coordinate Handling:**
```go
// Use image.Point for 2D coordinates
import "image"

type DayXXPuzzle struct {
    grid [][]byte  // not [][]rune
    dimY int       // height (rows)
    dimX int       // width (cols)
}

// Directions for grid traversal
directions := []image.Point{
    {X: 0, Y: -1}, // up
    {X: 0, Y: 1},  // down
    {X: -1, Y: 0}, // left
    {X: 1, Y: 0},  // right
}
```

**Algorithm State:**
- ✅ Local state in pure functions (e.g., `visited [][]bool`)
- ❌ Persistent mutable state in structs
- Keep data structures immutable, computation stateless

### Performance Guidelines

**Memory Efficiency:**
- Use `byte` instead of `rune` for ASCII
- Use `[]byte(string)` instead of manual character loops
- Leverage Go's built-in conversions

**Type Safety:**
- Use `uint` for inherently positive values
- Push type contracts up the call chain
- Let the type system prevent invalid states

## Example Implementation

See `day12.go` and `day12_test.go` for reference implementation following all these patterns.