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
	// Look for the Christmas tree pattern by finding when robots are most clustered
	maxClusterScore := 0
	bestTime := 0

	// Search through a reasonable range - robots repeat every LCM(dimX, dimY)
	maxTime := dimX * dimY
	if maxTime > 20000 {
		maxTime = 20000 // Reasonable upper bound
	}

	for t := 1; t < maxTime; t++ {
		score := calculateClusterScore(robots, dimX, dimY, t)
		if score > maxClusterScore {
			maxClusterScore = score
			bestTime = t
		}
	}

	return uint(bestTime)
}

func calculateClusterScore(robots [][4]int, dimX, dimY, seconds int) int {
	// Create a grid to count robot positions
	grid := make([][]int, dimY)
	for y := range grid {
		grid[y] = make([]int, dimX)
	}

	// Place robots on grid
	for _, robot := range robots {
		px, py, vx, vy := robot[0], robot[1], robot[2], robot[3]

		finalX := (px + vx*seconds) % dimX
		finalY := (py + vy*seconds) % dimY

		if finalX < 0 {
			finalX += dimX
		}
		if finalY < 0 {
			finalY += dimY
		}

		grid[finalY][finalX]++
	}

	// Calculate clustering score - count robots with neighbors
	score := 0
	for y := 0; y < dimY; y++ {
		for x := 0; x < dimX; x++ {
			if grid[y][x] > 0 {
				// Count neighbors
				neighbors := 0
				for dy := -1; dy <= 1; dy++ {
					for dx := -1; dx <= 1; dx++ {
						if dy == 0 && dx == 0 {
							continue
						}
						ny, nx := y+dy, x+dx
						if ny >= 0 && ny < dimY && nx >= 0 && nx < dimX && grid[ny][nx] > 0 {
							neighbors++
						}
					}
				}
				score += neighbors * grid[y][x]
			}
		}
	}

	return score
}
