package adventofcode2023

import (
	"regexp"
)

func Day03(program string, part1 bool) (sum uint) {
	var (
		mul  = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
		atoi = func(s string) (n uint) {
			for i := range s {
				n = 10*n + uint(s[i]-'0')
			}
			return
		}
	)

	gs := mul.FindAllStringSubmatch(program, -1)
	if part1 {
		for j := range gs {
			sum += atoi(gs[j][1]) * atoi(gs[j][2])
		}
		return
	}

	var (
		do      = regexp.MustCompile(`do\(\)|don't\(\)`)
		enabled = true // start enabled
	)
	enableds := make([]bool, len(program))
	for i := range enableds {
		enableds[i] = enabled
	}
	idxs := do.FindAllStringSubmatchIndex(program, -1)
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

	idxs = mul.FindAllStringSubmatchIndex(program, -1)
	for j := range gs {
		start := idxs[j][0]
		if enableds[start] {
			sum += atoi(gs[j][1]) * atoi(gs[j][2])
		}
	}
	return
}
