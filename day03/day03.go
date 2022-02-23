package day03

import (
	"strings"

	"github.com/pedantic79/aoc2020go/framework"
)

const day uint = 3

func init() {
	if framework.CheckDayAndPart(day, 1) {
		framework.Results = append(framework.Results, RunPart1)
	}

	if framework.CheckDayAndPart(day, 2) {
		framework.Results = append(framework.Results, RunPart2)
	}
}

func RunPart1() framework.AoCResult {
	return framework.Timer(day, 1, parse, part1)
}

func RunPart2() framework.AoCResult {
	return framework.Timer(day, 2, parse, part2)
}

func parse(input string) []string {
	return strings.Split(input, "\n")
}

func count(forest []string, down int, right int) int {
	width := len(forest[0])
	total := 0

	col := 0
	for row := 0; row < len(forest); row += down {
		if forest[row][col] == '#' {
			total += 1
		}

		col = (col + right) % width
	}

	return total
}

func part1(forest []string) int {
	return count(forest, 1, 3)
}

func part2(forest []string) int {
	total := 1

	for _, x := range [][2]int{{1, 1}, {1, 3}, {1, 5}, {1, 7}, {2, 1}} {
		total *= count(forest, x[0], x[1])
	}

	return total
}
