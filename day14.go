package adventofcode2024

func Day14(buf []byte, dimX, dimY int, seconds uint, part1 bool) uint {
	var sectors [4]uint
	halfX, halfY := dimX/2, dimY/2
	secondsInt := int(seconds)

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

		// Direct O(1) position calculation
		finalX := (px + vx*secondsInt) % dimX
		finalY := (py + vy*secondsInt) % dimY

		// Normalize to positive coordinates
		if finalX < 0 {
			finalX += dimX
		}
		if finalY < 0 {
			finalY += dimY
		}

		// Skip robots on middle axes (only for odd dimensions)
		if finalX == halfX || finalY == halfY {
			continue
		}

		// Quadrant mapping: 0=top-left, 1=top-right, 2=bottom-left, 3=bottom-right
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
