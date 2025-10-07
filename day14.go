package adventofcode2024

func Day14(buf []byte, dimX, dimY int, seconds uint, part1 bool) uint {
	parseInt := func(i int) (int, int) {
		negative := false
		if i < len(buf) && buf[i] == '-' {
			negative = true
			i++
		}

		value := 0
		for i < len(buf) && buf[i] >= '0' && buf[i] <= '9' {
			value = value*10 + int(buf[i]-'0')
			i++
		}

		if negative {
			value = -value
		}

		return value, i
	}

	// Parse robots once
	var robots [][4]int // [px, py, vx, vy]
	i := 0
	for i < len(buf) {
		// Skip "p="
		i += 2

		// Parse px
		px, newI := parseInt(i)
		i = newI + 1 // skip comma

		// Parse py
		py, newI := parseInt(i)
		i = newI + 3 // skip " v="

		// Parse vx
		vx, newI := parseInt(i)
		i = newI + 1 // skip comma

		// Parse vy
		vy, newI := parseInt(i)
		i = newI

		// Skip to next line
		for i < len(buf) && buf[i] != '\n' {
			i++
		}
		if i < len(buf) {
			i++ // skip '\n'
		}

		robots = append(robots, [4]int{px, py, vx, vy})
	}

	if part1 {
		return solvePart1(robots, dimX, dimY, int(seconds))
	} else {
		return solvePart2(robots, dimX, dimY)
	}
}

func solvePart1(robots [][4]int, dimX, dimY, seconds int) uint {
	var sectors [4]uint
	halfX, halfY := dimX/2, dimY/2

	for _, robot := range robots {
		px, py, vx, vy := robot[0], robot[1], robot[2], robot[3]

		// Direct O(1) position calculation
		finalX := (px + vx*seconds) % dimX
		finalY := (py + vy*seconds) % dimY

		// Normalize to positive coordinates
		if finalX < 0 {
			finalX += dimX
		}
		if finalY < 0 {
			finalY += dimY
		}

		// Skip robots on middle axes
		if finalX == halfX || finalY == halfY {
			continue
		}

		// Quadrant mapping
		quadrant := 0
		if finalX > halfX {
			quadrant += 1
		}
		if finalY > halfY {
			quadrant += 2
		}

		sectors[quadrant]++
	}

	return sectors[0] * sectors[1] * sectors[2] * sectors[3]
}

func solvePart2(robots [][4]int, dimX, dimY int) uint {
	// Pre-allocate reusable data structures to avoid allocations
	occupied := make(map[[2]int]bool, len(robots))
	robotPositions := make([][2]int, len(robots)) // Pre-sized slice

	maxClusterScore := 0
	bestTime := 0

	// Search through a reasonable range - robots repeat every LCM(dimX, dimY)
	maxTime := dimX * dimY
	if maxTime > 20000 {
		maxTime = 20000 // Reasonable upper bound
	}

	// Check a smaller range first for quick patterns
	quickSearchLimit := min(maxTime, 10000)

	for t := 1; t < quickSearchLimit; t++ {
		score := calculateClusterScoreUltra(robots, dimX, dimY, t, occupied, robotPositions)
		if score > maxClusterScore {
			maxClusterScore = score
			bestTime = t

			// Very aggressive early termination for obvious Christmas tree
			if score > len(robots)*2 { // Most robots clustered
				break
			}
		}

		// Skip iterations when score is very low (anti-clustering)
		if t > 500 && score < maxClusterScore/3 {
			t += 4 // Skip ahead when clearly no pattern
		}
	}

	// If we didn't find a strong pattern, extend search
	if maxClusterScore < len(robots) && quickSearchLimit < maxTime {
		for t := quickSearchLimit; t < maxTime; t += 5 { // Sample every 5th iteration
			score := calculateClusterScoreUltra(robots, dimX, dimY, t, occupied, robotPositions)
			if score > maxClusterScore {
				maxClusterScore = score
				bestTime = t

				// Fine-tune around the best found time
				for fine := max(1, t-4); fine <= min(maxTime-1, t+4); fine++ {
					if fine == t { continue }
					fineScore := calculateClusterScoreUltra(robots, dimX, dimY, fine, occupied, robotPositions)
					if fineScore > maxClusterScore {
						maxClusterScore = fineScore
						bestTime = fine
					}
				}
				break
			}
		}
	}

	return uint(bestTime)
}


func calculateClusterScoreUltra(robots [][4]int, dimX, dimY, seconds int, occupied map[[2]int]bool, robotPositions [][2]int) int {
	// Clear the reused map - much faster than allocating new map
	for k := range occupied {
		delete(occupied, k)
	}

	// Calculate robot positions and mark occupied cells (reuse slice)
	for i, robot := range robots {
		px, py, vx, vy := robot[0], robot[1], robot[2], robot[3]

		finalX := (px + vx*seconds) % dimX
		finalY := (py + vy*seconds) % dimY

		if finalX < 0 {
			finalX += dimX
		}
		if finalY < 0 {
			finalY += dimY
		}

		pos := [2]int{finalX, finalY}
		occupied[pos] = true
		robotPositions[i] = pos // Reuse pre-allocated slice
	}

	// Count clustering - accept some double counting for speed
	score := 0
	for i := 0; i < len(robots); i++ {
		pos := robotPositions[i]
		x, y := pos[0], pos[1]

		// Count neighbors using map lookups (4 cardinal directions)
		if occupied[[2]int{x-1, y}] { score++ }
		if occupied[[2]int{x+1, y}] { score++ }
		if occupied[[2]int{x, y-1}] { score++ }
		if occupied[[2]int{x, y+1}] { score++ }
	}

	return score
}

