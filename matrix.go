package adventofcode2024

// Eq represents a linear equation: A*x + B*y = X
type Eq struct {
	A, B, X int
}

// Cramer solves a 2x2 system using Cramer's rule
// eq1: A*x + B*y = X
// eq2: A*x + B*y = X
// Returns x, y, and whether integer solution exists
func Cramer(eq1, eq2 Eq) (int, int, bool) {
	// Calculate determinant of coefficient matrix
	det := eq1.A*eq2.B - eq1.B*eq2.A
	if det == 0 {
		return 0, 0, false // No unique solution
	}

	// Calculate numerators using Cramer's rule
	numeratorX := eq1.X*eq2.B - eq2.X*eq1.B
	numeratorY := eq1.A*eq2.X - eq2.A*eq1.X

	// Check if solutions are integers
	if numeratorX%det != 0 || numeratorY%det != 0 {
		return 0, 0, false // No integer solution
	}

	x := numeratorX / det
	y := numeratorY / det

	return x, y, true
}

// Bareiss solves an n×n system using Bareiss algorithm (fraction-free Gaussian elimination)
// Returns solution vector and whether solution exists
func Bareiss(eqs []Eq) ([]int, bool) {
	n := len(eqs)
	if n == 0 {
		return nil, false
	}

	// Create augmented matrix
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n+1)
		matrix[i][0] = eqs[i].A
		if n > 1 {
			matrix[i][1] = eqs[i].B
		}
		// For now, only support 2-variable systems
		if n > 2 {
			return nil, false // Not implemented for >2 variables yet
		}
		matrix[i][n] = eqs[i].X
	}

	// Forward elimination with Bareiss algorithm
	for k := 0; k < n-1; k++ {
		// Find pivot
		pivot := k
		for i := k + 1; i < n; i++ {
			if abs(matrix[i][k]) > abs(matrix[pivot][k]) {
				pivot = i
			}
		}

		// Swap rows if needed
		if pivot != k {
			matrix[k], matrix[pivot] = matrix[pivot], matrix[k]
		}

		// Check for zero pivot
		if matrix[k][k] == 0 {
			return nil, false
		}

		// Eliminate column k
		for i := k + 1; i < n; i++ {
			for j := k + 1; j <= n; j++ {
				if k == 0 {
					matrix[i][j] = matrix[k][k]*matrix[i][j] - matrix[i][k]*matrix[k][j]
				} else {
					if matrix[k-1][k-1] == 0 {
						return nil, false // Division by zero
					}
					matrix[i][j] = (matrix[k][k]*matrix[i][j] - matrix[i][k]*matrix[k][j]) / matrix[k-1][k-1]
				}
			}
			matrix[i][k] = 0
		}
	}

	// Back substitution
	solution := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		if matrix[i][i] == 0 {
			return nil, false // Zero on diagonal
		}
		
		sum := matrix[i][n]
		for j := i + 1; j < n; j++ {
			sum -= matrix[i][j] * solution[j]
		}

		// Check if solution is integer
		if sum%matrix[i][i] != 0 {
			return nil, false
		}
		solution[i] = sum / matrix[i][i]
	}

	return solution, true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
