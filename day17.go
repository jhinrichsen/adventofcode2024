package adventofcode2024

import (
	"fmt"
	"log"
	"strings"
)

func Day17(lines []string, part1 bool) (uint, uint, uint, string) {
	log.Println()
	// registers are not limited to 3 bits but can hold positive numbers of arbitrary length
	parse := func(s string) uint {
		var n, e uint = 0, 1
		for i := len(s) - 1; s[i] >= '0' && s[i] <= '9'; i-- {
			n += uint(s[i]-'0') * e
			e *= 10
		}
		return n
	}
	a, b, c := parse(lines[0]), parse(lines[1]), parse(lines[2])
	cmd := lines[4][9:]
	log.Println("command", cmd)

	var output strings.Builder
	for pc := 0; pc < len(cmd)-2; {
		opcode, operand := cmd[pc]-'0', cmd[pc+2]-'0'

		var combo uint
		switch operand {
		case 0:
			combo = 0
		case 1:
			combo = 1
		case 2:
			combo = 2
		case 3:
			combo = 3
		case 4:
			combo = a
		case 5:
			combo = b
		case 6:
			combo = c
		case 7:
			// NOP
		default:
			panic(fmt.Sprintf("illegal operand %d", operand))
		}
		log.Println("[ a =", a, "b =", b, "c =", c, "]", "opcode =", opcode, " operand =", operand, "combo =", combo, "output = ", output.String())

		switch opcode {
		case 0: // adv
			x := a / (1 << combo)
			log.Println("adv", 1<<combo)
			a = x
			pc += 4
		case 1: // bxl
			x := b ^ uint(operand)
			log.Println("bxl", x)
			b = x
			pc += 4
		case 2: // bst
			b = combo % 8
			log.Printf("b=%d\n", b)
			pc += 4
		case 3: // jnz
			if a == 0 {
				log.Println("jnz (a=0)")
				pc += 4
			} else {
				pc = int(operand) * 2
				log.Println("jnz", pc, "(", cmd, ")")
			}
		case 4: // bxc
			log.Println("bxc", b^c)
			b = b ^ c
			pc += 4
		case 5: // out
			b := byte((combo % 8) + '0')
			log.Println("out", string(b))
			if output.Len() > 0 {
				output.WriteByte(',')
			}
			output.WriteByte(b)
			pc += 4
		case 6: // bdv
			x := a / (1 << combo)
			log.Println("bdv", x)
			b = x
			pc += 4
		case 7: // cdv
			x := a / (1 << combo)
			log.Println("cdv", x)
			c = x
			pc += 4
		default:
			panic(fmt.Sprintf("illegal opcode %d", opcode))
		}
	}
	log.Println("[ a =", a, "b =", b, "c =", c, "output = ", output.String(), "]")
	return a, b, c, output.String()
}
