package adventofcode2024

import (
	"testing"
)

func TestDay04Part1Example(t *testing.T) {
	testWithParserNoErr(t, 4, exampleFilename, true, NewDay04, Day04, 18)
}

func TestDay04Part1(t *testing.T) {
	testWithParserNoErr(t, 4, filename, true, NewDay04, Day04, 2685)
}

func BenchmarkDay04Part1(b *testing.B) {
	benchWithParserNoErr(b, 4, true, NewDay04, Day04)
}

func TestDay04Part2Example(t *testing.T) {
	testWithParserNoErr(t, 4, exampleFilename, false, NewDay04, Day04, 9)
}

func TestDay04Part2(t *testing.T) {
	testWithParserNoErr(t, 4, filename, false, NewDay04, Day04, 2048)
}

func BenchmarkDay04Part2(b *testing.B) {
	benchWithParserNoErr(b, 4, false, NewDay04, Day04)
}
