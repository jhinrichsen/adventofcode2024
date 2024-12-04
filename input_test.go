package adventofcode2024

import (
	"bufio"
	"bytes"
	"os"
	"path/filepath"
	"slices"
	"testing"
)

func TestLinesFromFilename(t *testing.T) {
	lines, err := linesFromFilename("testdata/day00.txt")
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 1 {
		t.Fatalf("want 1 line but got %d", len(lines))
	}

}

func TestNumbersFromFilename(t *testing.T) {
	var (
		line0 = []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	)
	ns := mustNumbersFromFilename("testdata/numbers.txt")
	if len(ns) != 3 {
		t.Fatalf("want 3 lines but got %d", len(ns))
	}
	if !slices.Equal(line0, ns[0]) {
		t.Fatalf("line #1: want %v but got %v", line0, ns[0])
	}

	if len(ns[1]) != 10 {
		t.Fatalf("line #2: want 10 numbers but got %d", len(ns[1]))
	}

	if ns[2][0] != 1 {
		t.Fatalf("line #3: want 1 but got %d", ns[2][0])
	}
	last := len(ns[2]) - 1
	if ns[2][last] != 16384 {
		t.Fatalf("line #3: want 16384 but got %d", ns[2][last])
	}
}

func TestLinesAsNumbers(t *testing.T) {
	sample := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	got, err := linesAsNumbers(sample)
	if err != nil {
		t.Fatal(err)
	}
	if !slices.Equal(want, got) {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func BenchmarkBytesFromFilename(b *testing.B) {
	filenames, err := filepath.Glob("testdata/*.txt")
	if err != nil {
		b.Fatal(err)
	}
	readall := func() {
		for i := range filenames {
			_, _ = bytesFromFilename(filenames[i])
		}
	}
	// warm-up cache
	readall()
	b.ResetTimer()
	for range b.N {
		readall()
	}
}

func BenchmarkLinesFromFilename(b *testing.B) {
	filenames, err := filepath.Glob("testdata/*.txt")
	if err != nil {
		b.Fatal(err)
	}
	readall := func() {
		for i := range filenames {
			_, _ = linesFromFilename(filenames[i])
		}
	}
	// warm-up cache
	readall()
	b.ResetTimer()
	for range b.N {
		readall()
	}
}

func TestMagicConstants(t *testing.T) {
	filenames, err := filepath.Glob("testdata/*.txt")
	if err != nil {
		t.Fatal(err)
	}

	var gotLongestLine, gotMaxLines uint
	for i := range filenames {
		var lines uint
		buf, err := os.ReadFile(filenames[i])
		if err != nil {
			t.Fatal(err)
		}

		scanner := bufio.NewScanner(bytes.NewReader(buf))

		for scanner.Scan() {
			line := scanner.Text()
			lineLength := uint(len(line))
			if lineLength > gotLongestLine {
				gotLongestLine = lineLength
			}
			lines++
		}

		if err := scanner.Err(); err != nil {
			t.Fatal(err)
		}

		gotMaxLines = max(gotMaxLines, lines)
	}
	if MagicMaxLines != gotMaxLines {
		t.Fatalf("want %d but got %d", MagicMaxLines, gotMaxLines)
	}
	if MagicLongestLine != gotLongestLine {
		t.Fatalf("want %d but got %d", MagicLongestLine, gotLongestLine)
	}
}
