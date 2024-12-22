package adventofcode2024

type fa struct {
	id     uint
	length uint8
	file   bool // true: file, false: empty
}

func Day09(buf []byte) (checksum uint) {
	var (
		fat  = [20000]fa{}
		last int

		// ignore trailing empty blocks
		trackback = func() {
			last--
			for !fat[last].file {
				last--
			}
		}

		// create an empty fa at i by moving everything after i right
		mkEmpty = func(i int, length uint8) {
			last++
			for j := last; j > i; j-- {
				fat[j] = fat[j-1]
			}
			fat[i].id = 0
			fat[i].length = length
			fat[i].file = false
		}

		toggle bool
	)

	for i := range buf {
		// alternating file/empty
		toggle = !toggle

		length := buf[i] - '0'
		if length == 0 {
			continue
		}

		// store into FAT

		if toggle {
			fat[last].id = uint(i / 2)
			fat[last].file = toggle
		}
		fat[last].length = uint8(length)
		last++
	}
	trackback()

	for i := 0; i <= last; i++ {
		if fat[i].file {
			continue
		}

		free := fat[i].length
		avail := fat[last].length

		if free == avail {
			fat[i] = fat[last]
			trackback()
		} else if free < avail {
			fat[i].id = fat[last].id
			fat[i].file = true
			fat[last].length -= free
		} else { // free > avail
			fat[i].id = fat[last].id
			fat[i].length = avail
			fat[i].file = true
			mkEmpty(i+1, free-avail)
			trackback()
		}
	}

	var position uint
	for i := 0; i <= last; i++ {
		for range fat[i].length {
			checksum += position * fat[i].id
			position++
		}
	}
	return
}
