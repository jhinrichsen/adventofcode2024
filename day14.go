package adventofcode2024

import (
	"fmt"
	"regexp"
	"strconv"
)

type R4 struct { // RestRoom Redoubt Robot
	px, py int
	vx, vy int
}

type Day14Puzzle struct {
	dimX, dimY int // dimension of plane/ room/ space
	robots     []R4
}

func NewDay14(lines []string, dimX, dimY int) (Day14Puzzle, error) {
	p := Day14Puzzle{dimX, dimY, nil}
	re := regexp.MustCompile(`^p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)$`)

	for i, line := range lines {
		m := re.FindStringSubmatch(line)
		if m == nil {
			return p, fmt.Errorf("error parsing line %d: %s", i+1, line)
		}

		px, err := strconv.Atoi(m[1])
		if err != nil {
			return p, fmt.Errorf("error parsing X position of line %d: %v", i+1, err)
		}
		py, err := strconv.Atoi(m[2])
		if err != nil {
			return p, fmt.Errorf("error parsing Y position of line %d: %v", i+1, err)
		}
		vx, err := strconv.Atoi(m[3])
		if err != nil {
			return p, fmt.Errorf("error parsing X velocity of line %d: %v", i+1, err)
		}
		vy, err := strconv.Atoi(m[4])
		if err != nil {
			return p, fmt.Errorf("error parsing Y velocity of line %d: %v", i+1, err)
		}
		p.robots = append(p.robots, R4{px, py, vx, vy})
	}

	return p, nil
}

func Day14(p Day14Puzzle, seconds uint, part1 bool) uint {
	for range seconds {
		for i := range p.robots {
			p.robots[i].px = (p.robots[i].px + p.robots[i].vx) % p.dimX
			p.robots[i].py = (p.robots[i].py + p.robots[i].vy) % p.dimY
		}
	}

	var sectors [5]uint // neutral middle lane and four sectors

	// ignore robots on x and y middle axis, which can only happen on odd dimensions
	oddX, oddY := p.dimX%2 != 0, p.dimY%2 != 0
	mx, my := p.dimX/2, p.dimY/2
	for _, r := range p.robots {
		// normalize into positive coordinates
		if r.px < 0 {
			r.px += p.dimX
		}
		if r.py < 0 {
			r.py += p.dimY
		}

		// count sector occurences

		// 0 == neutral middle lane
		if oddX && r.px == mx || oddY && r.py == my {
			sectors[0]++
		} else if r.px <= mx && r.py <= my {
			sectors[1]++
		} else if r.px >= mx && r.py <= my {
			sectors[2]++
		} else if r.px <= mx && r.py >= my {
			sectors[3]++
		} else if r.px >= mx && r.py >= my {
			sectors[4]++
		} else {
			panic("bad sector")
		}
	}
	fmt.Printf("robots in sectors: %d, %d, %d, %d\n", sectors[1], sectors[2], sectors[3], sectors[4])
	return sectors[1] * sectors[2] * sectors[3] * sectors[4]
}
