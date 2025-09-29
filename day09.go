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
		// Part 2: Move whole files from highest ID to lowest
		maxFileId := 0
		for _, block := range blocks {
			if block > maxFileId {
				maxFileId = block
			}
		}
		
		for currentFileId := maxFileId; currentFileId >= 1; currentFileId-- {
			// Find the file's position and size
			fileStart := -1
			fileSize := 0
			
			for i, block := range blocks {
				if block == currentFileId {
					if fileStart == -1 {
						fileStart = i
					}
					fileSize++
				} else if fileStart != -1 {
					break // Found end of file
				}
			}
			
			if fileStart == -1 || fileSize == 0 {
				continue
			}
			
			// Find leftmost contiguous free space that can fit this file
			moved := false
			for i := 0; i < fileStart && !moved; i++ {
				if blocks[i] == -1 {
					// Count contiguous free space starting at i
					freeSize := 0
					for j := i; j < len(blocks) && blocks[j] == -1; j++ {
						freeSize++
					}
					
					if freeSize >= fileSize {
						// Found suitable contiguous space - move the file
						for j := 0; j < fileSize; j++ {
							blocks[i+j] = currentFileId
							blocks[fileStart+j] = -1
						}
						moved = true
					}
					// Skip to end of this free space block
					i += freeSize - 1 // -1 because loop will increment
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
