package adventofcode2024

import (
	"testing"
)

func TestDay06Part1Example(t *testing.T) {
	testWithParserNoErr(t, 6, exampleFilename, true, NewDay06, Day06, 41)
}

func TestDay06Part2Example(t *testing.T) {
	testWithParserNoErr(t, 6, exampleFilename, false, NewDay06, Day06, 6)
}

func TestDay06Part1(t *testing.T) {
	testWithParserNoErr(t, 6, filename, true, NewDay06, Day06, 4903)
}

func TestDay06Part2(t *testing.T) {
	testWithParserNoErr(t, 6, filename, false, NewDay06, Day06, 1911)
}

func BenchmarkDay06Part1(b *testing.B) {
	benchWithParserNoErr(b, 6, true, NewDay06, Day06)
}

func BenchmarkDay06Part2(b *testing.B) {
	benchWithParserNoErr(b, 6, false, NewDay06, Day06)
}
