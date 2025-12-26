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
	// X and Y coordinates move independently!
	// X repeats every dimX steps, Y repeats every dimY steps
	// Find time with minimum variance for each axis, then combine with CRT

	n := len(robots)

	// Find time tx (0 to dimX-1) where X-coordinates have minimum variance
	minVarX := int64(1 << 62)
	bestTx := 0
	for t := 0; t < dimX; t++ {
		var sumX, sumX2 int64
		for _, robot := range robots {
			px, vx := robot[0], robot[2]
			x := (px + vx*t) % dimX
			if x < 0 {
				x += dimX
			}
			sumX += int64(x)
			sumX2 += int64(x) * int64(x)
		}
		// Variance = E[X²] - E[X]² (scaled by n² to avoid division)
		variance := int64(n)*sumX2 - sumX*sumX
		if variance < minVarX {
			minVarX = variance
			bestTx = t
		}
	}

	// Find time ty (0 to dimY-1) where Y-coordinates have minimum variance
	minVarY := int64(1 << 62)
	bestTy := 0
	for t := 0; t < dimY; t++ {
		var sumY, sumY2 int64
		for _, robot := range robots {
			py, vy := robot[1], robot[3]
			y := (py + vy*t) % dimY
			if y < 0 {
				y += dimY
			}
			sumY += int64(y)
			sumY2 += int64(y) * int64(y)
		}
		variance := int64(n)*sumY2 - sumY*sumY
		if variance < minVarY {
			minVarY = variance
			bestTy = t
		}
	}

	// Chinese Remainder Theorem: find t where t ≡ bestTx (mod dimX) and t ≡ bestTy (mod dimY)
	// Since dimX=101 and dimY=103 are coprime, solution exists and is unique mod (dimX*dimY)
	// t = bestTx + dimX * k, find k such that bestTx + dimX*k ≡ bestTy (mod dimY)
	// dimX * k ≡ (bestTy - bestTx) (mod dimY)
	// k ≡ (bestTy - bestTx) * modInverse(dimX, dimY) (mod dimY)

	diff := bestTy - bestTx
	for diff < 0 {
		diff += dimY
	}

	// modInverse of dimX mod dimY using extended Euclidean algorithm
	inv := modInverse(dimX, dimY)
	k := (diff * inv) % dimY

	result := bestTx + dimX*k
	return uint(result)
}

func modInverse(a, m int) int {
	// Extended Euclidean algorithm
	g, x, _ := extGCD(a, m)
	if g != 1 {
		return -1 // No inverse exists
	}
	return (x%m + m) % m
}

func extGCD(a, b int) (g, x, y int) {
	if a == 0 {
		return b, 0, 1
	}
	g, x1, y1 := extGCD(b%a, a)
	return g, y1 - (b/a)*x1, x1
}
