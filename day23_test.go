package adventofcode2024

import "testing"

func TestDay23Part1Example(t *testing.T) {
	testWithParserNoErr(t, 23, exampleFilename, true, NewDay23, Day23, "7")
}

func TestDay23Part1(t *testing.T) {
	testWithParserNoErr(t, 23, filename, true, NewDay23, Day23, "1230")
}

func TestDay23Part2Example(t *testing.T) {
	testWithParserNoErr(t, 23, exampleFilename, false, NewDay23, Day23, "co,de,ka,ta")
}

func TestDay23Part2(t *testing.T) {
	testWithParserNoErr(t, 23, filename, false, NewDay23, Day23, "az,cj,kp,lm,lt,nj,rf,rx,sn,ty,ui,wp,zo")
}

func BenchmarkDay23Part1(b *testing.B) {
	benchWithParserNoErr(b, 23, true, NewDay23, Day23)
}

func BenchmarkDay23Part2(b *testing.B) {
	benchWithParserNoErr(b, 23, false, NewDay23, Day23)
}
