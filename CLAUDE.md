# Claude Code Instructions - Advent of Code 2024

## Critical Rules

### Function Signatures (PRIMARY RULE)
- **MUST** implement: `func DayXX(<input>) uint`
- **SHOULD** use: `func DayXX(<input>, part1 bool) uint` unless alternatives are more elegant
- **IF** parser required (input cannot be directly processed by `input_test.go` functions):
  - Parser: `func NewDayXX(<input>) DayXXPuzzle` (return by value)
  - Combined: `func DayXX(puzzle DayXXPuzzle) uint`
- **NEVER** use methods: `func (p *DayXXPuzzle) DayXX() uint`

### File Access Prohibition
- Puzzles must not perform I/O
- **ONLY** tests may read files using `input_test.go` functions

### uint Pattern (MANDATORY)
- **ALL** puzzle return types that are counts/sums/totals/amounts must be `uint`
- Push `uint` contract up the entire call chain
- Area, perimeter, distances, prices are inherently non-negative
- Example: `func exploreRegion(...) (area, perimeter uint)`

### Coordinate System
- **ALWAYS** use `x`/`y` throughout (never row/col)
- `dimX` (width), `dimY` (height) for dimensions
- `grid[y][x]` indexing pattern
- `startY, startX int` parameter order

### Data Types
- **ALWAYS** use `byte` for ASCII characters (A-Z, 0-9, symbols)
- **NEVER** use `rune` - unnecessary UTF-8 overhead for AoC
- Use `[]byte(string)` for conversion, not manual loops

### Algorithm Requirements
- **NEVER** use recursion
- **ALWAYS** use iterative with explicit stacks: `[]image.Point`
- Use `image.Point{X: x, Y: y}` for coordinates

### Modern Go Patterns (MANDATORY)
- **ALWAYS** use latest Go 1.24+ features where applicable
- Use `for range N` instead of `for i := 0; i < N; i++` (range over integers)
- Use `slices` package: `slices.Equal`, `slices.Contains`, `slices.Sort`
- Use `maps` package: `maps.Equal`, `maps.Clone` when needed
- Use `clear(map)` and `clear(slice)` for efficient clearing
- Use `min()` and `max()` built-in functions

### Test Structure
- Table-driven tests with external files
- `testdata/dayXX_example1.txt` not inline strings
- **NEVER** use multiline string literals in tests - always use external testdata files
- `lines := linesFromFilename(t, filename)` in tests only
- Multiple examples: use `example1Filename(day)`, `example2Filename(day)`, etc.
- Available filename functions: `exampleFilename()`, `exampleNFilename()`, `example1Filename()`, `example2Filename()`, `example3Filename()`, `filename()`

### Input Parsing (Flexible)
- **Parser is optional** - only use if beneficial for complexity
- `func DayXX(input []byte)` - fine if puzzle can parse bytes directly
- `func DayXX(lines []string)` - fine if puzzle needs line-based input
- `func NewDayXX()` + `DayXX(puzzle)` - use for complex data structures
- Choose the most appropriate input format for each puzzle's needs
- Use appropriate `input_test.go` helper functions

## Template

```go
// dayXX.go
type DayXXPuzzle struct {
    grid [][]byte  // not rune
    dimY int       // height
    dimX int       // width
}

func NewDayXX(lines []string) DayXXPuzzle {
    dimY := len(lines)
    grid := make([][]byte, dimY)
    for y := range grid {  // modern range syntax
        grid[y] = []byte(lines[y])
    }
    return DayXXPuzzle{grid: grid, dimY: dimY, dimX: len(lines[0])}
}

func DayXX(puzzle DayXXPuzzle) uint {  // uint return
    var total uint
    visited := make([][]bool, puzzle.dimY)
    for y := range visited {
        visited[y] = make([]bool, puzzle.dimX)
    }

    for y := range puzzle.dimY {  // range over int (Go 1.24+)
        for x := range puzzle.dimX {
            // algorithm logic
        }
    }
    return total
}

// dayXX_test.go
func TestDayXXPart1Example(t *testing.T) {
    tests := []struct {
        name string
        file string
        want uint  // uint expectation
    }{
        {"example", exampleFilename(XX), 42},
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

## Reference Implementation
See `day12.go` and `day12_test.go` for perfect examples following all patterns.