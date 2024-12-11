package adventofcode2024

func Day07(lines []string) (sum uint) {
	for _, line := range lines {
		// parse test value

		var (
			want uint

			vals   = [12]uint{}
			idx    uint
			append = func(x uint) {
				vals[idx] = x
				idx++
			}
			j int
			x uint
		)
		for j = range line {
			if line[j] == ':' {
				j++
				j++
				break
			}
			want = 10*want + uint(line[j]-'0')
		}

		// parse equation

		for _, b := range line[j:] {
			if b == ' ' {
				append(x)
				x = 0
				continue
			}
			x = 10*x + uint(b-'0')
		}
		// append final digit (no trailing separator)
		append(x)

		combinations := 1 << (idx - 1) // 2^(number of operators)
		for i := range combinations {
			got := vals[0]
			// deduct operators from bits
			for j, val := range vals[1:idx] {
				mask := 1 << j
				// use RISC-V style because why not
				if i&mask == 0 {
					got += val // funct7=0000000
				} else {
					got *= val // funct7=0000001
				}
			}
			if want == got {
				sum += want
				// don't count all combinations, just one
				break
			}
		}
	}
	return
}
