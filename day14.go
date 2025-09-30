package adventofcode2024

import (
	"fmt"
	"math"
)

type R4 struct { // RestRoom Redoubt Robot
	px, py int16
	vx, vy int16
}

type Day14Puzzle struct {
	dimX, dimY int16 // dimension of plane/ room/ space
	robots     []R4
}

func NewDay14(lines []string, dimX, dimY int) (Day14Puzzle, error) {
	p := Day14Puzzle{int16(dimX), int16(dimY), make([]R4, 0, len(lines))}

	for i, line := range lines {
		px, py, vx, vy, err := parseRobotLine(line)
		if err != nil {
			return p, fmt.Errorf("error parsing line %d: %v", i+1, err)
		}
		p.robots = append(p.robots, R4{px, py, vx, vy})
	}

	return p, nil
}

// parseRobotLine parses "p=px,py v=vx,vy" format without allocations.
// Returns error if any coordinate or velocity value is outside int16 range (-32768 to 32767).
func parseRobotLine(line string) (px, py, vx, vy int16, err error) {
	// Expected format: "p=px,py v=vx,vy"
	if len(line) < 9 || line[0] != 'p' || line[1] != '=' {
		return 0, 0, 0, 0, fmt.Errorf("invalid format: expected 'p=' at start")
	}

	i := 2 // skip "p="

	// Parse px
	pxInt, i, err := parseInt(line, i)
	if err != nil {
		return 0, 0, 0, 0, fmt.Errorf("parsing px: %v", err)
	}
	if pxInt < math.MinInt8 || pxInt > math.MaxInt8 {
		return 0, 0, 0, 0, fmt.Errorf("px value %d out of int8 range", pxInt)
	}
	px = int16(pxInt)

	// Expect comma
	if i >= len(line) || line[i] != ',' {
		return 0, 0, 0, 0, fmt.Errorf("expected ',' after px")
	}
	i++ // skip comma

	// Parse py
	pyInt, i, err := parseInt(line, i)
	if err != nil {
		return 0, 0, 0, 0, fmt.Errorf("parsing py: %v", err)
	}
	if pyInt < math.MinInt8 || pyInt > math.MaxInt8 {
		return 0, 0, 0, 0, fmt.Errorf("py value %d out of int8 range", pyInt)
	}
	py = int16(pyInt)

	// Expect " v="
	if i+3 >= len(line) || line[i] != ' ' || line[i+1] != 'v' || line[i+2] != '=' {
		return 0, 0, 0, 0, fmt.Errorf("expected ' v=' after py")
	}
	i += 3 // skip " v="

	// Parse vx
	vxInt, i, err := parseInt(line, i)
	if err != nil {
		return 0, 0, 0, 0, fmt.Errorf("parsing vx: %v", err)
	}
	if vxInt < math.MinInt8 || vxInt > math.MaxInt8 {
		return 0, 0, 0, 0, fmt.Errorf("vx value %d out of int8 range", vxInt)
	}
	vx = int16(vxInt)

	// Expect comma
	if i >= len(line) || line[i] != ',' {
		return 0, 0, 0, 0, fmt.Errorf("expected ',' after vx")
	}
	i++ // skip comma

	// Parse vy
	vyInt, i, err := parseInt(line, i)
	if err != nil {
		return 0, 0, 0, 0, fmt.Errorf("parsing vy: %v", err)
	}
	if vyInt < math.MinInt8 || vyInt > math.MaxInt8 {
		return 0, 0, 0, 0, fmt.Errorf("vy value %d out of int8 range", vyInt)
	}
	vy = int16(vyInt)

	// Should be at end of line
	if i != len(line) {
		return 0, 0, 0, 0, fmt.Errorf("unexpected characters after vy")
	}

	return px, py, vx, vy, nil
}

// parseInt parses an integer starting at position i, returns value and new position
func parseInt(s string, i int) (int, int, error) {
	if i >= len(s) {
		return 0, i, fmt.Errorf("unexpected end of string")
	}

	negative := false
	if s[i] == '-' {
		negative = true
		i++
	}

	if i >= len(s) || s[i] < '0' || s[i] > '9' {
		return 0, i, fmt.Errorf("expected digit")
	}

	value := 0
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		value = value*10 + int(s[i]-'0')
		i++
	}

	if negative {
		value = -value
	}

	return value, i, nil
}

func Day14(p Day14Puzzle, seconds uint, part1 bool) uint {
	// Direct calculation: final_pos = (initial_pos + velocity * time) mod dimension
	secondsInt := int16(seconds)

	for i := range p.robots {
		robot := &p.robots[i]
		robot.px = (robot.px + robot.vx*secondsInt) % p.dimX
		robot.py = (robot.py + robot.vy*secondsInt) % p.dimY
	}

	var sectors [4]uint // four quadrants

	halfX, halfY := p.dimX/2, p.dimY/2
	for _, r := range p.robots {
		// Skip robots exactly on middle axes (only possible with odd dimensions)
		if (r.px % p.dimX) == halfX || (r.py % p.dimY) == halfY {
			continue
		}

		// Direct quadrant mapping using sign and magnitude
		// qx: 0 for left half (px < halfX), 1 for right half (px >= halfX) 
		// qy: 0 for top half (py < halfY), 1 for bottom half (py >= halfY)
		
		// For modulo results: negative values are in the "upper" part of their respective axis
		qx := int16(0)
		if r.px >= halfX || (r.px < 0 && r.px >= -halfX) {
			qx = 1
		}
		
		qy := int16(0)
		if r.py >= halfY || (r.py < 0 && r.py >= -halfY) {
			qy = 1
		}
		
		sectors[qy*2+qx]++
	}
	return sectors[0] * sectors[1] * sectors[2] * sectors[3]
}
