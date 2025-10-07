package adventofcode2024

import "testing"

func TestCramer(t *testing.T) {
	tests := []struct {
		name     string
		eq1, eq2 Eq
		wantX    int
		wantY    int
		wantOk   bool
	}{
		{
			name:   "simple system",
			eq1:    Eq{2, 3, 7},  // 2x + 3y = 7
			eq2:    Eq{1, -1, 1}, // x - y = 1
			wantX:  2,
			wantY:  1,
			wantOk: true,
		},
		{
			name:   "no solution - zero determinant",
			eq1:    Eq{2, 4, 6}, // 2x + 4y = 6
			eq2:    Eq{1, 2, 4}, // x + 2y = 4 (parallel lines)
			wantX:  0,
			wantY:  0,
			wantOk: false,
		},
		{
			name:   "no integer solution",
			eq1:    Eq{2, 3, 8}, // 2x + 3y = 8  -> x=1, y=2 (actually has integer solution)
			eq2:    Eq{1, 1, 3}, // x + y = 3
			wantX:  1,
			wantY:  2,
			wantOk: true,
		},
		{
			name:   "negative solution",
			eq1:    Eq{1, 1, 0},  // x + y = 0
			eq2:    Eq{1, -1, 2}, // x - y = 2
			wantX:  1,
			wantY:  -1,
			wantOk: true,
		},
		{
			name:   "large numbers",
			eq1:    Eq{94, 22, 8400}, // 94x + 22y = 8400
			eq2:    Eq{34, 67, 5400}, // 34x + 67y = 5400
			wantX:  80,
			wantY:  40,
			wantOk: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY, gotOk := Cramer(tt.eq1, tt.eq2)
			if gotX != tt.wantX || gotY != tt.wantY || gotOk != tt.wantOk {
				t.Errorf("Cramer(%v, %v) = (%d, %d, %t), want (%d, %d, %t)",
					tt.eq1, tt.eq2, gotX, gotY, gotOk, tt.wantX, tt.wantY, tt.wantOk)
			}
		})
	}
}

func TestBareiss(t *testing.T) {
	tests := []struct {
		name    string
		eqs     []Eq
		wantSol []int
		wantOk  bool
	}{
		{
			name: "simple 2x2 system",
			eqs: []Eq{
				{2, 3, 7},  // 2x + 3y = 7
				{1, -1, 1}, // x - y = 1
			},
			wantSol: []int{2, 1},
			wantOk:  true,
		},
		{
			name: "no solution - zero determinant",
			eqs: []Eq{
				{2, 4, 6}, // 2x + 4y = 6
				{1, 2, 4}, // x + 2y = 4
			},
			wantSol: nil,
			wantOk:  false,
		},
		{
			name: "integer solution",
			eqs: []Eq{
				{2, 3, 8}, // 2x + 3y = 8  -> x=1, y=2
				{1, 1, 3}, // x + y = 3
			},
			wantSol: []int{1, 2},
			wantOk:  true,
		},
		{
			name: "negative solution",
			eqs: []Eq{
				{1, 1, 0},  // x + y = 0
				{1, -1, 2}, // x - y = 2
			},
			wantSol: []int{1, -1},
			wantOk:  true,
		},
		{
			name:    "empty system",
			eqs:     []Eq{},
			wantSol: nil,
			wantOk:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSol, gotOk := Bareiss(tt.eqs)
			if gotOk != tt.wantOk {
				t.Errorf("Bareiss(%v) ok = %t, want %t", tt.eqs, gotOk, tt.wantOk)
				return
			}
			if !gotOk {
				return // Expected failure
			}
			if len(gotSol) != len(tt.wantSol) {
				t.Errorf("Bareiss(%v) solution length = %d, want %d", tt.eqs, len(gotSol), len(tt.wantSol))
				return
			}
			for i, got := range gotSol {
				if got != tt.wantSol[i] {
					t.Errorf("Bareiss(%v) solution[%d] = %d, want %d", tt.eqs, i, got, tt.wantSol[i])
				}
			}
		})
	}
}

func TestCramerVsBareiss(t *testing.T) {
	tests := []struct {
		name     string
		eq1, eq2 Eq
	}{
		{
			name: "simple system",
			eq1:  Eq{2, 3, 7},
			eq2:  Eq{1, -1, 1},
		},
		{
			name: "large numbers",
			eq1:  Eq{94, 22, 8400},
			eq2:  Eq{34, 67, 5400},
		},
		{
			name: "negative coefficients",
			eq1:  Eq{-2, 3, 1},
			eq2:  Eq{4, -1, 7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test Cramer
			x1, y1, ok1 := Cramer(tt.eq1, tt.eq2)

			// Test Bareiss
			sol2, ok2 := Bareiss([]Eq{tt.eq1, tt.eq2})

			// Both should have same success/failure
			if ok1 != ok2 {
				t.Errorf("Cramer ok=%t, Bareiss ok=%t, should be equal", ok1, ok2)
				return
			}

			// If both succeeded, solutions should match
			if ok1 && ok2 {
				if len(sol2) != 2 {
					t.Errorf("Bareiss returned %d solutions, want 2", len(sol2))
					return
				}
				if x1 != sol2[0] || y1 != sol2[1] {
					t.Errorf("Cramer=(%d,%d), Bareiss=(%d,%d), should be equal",
						x1, y1, sol2[0], sol2[1])
				}
			}
		})
	}
}

func BenchmarkCramer(b *testing.B) {
	eq1 := Eq{94, 22, 8400}
	eq2 := Eq{34, 67, 5400}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Cramer(eq1, eq2)
	}
}

func BenchmarkBareiss(b *testing.B) {
	eqs := []Eq{
		{94, 22, 8400},
		{34, 67, 5400},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Bareiss(eqs)
	}
}
