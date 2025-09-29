package adventofcode2024

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

// Helper functions using testing.TB - only way to access files
func linesFromFilename(tb testing.TB, filename string) []string {
	tb.Helper()
	f, err := os.Open(filename)
	if err != nil {
		tb.Fatal(err)
	}
	defer f.Close()

	var lines []string
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		lines = append(lines, line)
	}
	if err := sc.Err(); err != nil {
		tb.Fatal(err)
	}

	// Reset timer if this is a benchmark
	if b, ok := tb.(*testing.B); ok {
		b.ResetTimer()
	}

	return lines
}

func exampleFilename(day uint8) string {
	return fmt.Sprintf("testdata/day%02d_example.txt", day)
}

func exampleNFilename(day uint8, n int) string {
	return fmt.Sprintf("testdata/day%02d_example%d.txt", day, n)
}

func example1Filename(day uint8) string {
	return exampleNFilename(day, 1)
}

func example2Filename(day uint8) string {
	return exampleNFilename(day, 2)
}

func example3Filename(day uint8) string {
	return exampleNFilename(day, 3)
}

func example4Filename(day uint8) string {
	return exampleNFilename(day, 4)
}

func filename(day uint8) string {
	return fmt.Sprintf("testdata/day%02d.txt", day)
}

func file(tb testing.TB, day uint8) []byte {
	tb.Helper()
	buf, err := os.ReadFile(filename(day))
	if err != nil {
		tb.Fatal(err)
	}

	// Reset timer if this is a benchmark
	if b, ok := tb.(*testing.B); ok {
		b.ResetTimer()
	}

	return buf
}

func exampleFile(tb testing.TB, day uint8) []byte {
	tb.Helper()
	buf, err := os.ReadFile(exampleFilename(day))
	if err != nil {
		tb.Fatal(err)
	}

	// Reset timer if this is a benchmark
	if b, ok := tb.(*testing.B); ok {
		b.ResetTimer()
	}

	return buf
}

const (
	MagicMaxLines    = 3999  // maximum number of lines for any puzzle input
	MagicLongestLine = 19999 // longest line of any puzzle input
)

func TestLinesFromFilename(t *testing.T) {
	lines := linesFromFilename(t, "testdata/day01.txt")
	const want = 1000
	got := len(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
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
