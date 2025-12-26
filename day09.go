package adventofcode2024

type Day09Puzzle struct {
	buf []byte
}

func NewDay09(buf []byte) Day09Puzzle {
	return Day09Puzzle{buf: buf}
}

func Day09(puzzle Day09Puzzle, part1 bool) uint {
	if part1 {
		return day09Part1(puzzle.buf)
	}
	return day09Part2(puzzle.buf)
}

func day09Part1(buf []byte) uint {
	// Expand to blocks for part 1 (simpler algorithm)
	var blocks []int
	fileId := 0
	isFile := true
	for _, b := range buf {
		if b < '0' || b > '9' {
			continue
		}
		length := int(b - '0')
		val := -1
		if isFile {
			val = fileId
			fileId++
		}
		for range length {
			blocks = append(blocks, val)
		}
		isFile = !isFile
	}

	left, right := 0, len(blocks)-1
	for left < right {
		for left < len(blocks) && blocks[left] != -1 {
			left++
		}
		for right >= 0 && blocks[right] == -1 {
			right--
		}
		if left < right {
			blocks[left], blocks[right] = blocks[right], -1
			left++
			right--
		}
	}

	var checksum uint
	for i, block := range blocks {
		if block >= 0 {
			checksum += uint(i) * uint(block)
		}
	}
	return checksum
}

func day09Part2(buf []byte) uint {
	// Parse into segments - don't expand to individual blocks
	type segment struct {
		pos  int // absolute position
		size int
		id   int // -1 for free, >= 0 for file
	}

	var segments []segment
	pos := 0
	fileId := 0
	isFile := true

	for _, b := range buf {
		if b < '0' || b > '9' {
			continue
		}
		size := int(b - '0')
		if size > 0 {
			id := -1
			if isFile {
				id = fileId
			}
			segments = append(segments, segment{pos, size, id})
		}
		pos += size
		if isFile {
			fileId++
		}
		isFile = !isFile
	}

	// Build file index: fileId -> segment index
	fileIdx := make([]int, fileId)
	for i, seg := range segments {
		if seg.id >= 0 {
			fileIdx[seg.id] = i
		}
	}

	// Process files from highest to lowest
	for fid := fileId - 1; fid >= 0; fid-- {
		segIdx := fileIdx[fid]
		file := segments[segIdx]

		// Find leftmost free segment that fits
		for i := 0; i < segIdx; i++ {
			if segments[i].id == -1 && segments[i].size >= file.size {
				// Move file here
				freePos := segments[i].pos
				freeSize := segments[i].size

				// Update file position
				segments[segIdx].pos = freePos

				if freeSize == file.size {
					// Free segment fully consumed - mark as used
					segments[i].id = -2 // consumed marker
				} else {
					// Shrink free segment
					segments[i].pos += file.size
					segments[i].size -= file.size
				}
				break
			}
		}
	}

	// Calculate checksum: sum of pos*id for each block
	// For a file at position p with size s and id i:
	// checksum contribution = i * (p + p+1 + ... + p+s-1) = i * (s*p + s*(s-1)/2)
	var checksum uint
	for _, seg := range segments {
		if seg.id >= 0 {
			// Arithmetic sum: p + (p+1) + ... + (p+s-1) = s*p + s*(s-1)/2
			sum := uint(seg.size)*uint(seg.pos) + uint(seg.size)*uint(seg.size-1)/2
			checksum += uint(seg.id) * sum
		}
	}
	return checksum
}
