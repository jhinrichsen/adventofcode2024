package adventofcode2024

import (
	"regexp"
)

func Day03(program string, part1 bool) (sum uint) {
	var (
		atoi = func(s string) (n uint) {
			for i := range s {
				n = 10*n + uint(s[i]-'0')
			}
			return
		}
		pattern = `mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`
		enabled = true
	)
	if part1 {
		pattern = `mul\((\d{1,3}),(\d{1,3})\)`
	}

	re := regexp.MustCompile(pattern)
	gs := re.FindAllStringSubmatch(program, -1)
	for j := range gs {
		switch gs[j][0] {
		case "do()":
			enabled = true
		case "don't()":
			enabled = false
		default:
			if enabled {
				sum += atoi(gs[j][1]) * atoi(gs[j][2])
			}
		}
	}
	return
}
