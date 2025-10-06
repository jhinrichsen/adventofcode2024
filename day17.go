package adventofcode2024

import (
	"fmt"
)

func Day17(lines []string, part1 bool) (uint, uint, uint, string) {
	const (
		REGISTER_A = 4 // index into registers
		REGISTER_B = 5
		REGISTER_C = 6
	)
	// registers are not limited to 3 bits but can hold positive numbers of arbitrary length
	registers := [8]uint{0, 1, 2, 3, 0, 0, 0, 0}

	parse := func(s string, register uint8) {
		var e uint = 1
		for i := len(s) - 1; s[i] >= '0' && s[i] <= '9'; i-- {
			registers[register] += uint(s[i]-'0') * e
			e *= 10
		}
	}
	parse(lines[0], REGISTER_A)
	parse(lines[1], REGISTER_B)
	parse(lines[2], REGISTER_C)
	// empty line separator
	cmd := lines[4][9:]

	var output [1024]byte
	idx := 0
	l := len(cmd) - 2
	for pc := 0; pc < l; {
		opcode, operand := cmd[pc]-'0', cmd[pc+2]-'0'
		combo := registers[operand]

		switch opcode {
		case 0: // adv
			x := registers[REGISTER_A] / (1 << combo)
			registers[REGISTER_A] = x
			pc += 4
		case 1: // bxl
			x := registers[REGISTER_B] ^ uint(operand)
			registers[REGISTER_B] = x
			pc += 4
		case 2: // bst
			registers[REGISTER_B] = combo % 8
			pc += 4
		case 3: // jnz
			if registers[REGISTER_A] == 0 {
				pc += 4
			} else {
				pc = int(operand) * 2
			}
		case 4: // bxc
			registers[REGISTER_B] ^= registers[REGISTER_C]
			pc += 4
		case 5: // out
			if idx > 0 {
				output[idx] = ','
				idx++
			}
			output[idx] = byte((combo % 8) + '0')
			idx++
			pc += 4
		case 6: // bdv
			x := registers[REGISTER_A] / (1 << combo)
			registers[REGISTER_B] = x
			pc += 4
		case 7: // cdv
			x := registers[REGISTER_A] / (1 << combo)
			registers[REGISTER_C] = x
			pc += 4
		default:
			panic(fmt.Sprintf("illegal opcode %d", opcode))
		}
	}
	return registers[REGISTER_A], registers[REGISTER_B], registers[REGISTER_C], string(output[:idx])
}
