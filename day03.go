package adventofcode2023

import (
	"regexp"
)

func Day03(lines []string, part1 bool) (sum uint) {
	var (
		mul  = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
		atoi = func(s string) (n uint) {
			for i := range s {
				n = 10*n + uint(s[i]-'0')
			}
			return
		}
	)

	gs := mul.FindAllStringSubmatch(lines[i], -1) // match groups
	if part1 {
		for i := range lines {
			gs := mul.FindAllStringSubmatch(lines[i], -1)
			for j := range gs {
				sum += atoi(gs[j][1]) * atoi(gs[j][2])
			}
			return
		}
	}

	var (
		do      = regexp.MustCompile(`do\(\)|don't\(\)`)
		enabled = true // start enabled
	)
	for i := range lines {
		enableds := make([]bool, len(lines[i]))
		for i := range enableds {
			enableds[i] = enabled
		}
		idxs := do.FindAllStringSubmatchIndex(lines[i], -1)
		for j := range idxs {
			start := idxs[j][0]
			end := idxs[j][1]
			enabled = true
			if end-start == 7 { // "don't()"
				enabled = false
			}
			for j := start; j < len(enableds); j++ {
				enableds[j] = enabled
			}
		}

		gs := mul.FindAllStringSubmatch(lines[i], -1)       // mul groups
		idxs = mul.FindAllStringSubmatchIndex(lines[i], -1) // and corresponding indices
		for j := range gs {
			start := idxs[j][0]
			if enableds[start] {
				sum += atoi(gs[j][1]) * atoi(gs[j][2])
			}
		}
	}
	return
}
