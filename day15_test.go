package adventofcode2024

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

var day15TestHook func(gridBytes []byte)

func formatGrid(gridBytes []byte, dimX, dimY int) string {
	var result []byte
	for i := 0; i < len(gridBytes); i++ {
		if i > 0 && i%(dimX+1) == dimX {
			continue
		}
		result = append(result, gridBytes[i])
		if (i+1)%(dimX+1) == dimX {
			result = append(result, '\n')
		}
	}
	if len(result) > 0 && result[len(result)-1] == '\n' {
		result = result[:len(result)-1]
	}
	return string(result)
}

func TestDay15Part1Example(t *testing.T) {
	tests := []struct {
		name        string
		file        string
		want        uint
		stepFiles   []string
		expectSteps int
	}{
		{
			name:        "1",
			file:        example1Filename(15),
			want:        10092,
			stepFiles:   []string{"testdata/day15_example1_last_step.txt"},
			expectSteps: 1, // Verify final state only
		},
		{
			name: "2",
			file: example2Filename(15),
			want: 2028,
			stepFiles: func() []string {
				var files []string
				for i := 0; i <= 15; i++ {
					files = append(files, fmt.Sprintf("testdata/day15_example2_step%d.txt", i))
				}
				return files
			}(),
			expectSteps: 15, // 15 hook calls, not 16
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stepIndex int
			var finalState []byte
			var hookCallCount int
			var initialBoxCount int
			day15TestHook = func(gridBytes []byte) {
				hookCallCount++

				if tt.expectSteps == 1 {
					// Count boxes in current state
					boxCount := 0
					for _, b := range gridBytes {
						if b == 'O' {
							boxCount++
						}
					}

					// Set initial count on first call
					if hookCallCount == 1 {
						initialBoxCount = boxCount
					}

					// Check for box count divergence
					if boxCount != initialBoxCount {
						t.Fatalf("Box count changed! Initial: %d, Current: %d at step %d",
							initialBoxCount, boxCount, hookCallCount)
					}

					// Capture final state
					finalState = make([]byte, len(gridBytes))
					copy(finalState, gridBytes)
					return
				}

				stepIndex++
				if stepIndex > len(tt.stepFiles)-1 {
					t.Errorf("Too many hook calls: got call %d, expected max %d", stepIndex, len(tt.stepFiles)-1)
					return
				}

				want, err := os.ReadFile(tt.stepFiles[stepIndex])
				if err != nil {
					t.Fatalf("Failed to read step file %s: %v", tt.stepFiles[stepIndex], err)
				}

				if !bytes.Equal(want, gridBytes) {
					t.Fatalf("Step %d mismatch:\nwant:\n%s\ngot:\n%s",
						stepIndex, string(want), string(gridBytes))
				}
			}
			defer func() { day15TestHook = nil }()

			input, err := os.ReadFile(tt.file)
			if err != nil {
				t.Fatal(err)
			}

			got, err := Day15(input, true)
			if err != nil {
				t.Fatal(err)
			}

			// Check final state for example1
			if tt.expectSteps == 1 && finalState != nil {
				want, err := os.ReadFile(tt.stepFiles[0])
				if err != nil {
					t.Fatalf("Failed to read final step file %s: %v", tt.stepFiles[0], err)
				}

				// Normalize by trimming trailing whitespace and newlines
				wantStr := strings.TrimSpace(string(want))
				gotStr := strings.TrimSpace(string(finalState))

				if wantStr != gotStr {
					t.Fatalf("Final state mismatch:\nwant:\n%s\ngot:\n%s",
						wantStr, gotStr)
				}
			}

			if tt.want != got {
				t.Errorf("want %v, got %v", tt.want, got)
			}
		})
	}
}

func generateAllSteps(t *testing.T, filename string, outputDir string) {
	originalInput, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	// Make a copy for Day15 to modify
	input := make([]byte, len(originalInput))
	copy(input, originalInput)

	dimX := 0
	for i := 0; i < len(input); i++ {
		if input[i] == '\n' {
			dimX = i
			break
		}
	}

	dimY := 0
	for i := 0; i < len(input); i++ {
		if input[i] == '\n' {
			dimY++
			if i+1 < len(input) && input[i+1] == '\n' {
				break
			}
		}
	}

	stepCount := 1

	day15TestHook = func(gridBytes []byte) {
		formatted := formatGrid(gridBytes, dimX, dimY)
		filename := fmt.Sprintf("%s/step%d.txt", outputDir, stepCount)
		err := os.WriteFile(filename, []byte(formatted), 0644)
		if err != nil {
			t.Fatalf("Failed to write step file %s: %v", filename, err)
		}
		stepCount++
		if stepCount%50 == 0 {
			fmt.Printf("Generated step %d\n", stepCount)
		}
	}
	defer func() { day15TestHook = nil }()

	var gridEnd int
	for i := 0; i < len(input)-1; i++ {
		if input[i] == '\n' && input[i+1] == '\n' {
			gridEnd = i // Grid ends at the first newline of the blank line
			break
		}
	}

	// Create initial state BEFORE calling Day15 - use original input
	initialFormatted := string(originalInput[:gridEnd])
	err = os.WriteFile(fmt.Sprintf("%s/step0.txt", outputDir), []byte(initialFormatted), 0644)
	if err != nil {
		t.Fatalf("Failed to write initial step file: %v", err)
	}

	result, err := Day15(input, true)
	if err != nil {
		t.Fatalf("Day15() error = %v", err)
	}

	fmt.Printf("Generated %d step files in %s/ directory\n", stepCount, outputDir)
	fmt.Printf("Final result: %d\n", result)
}

func TestDay15Part1(t *testing.T) {
	const want = 1451928

	got, err := Day15(file(t, 15), true)
	if err != nil {
		t.Fatal(err)
	}

	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestDay15GenerateSteps(t *testing.T) {
	generateAllSteps(t, example1Filename(15), "me")
}

func TestDay15Scaled(t *testing.T) {
	// Test with example1 and compare against the scaled version from spec
	input, err := os.ReadFile(example1Filename(15))
	if err != nil {
		t.Fatal(err)
	}

	expected, err := os.ReadFile("testdata/day15_example1_scaled.txt")
	if err != nil {
		t.Fatal(err)
	}

	result := make([]byte, len(input)*2)
	widen(input, result)

	// Extract just the grid portion (before the blank line)
	resultLines := strings.Split(string(result), "\n")
	expectedLines := strings.Split(string(expected), "\n")

	// Compare grid portions
	minLines := min(len(resultLines), len(expectedLines))
	for i := 0; i < minLines; i++ {
		if resultLines[i] != expectedLines[i] {
			t.Errorf("Line %d mismatch:\nwant: %s\ngot:  %s", i, expectedLines[i], resultLines[i])
		}
	}
}

func TestDay15Part2Example(t *testing.T) {
	const want = 9021 // From spec

	input, err := os.ReadFile(example1Filename(15))
	if err != nil {
		t.Fatal(err)
	}

	got, err := Day15(input, false)
	if err != nil {
		t.Fatal(err)
	}

	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestDay15Part2(t *testing.T) {
	const want = 1462788

	got, err := Day15(file(t, 15), false)
	if err != nil {
		t.Fatal(err)
	}

	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}

func BenchmarkDay15Part1(b *testing.B) {
	input := file(b, 15)
	for range b.N {
		_, _ = Day15(input, true)
	}
}

func BenchmarkDay15Part2(b *testing.B) {
	input := file(b, 15)
	for range b.N {
		_, _ = Day15(input, false)
	}
}
