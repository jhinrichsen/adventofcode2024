package adventofcode2024

func Day17(lines []string, part1 bool) (uint, uint, uint, string) {
	run := func(a, b, c uint, cmd string) (uint, uint, uint, string) {
		const (
			REGISTER_A = 4 // index into registers
			REGISTER_B = 5
			REGISTER_C = 6
		)
		// registers are not limited to 3 bits but can hold positive numbers of arbitrary length
		registers := [8]uint{0, 1, 2, 3, a, b, c, 0}

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
				panic([]any{"illegal opcode", opcode})
			}
		}
		return registers[REGISTER_A], registers[REGISTER_B], registers[REGISTER_C], string(output[:idx])
	}

	parse := func(s string) uint {
		var (
			e uint = 1
			n uint
		)
		for i := len(s) - 1; s[i] >= '0' && s[i] <= '9'; i-- {
			n += uint(s[i]-'0') * e
			e *= 10
		}
		return n
	}
	a := parse(lines[0])
	b := parse(lines[1])
	c := parse(lines[2])
	// empty line separator
	cmd := lines[4][9:]

	if part1 {
		return run(a, b, c, cmd)
	}

	// fast parser: "d,d,d,..." -> []uint8 (no minus support needed)
	asNumbers := func(s string) []uint8 {
		if len(s) == 0 {
			return nil
		}
		out := make([]uint8, 0, 64)
		n := 0
		have := false
		for i := 0; i <= len(s); i++ {
			if i < len(s) {
				c := s[i]
				if c >= '0' && c <= '9' {
					n = n*10 + int(c-'0')
					have = true
					continue
				}
				if c == ',' || c == ' ' {
					// fall through to flush if we have a number
				} else {
					// ignore any other char (shouldn't happen)
					continue
				}
			}
			if have {
				out = append(out, uint8(n))
				n, have = 0, false
			}
		}
		return out
	}

	want := asNumbers(cmd)

	// we’ll grow candidates from the back: each step fixes one more trailing symbol
	cands := []uint{0}

	for i := len(want) - 1; i >= 0; i-- {
		need := want[i:] // required suffix
		next := make([]uint, 0, len(cands)*8)

		for _, hi := range cands {
			for d := uint(0); d < 8; d++ {
				a := (hi << 3) | d
				_, _, _, output := run(a, b, c, cmd)
				got := asNumbers(output)
				if len(got) < len(need) {
					continue
				}
				ok := true
				off := len(got) - len(need)
				for k := range need {
					if got[off+k] != need[k] {
						ok = false
						break
					}
				}
				if ok {
					next = append(next, a)
				}
			}
		}

		// must have at least one surviving candidate
		if len(next) == 0 {
			panic("no candidate keeps the required suffix")
		}
		cands = next
	}

	// pick any candidate that fully matches (usually the first)
	for _, a := range cands {
		_, _, _, output := run(a, b, c, cmd)
		if cmd == output {
			return a, b, c, output
		}
	}
	// Shouldn’t happen if the loop logic is correct
	panic("no candidate yields exact match")
}
