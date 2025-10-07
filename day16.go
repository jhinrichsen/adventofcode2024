package adventofcode2024

import (
	"container/heap"
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

// Sentinel error for errors.Is() compatibility
var ErrNoSolutionFound = &NoSolutionError{}

type state struct {
	pos image.Point
	dir int // 0=North, 1=East, 2=South, 3=West
}

type dijkstraNode struct {
	state state
	cost  uint
	index int // for heap interface
}

type priorityQueue []*dijkstraNode

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	node := x.(*dijkstraNode)
	node.index = n
	*pq = append(*pq, node)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	old[n-1] = nil  // avoid memory leak
	node.index = -1 // for safety
	*pq = old[0 : n-1]
	return node
}

func Day16(grid []byte, part1 bool) (uint, error) {
	// Find dimensions by scanning for first newline
	dimX := 0
	for i, b := range grid {
		if b == '\n' {
			dimX = i
			break
		}
	}

	// Find start and end positions
	var start, end image.Point
	var startFound, endFound bool

	for i, b := range grid {
		if b == '\n' {
			continue
		}
		x := i % (dimX + 1)
		y := i / (dimX + 1)

		switch b {
		case 'S':
			start = image.Point{X: x, Y: y}
			startFound = true
		case 'E':
			end = image.Point{X: x, Y: y}
			endFound = true
		}
	}

	if !startFound {
		return 0, ErrNoStartFound
	}
	if !endFound {
		return 0, ErrNoEndFound
	}
	// Direction vectors: North, East, South, West
	dirs := []image.Point{
		{X: 0, Y: -1}, // North
		{X: 1, Y: 0},  // East
		{X: 0, Y: 1},  // South
		{X: -1, Y: 0}, // West
	}

	// Start facing East (direction 1)
	startState := state{pos: start, dir: 1}

	// Priority queue for Dijkstra's
	pq := &priorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &dijkstraNode{state: startState, cost: 0})

	// Track visited states with their minimum cost
	visited := make(map[state]uint)
	var minCost uint
	var foundSolution bool

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*dijkstraNode)

		// Skip if we've already found a better path to this state
		if prevCost, exists := visited[current.state]; exists && prevCost < current.cost {
			continue
		}
		visited[current.state] = current.cost

		// Check if we reached the end
		if current.state.pos == end {
			if part1 {
				return current.cost, nil
			}
			if !foundSolution {
				minCost = current.cost
				foundSolution = true
			} else if current.cost > minCost {
				break // No more optimal paths
			}
		}

		// Try moving forward
		newPos := current.state.pos.Add(dirs[current.state.dir])
		if newPos.X >= 0 && newPos.X < dimX && newPos.Y >= 0 {
			gridPos := newPos.Y*(dimX+1) + newPos.X
			if gridPos < len(grid) && grid[gridPos] != '#' {
				newState := state{pos: newPos, dir: current.state.dir}
				if prevCost, exists := visited[newState]; !exists || prevCost > current.cost+1 {
					heap.Push(pq, &dijkstraNode{state: newState, cost: current.cost + 1})
				}
			}
		}

		// Try turning left
		leftDir := (current.state.dir + 3) % 4 // -1 mod 4
		leftState := state{pos: current.state.pos, dir: leftDir}
		if prevCost, exists := visited[leftState]; !exists || prevCost > current.cost+1000 {
			heap.Push(pq, &dijkstraNode{state: leftState, cost: current.cost + 1000})
		}

		// Try turning right
		rightDir := (current.state.dir + 1) % 4
		rightState := state{pos: current.state.pos, dir: rightDir}
		if prevCost, exists := visited[rightState]; !exists || prevCost > current.cost+1000 {
			heap.Push(pq, &dijkstraNode{state: rightState, cost: current.cost + 1000})
		}
	}

	if !foundSolution {
		return 0, &NoSolutionError{Start: start, End: end}
	}

	// Part 2: Count tiles on optimal paths using backtracking
	optimalTiles := make(map[image.Point]bool)

	// Find all end states with minimum cost
	endStates := []state{}
	for s, cost := range visited {
		if s.pos == end && cost == minCost {
			endStates = append(endStates, s)
		}
	}

	// Backtrack from all optimal end states
	stack := make([]state, 0, len(endStates))
	stack = append(stack, endStates...)
	processed := make(map[state]bool)

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if processed[current] {
			continue
		}
		processed[current] = true
		optimalTiles[current.pos] = true

		currentCost := visited[current]

		// Check all possible predecessors
		// 1. Move backward (came from forward move)
		backDir := (current.dir + 2) % 4 // opposite direction
		backPos := current.pos.Add(dirs[backDir])
		if backPos.X >= 0 && backPos.X < dimX && backPos.Y >= 0 {
			gridPos := backPos.Y*(dimX+1) + backPos.X
			if gridPos < len(grid) && grid[gridPos] != '#' {
				prevState := state{pos: backPos, dir: current.dir}
				if cost, exists := visited[prevState]; exists && cost == currentCost-1 {
					stack = append(stack, prevState)
				}
			}
		}

		// 2. Turn left (came from right turn)
		leftDir := (current.dir + 1) % 4
		leftState := state{pos: current.pos, dir: leftDir}
		if cost, exists := visited[leftState]; exists && cost == currentCost-1000 {
			stack = append(stack, leftState)
		}

		// 3. Turn right (came from left turn)
		rightDir := (current.dir + 3) % 4
		rightState := state{pos: current.pos, dir: rightDir}
		if cost, exists := visited[rightState]; exists && cost == currentCost-1000 {
			stack = append(stack, rightState)
		}
	}

	return uint(len(optimalTiles)), nil
}

func Day16Part1(grid []byte) (uint, error) {
	return Day16(grid, true)
}

func Day16Part2(grid []byte) (uint, error) {
	return Day16(grid, false)
}
