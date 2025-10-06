package adventofcode2024

import (
	"fmt"
	"strings"
)

func Day17(lines []string, part1 bool) (uint, uint, uint, string) {
	fmt.Println()
	// registers are not limited to 3 bits but can hold positive numbers of arbitrary length
	parse := func(s string) uint {
		var n, e uint = 0, 1
		for i := len(s) - 1 - 1; s[i] >= '0' && s[i] <= '9'; i-- {
			n += uint(s[i]-'0') * e
			e *= 10
		}
		return n
	}
	a, b, c := parse(lines[0]), parse(lines[1]), parse(lines[2])
	fmt.Println("[ a =", a, "b =", b, "c =", c, "]")
	cmd := lines[4][9:]
	fmt.Println("command", cmd)

	var output strings.Builder
	for pc := 0; pc < len(cmd)-2; {
		opcode, operand := cmd[pc]-'0', cmd[pc+2]-'0'
		fmt.Println("opcode =", opcode, " operand =", operand)

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
		fmt.Printf("combo=%d\n", combo)

		switch opcode {
		case 0: // adv
			fmt.Println("adv", 2<<combo)
			a = a / (2 << combo)
			pc += 4
		case 1: // bxl
			fmt.Println("todo bxl")
			pc += 4
		case 2: // bst
			b = combo % 8
			fmt.Printf("b=%d\n", b)
			pc += 4
		case 3: // jnz
			if a == 0 {
				fmt.Println("jnz (a=0)")
				pc += 4
			} else {
				pc = int(opcode) * 2
				fmt.Println("jnz", pc, "(", cmd, ")")
			}
		case 4: // bxc
			fmt.Println("bxc", b^c)
			b = b ^ c
			pc += 4
		case 5: // out
			b := byte((combo % 8) + '0')
			fmt.Println("out", string(b))
			if output.Len() > 0 {
				output.WriteByte(',')
			}
			output.WriteByte(b)
			pc += 4
		case 6: // bdv
			fmt.Println("todo bdv")
			pc += 4
		case 7: // cdv
			fmt.Println("todo cdv")
			pc += 4
		default:
			panic(fmt.Sprintf("illegal opcode %d", opcode))
		}
	}
	fmt.Println("[ a =", a, "b =", b, "c =", c, "output = ", output.String(), "]")
	return a, b, c, output.String()
}
