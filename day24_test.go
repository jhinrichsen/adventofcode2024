package adventofcode2024

import "testing"

func TestDay24Part1Example1(t *testing.T) {
	testWithParser(t, 24, example1Filename, true, NewDay24, Day24, "4")
}

func TestDay24Part1Example(t *testing.T) {
	testWithParser(t, 24, exampleFilename, true, NewDay24, Day24, "2024")
}

func TestDay24Part1(t *testing.T) {
	testWithParser(t, 24, filename, true, NewDay24, Day24, "59336987801432")
}

func TestDay24Part2(t *testing.T) {
	testWithParser(t, 24, filename, false, NewDay24, Day24, "ctg,dmh,dvq,rpb,rpv,z11,z31,z38")
}

func BenchmarkDay24Part1(b *testing.B) {
	benchWithParser(b, 24, true, NewDay24, Day24)
}
