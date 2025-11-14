package adventofcode2024

import "testing"

func TestDay23Part1Example(t *testing.T) {
	const want = 7
	got := Day23(NewDay23(linesFromFilename(t, exampleFilename(23))), true)
	if got != want {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay23Part1(t *testing.T) {
	const want = 1230
	got := Day23(NewDay23(linesFromFilename(t, filename(23))), true)
	if got != want {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay23Part2Example(t *testing.T) {
	const want = "co,de,ka,ta"
	got := Day23Part2Password(NewDay23(linesFromFilename(t, exampleFilename(23))))
	if got != want {
		t.Fatalf("want %q but got %q", want, got)
	}
}

func TestDay23Part2(t *testing.T) {
	const want = "az,cj,kp,lm,lt,nj,rf,rx,sn,ty,ui,wp,zo"
	got := Day23Part2Password(NewDay23(linesFromFilename(t, filename(23))))
	if got != want {
		t.Fatalf("want %q but got %q", want, got)
	}
}

func BenchmarkDay23Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(23))
	for range b.N {
		Day23(NewDay23(lines), true)
	}
}

func BenchmarkDay23Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(23))
	for range b.N {
		Day23Part2Password(NewDay23(lines))
	}
}
