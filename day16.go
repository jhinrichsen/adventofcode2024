package adventofcode2024

import (
	"errors"
	"fmt"
	"image"
)

// Day16 specific errors
var (
	ErrNoStartFound = errors.New("no start position 'S' found in maze")
	ErrNoEndFound   = errors.New("no end position 'E' found in maze")
)

// NoSolutionError provides detailed information when no path exists
type NoSolutionError struct {
	Start image.Point
	End   image.Point
}

func (e *NoSolutionError) Error() string {
	return fmt.Sprintf("no path found from start S(%d,%d) to end E(%d,%d)",
		e.Start.X, e.Start.Y, e.End.X, e.End.Y)
}

func (e *NoSolutionError) Is(target error) bool {
	_, ok := target.(*NoSolutionError)
	return ok
}

var ErrNoSolutionFound = &NoSolutionError{}

// Direction deltas: North=0, East=1, South=2, West=3
var day16DX = [4]int{0, 1, 0, -1}
var day16DY = [4]int{-1, 0, 1, 0}

const day16MaxCost = uint(1<<31 - 1)

func Day16(grid []byte, part1 bool) (uint, error) {
	// Find dimensions
	dimX := 0
	for i, b := range grid {
		if b == '\n' {
			dimX = i
			break
		}
	}
	stride := dimX + 1
	dimY := (len(grid) + 1) / stride

	// Find start and end
	startX, startY, endX, endY := -1, -1, -1, -1
	for i, b := range grid {
		if b == 'S' {
			startX, startY = i%stride, i/stride
		} else if b == 'E' {
			endX, endY = i%stride, i/stride
		}
	}
	if startX < 0 {
		return 0, ErrNoStartFound
	}
	if endX < 0 {
		return 0, ErrNoEndFound
	}

	// State encoding: (y * dimX + x) * 4 + dir
	numStates := dimX * dimY * 4
	cost := make([]uint, numStates)
	for i := range cost {
		cost[i] = day16MaxCost
	}

	// Bucket queue for Dijkstra (costs are 1 or 1000)
	// Use two levels: buckets[i] = states with cost in [i*1000, (i+1)*1000)
	// Within each bucket, use deque-like processing
	maxBuckets := (dimX + dimY) * 2 // rough upper bound on turns
	buckets := make([][]int, maxBuckets)
	for i := range buckets {
		buckets[i] = make([]int, 0, 64)
	}

	// Start facing East (dir=1)
	startState := (startY*dimX + startX) * 4 + 1
	cost[startState] = 0
	buckets[0] = append(buckets[0], startState)

	endPos := endY*dimX + endX
	var minCost uint = day16MaxCost
	currentBucket := 0

	for currentBucket < maxBuckets {
		if len(buckets[currentBucket]) == 0 {
			currentBucket++
			continue
		}

		// Pop from current bucket
		s := buckets[currentBucket][len(buckets[currentBucket])-1]
		buckets[currentBucket] = buckets[currentBucket][:len(buckets[currentBucket])-1]

		pos := s / 4
		dir := s % 4
		c := cost[s]

		// Check if outdated
		if c > cost[s] {
			continue
		}

		// Check if reached end
		if pos == endPos {
			if part1 {
				return c, nil
			}
			if c < minCost {
				minCost = c
			}
			continue
		}

		// Early termination for part2
		if c > minCost {
			continue
		}

		x, y := pos%dimX, pos/dimX

		// Move forward
		nx, ny := x+day16DX[dir], y+day16DY[dir]
		if nx >= 0 && nx < dimX && ny >= 0 && ny < dimY {
			if grid[ny*stride+nx] != '#' {
				newState := (ny*dimX+nx)*4 + dir
				newCost := c + 1
				if newCost < cost[newState] {
					cost[newState] = newCost
					bucket := int(newCost / 1000)
					if bucket < maxBuckets {
						buckets[bucket] = append(buckets[bucket], newState)
					}
				}
			}
		}

		// Turn left (dir-1 mod 4)
		leftDir := (dir + 3) % 4
		leftState := pos*4 + leftDir
		leftCost := c + 1000
		if leftCost < cost[leftState] {
			cost[leftState] = leftCost
			bucket := int(leftCost / 1000)
			if bucket < maxBuckets {
				buckets[bucket] = append(buckets[bucket], leftState)
			}
		}

		// Turn right (dir+1 mod 4)
		rightDir := (dir + 1) % 4
		rightState := pos*4 + rightDir
		rightCost := c + 1000
		if rightCost < cost[rightState] {
			cost[rightState] = rightCost
			bucket := int(rightCost / 1000)
			if bucket < maxBuckets {
				buckets[bucket] = append(buckets[bucket], rightState)
			}
		}
	}

	if minCost == day16MaxCost {
		return 0, &NoSolutionError{
			Start: image.Point{X: startX, Y: startY},
			End:   image.Point{X: endX, Y: endY},
		}
	}

	// Part 2: backtrack to find all optimal tiles
	onOptimal := make([]bool, numStates)
	optimalTiles := make([]bool, dimX*dimY)

	// Mark end states with minCost
	stack := make([]int, 0, 256)
	for dir := range 4 {
		s := endPos*4 + dir
		if cost[s] == minCost {
			stack = append(stack, s)
			onOptimal[s] = true
		}
	}

	for len(stack) > 0 {
		s := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		pos := s / 4
		dir := s % 4
		optimalTiles[pos] = true
		c := cost[s]

		x, y := pos%dimX, pos/dimX

		// Predecessor: moved from backward position
		backDir := (dir + 2) % 4
		bx, by := x+day16DX[backDir], y+day16DY[backDir]
		if bx >= 0 && bx < dimX && by >= 0 && by < dimY {
			if grid[by*stride+bx] != '#' {
				prevState := (by*dimX+bx)*4 + dir
				if cost[prevState] == c-1 && !onOptimal[prevState] {
					onOptimal[prevState] = true
					stack = append(stack, prevState)
				}
			}
		}

		// Predecessor: turned right to get here (was facing left)
		leftDir := (dir + 3) % 4
		leftState := pos*4 + leftDir
		if cost[leftState] == c-1000 && !onOptimal[leftState] {
			onOptimal[leftState] = true
			stack = append(stack, leftState)
		}

		// Predecessor: turned left to get here (was facing right)
		rightDir := (dir + 1) % 4
		rightState := pos*4 + rightDir
		if cost[rightState] == c-1000 && !onOptimal[rightState] {
			onOptimal[rightState] = true
			stack = append(stack, rightState)
		}
	}

	var count uint
	for _, v := range optimalTiles {
		if v {
			count++
		}
	}
	return count, nil
}

func Day16Part1(grid []byte) (uint, error) {
	return Day16(grid, true)
}

func Day16Part2(grid []byte) (uint, error) {
	return Day16(grid, false)
}
