package adventofcode2024

import (
	"errors"
	"image"
	"os"
	"testing"
)

func TestDay16Part1Example(t *testing.T) {
	tests := []struct {
		name string
		file string
		want uint
	}{
		{
			name: "1",
			file: example1Filename(16),
			want: 7036,
		},
		{
			name: "2",
			file: example2Filename(16),
			want: 11048,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input, err := os.ReadFile(tt.file)
			if err != nil {
				t.Fatal(err)
			}
			got, err := Day16(input, true)
			if err != nil {
				t.Fatal(err)
			}
			if got != tt.want {
				t.Errorf("Day16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay16Part1(t *testing.T) {
	const want = 107512

	input := file(t, 16)
	got, err := Day16(input, true)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("Day16() = %v, want %v", got, want)
	}
}

func TestDay16Part2Example(t *testing.T) {
	tests := []struct {
		name string
		file string
		want uint
	}{
		{
			name: "1",
			file: example1Filename(16),
			want: 45,
		},
		{
			name: "2",
			file: example2Filename(16),
			want: 64,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input, err := os.ReadFile(tt.file)
			if err != nil {
				t.Fatal(err)
			}
			got, err := Day16(input, false)
			if err != nil {
				t.Fatal(err)
			}
			if got != tt.want {
				t.Errorf("Day16() part 2 = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay16Part2(t *testing.T) {
	const want = 561

	input := file(t, 16)
	got, err := Day16(input, false)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("Day16() part 2 = %v, want %v", got, want)
	}
}

func TestDay16Errors(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:    "no start position",
			input:   []byte("###\n#.E#\n###"),
			wantErr: ErrNoStartFound,
		},
		{
			name:    "no end position",
			input:   []byte("###\n#S.#\n###"),
			wantErr: ErrNoEndFound,
		},
		{
			name:    "no solution - blocked",
			input:   []byte("###\n#S#E#\n###"),
			wantErr: ErrNoSolutionFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Day16(tt.input, true)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Day16() error = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestDay16NoSolutionErrorDetails(t *testing.T) {
	const (
		size = 10
		S    = 3
		E    = 8
	)
	var buf [size*size + size]byte // grid + newlines (including final)

	// Fill with walls and newlines
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			pos := y*(size+1) + x
			buf[pos] = '#'
		}
		buf[y*(size+1)+size] = '\n' // newline after each row
	}

	// Place S and E at specified positions
	start := image.Point{X: S, Y: S}
	end := image.Point{X: E, Y: E}
	buf[start.Y*(size+1)+start.X] = 'S'
	buf[end.Y*(size+1)+end.X] = 'E'

	_, err := Day16(buf[:], true)

	var noSolErr *NoSolutionError
	if !errors.As(err, &noSolErr) {
		t.Fatalf("Expected NoSolutionError, got %T", err)
	}

	if noSolErr.Start != start {
		t.Errorf("Start position = %v, want %v", noSolErr.Start, start)
	}
	if noSolErr.End != end {
		t.Errorf("End position = %v, want %v", noSolErr.End, end)
	}
}
