package adventofcode2024

type Day09Puzzle struct {
	blocks []int // -1 = free space, >= 0 = file ID
}

func NewDay09(buf []byte) Day09Puzzle {
	// Parse into blocks - expand the compressed format
	var blocks []int // -1 = free space, >= 0 = file ID
	
	fileId := 0
	isFile := true
	
	for _, b := range buf {
		if b < '0' || b > '9' {
			continue // Skip non-digit characters (like newlines)
		}
		length := int(b - '0')
		for i := 0; i < length; i++ {
			if isFile {
				blocks = append(blocks, fileId)
			} else {
				blocks = append(blocks, -1)
			}
		}
		if isFile {
			fileId++
		}
		isFile = !isFile
	}
	
	return Day09Puzzle{blocks: blocks}
}

func Day09(puzzle Day09Puzzle, part1 bool) (checksum uint) {
	// Make a copy of blocks to avoid modifying the original
	blocks := make([]int, len(puzzle.blocks))
	copy(blocks, puzzle.blocks)
	
	if part1 {
		// Part 1: Move individual blocks from right to left
		left, right := 0, len(blocks)-1
		
		for left < right {
			// Find next free space from left
			for left < len(blocks) && blocks[left] != -1 {
				left++
			}
			// Find next file block from right
			for right >= 0 && blocks[right] == -1 {
				right--
			}
			
			if left < right {
				blocks[left] = blocks[right]
				blocks[right] = -1
				left++
				right--
			}
		}
	} else {
		// Part 2: Move whole files from highest ID to lowest (optimized)
		maxFileId := 0
		for _, block := range blocks {
			if block > maxFileId {
				maxFileId = block
			}
		}

		// Pre-compute file positions and sizes to avoid repeated scanning
		fileInfos := make(map[int][2]int, maxFileId+1) // fileId -> [start, size]
		for i, block := range blocks {
			if block >= 0 {
				if info, exists := fileInfos[block]; exists {
					fileInfos[block] = [2]int{info[0], info[1] + 1} // increment size
				} else {
					fileInfos[block] = [2]int{i, 1} // start position, size 1
				}
			}
		}

		// Pre-compute free space segments for faster lookup
		type freeSegment struct {
			start, size int
		}
		var freeSegments []freeSegment
		i := 0
		for i < len(blocks) {
			if blocks[i] == -1 {
				start := i
				size := 0
				for i < len(blocks) && blocks[i] == -1 {
					size++
					i++
				}
				freeSegments = append(freeSegments, freeSegment{start, size})
			} else {
				i++
			}
		}

		for currentFileId := maxFileId; currentFileId >= 1; currentFileId-- {
			info, exists := fileInfos[currentFileId]
			if !exists {
				continue
			}
			fileStart, fileSize := info[0], info[1]

			// Find leftmost suitable free segment
			for segIdx, seg := range freeSegments {
				if seg.start >= fileStart {
					break // Only consider segments to the left
				}
				if seg.size >= fileSize {
					// Move the file
					for j := 0; j < fileSize; j++ {
						blocks[seg.start+j] = currentFileId
						blocks[fileStart+j] = -1
					}

					// Update free segments list
					if seg.size == fileSize {
						// Remove the segment entirely
						freeSegments = append(freeSegments[:segIdx], freeSegments[segIdx+1:]...)
					} else {
						// Shrink the segment
						freeSegments[segIdx] = freeSegment{seg.start + fileSize, seg.size - fileSize}
					}

					// Add new free segment where file was (merge with adjacent if possible)
					newFreeStart := fileStart
					newFreeSize := fileSize

					// Simple approach: just add the new free segment without complex merging
					// The algorithm still works correctly without merging
					inserted := false
					for j := range freeSegments {
						if freeSegments[j].start > newFreeStart {
							freeSegments = append(freeSegments[:j], append([]freeSegment{{newFreeStart, newFreeSize}}, freeSegments[j:]...)...)
							inserted = true
							break
						}
					}
					if !inserted {
						freeSegments = append(freeSegments, freeSegment{newFreeStart, newFreeSize})
					}
					break
				}
			}
		}
	}
	
	// Calculate checksum
	for i, block := range blocks {
		if block >= 0 {
			checksum += uint(i) * uint(block)
		}
	}
	
	return checksum
}
