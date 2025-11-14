package adventofcode2024

import (
	"regexp"
	"strconv"
	"strings"
)

type ClawMachine struct {
	ButtonA Point
	ButtonB Point
	Prize   Point
}

type Point struct {
	X, Y int
}

type Day13Puzzle struct {
	Machines []ClawMachine
}

func NewDay13(lines []string) Day13Puzzle {
	var machines []ClawMachine
	var current ClawMachine

	buttonRegex := regexp.MustCompile(`Button ([AB]): X\+(\d+), Y\+(\d+)`)
	prizeRegex := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if matches := buttonRegex.FindStringSubmatch(line); matches != nil {
			x, _ := strconv.Atoi(matches[2])
			y, _ := strconv.Atoi(matches[3])
			point := Point{X: x, Y: y}

			if matches[1] == "A" {
				current.ButtonA = point
			} else {
				current.ButtonB = point
			}
		} else if matches := prizeRegex.FindStringSubmatch(line); matches != nil {
			x, _ := strconv.Atoi(matches[1])
			y, _ := strconv.Atoi(matches[2])
			current.Prize = Point{X: x, Y: y}

			// Complete machine, add to list
			machines = append(machines, current)
			current = ClawMachine{} // Reset for next machine
		}
	}

	return Day13Puzzle{Machines: machines}
}

func Day13(puzzle Day13Puzzle, part1 bool) int {
	totalTokens := 0

	for _, machine := range puzzle.Machines {
		tokens := solveMachine(machine, part1)
		totalTokens += tokens
	}

	return totalTokens
}

func solveMachine(machine ClawMachine, part1 bool) int {
	// Part 2 default: Add 10000000000000 to prize coordinates
	prizeX := machine.Prize.X + 10000000000000
	prizeY := machine.Prize.Y + 10000000000000

	if part1 {
		prizeX = machine.Prize.X
		prizeY = machine.Prize.Y
	}

	// System of linear equations:
	// a * ButtonA.X + b * ButtonB.X = prizeX
	// a * ButtonA.Y + b * ButtonB.Y = prizeY
	eq1 := Eq{machine.ButtonA.X, machine.ButtonB.X, prizeX}
	eq2 := Eq{machine.ButtonA.Y, machine.ButtonB.Y, prizeY}

	a, b, hasIntSolution := Cramer(eq1, eq2)

	if !hasIntSolution || a < 0 || b < 0 {
		return 0
	}

	// Part 1: Check button press limit
	if part1 && (a > 100 || b > 100) {
		return 0
	}

	// Verify solution
	if a*machine.ButtonA.X+b*machine.ButtonB.X == prizeX &&
		a*machine.ButtonA.Y+b*machine.ButtonB.Y == prizeY {
		return a*3 + b*1 // A costs 3 tokens, B costs 1 token
	}

	return 0
}
