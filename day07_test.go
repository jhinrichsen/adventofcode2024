package adventofcode2024

import (
	"testing"
)

func TestDay07Part1Example(t *testing.T) {
	testWithParserNoErr(t, 7, exampleFilename, true, NewDay07, Day07, 3749)
}

func TestDay07Part2Example(t *testing.T) {
	testWithParserNoErr(t, 7, exampleFilename, false, NewDay07, Day07, 11387)
}

func TestDay07Part1(t *testing.T) {
	testWithParserNoErr(t, 7, filename, true, NewDay07, Day07, 20281182715321)
}

func TestDay07Part2(t *testing.T) {
	testWithParserNoErr(t, 7, filename, false, NewDay07, Day07, 159490400628354)
}

func BenchmarkDay07Part1(b *testing.B) {
	benchWithParserNoErr(b, 7, true, NewDay07, Day07)
}

func BenchmarkDay07Part2(b *testing.B) {
	benchWithParserNoErr(b, 7, false, NewDay07, Day07)
}
