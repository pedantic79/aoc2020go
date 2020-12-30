package day03

import (
	"strings"
	"time"

	"github.com/pedantic79/aoc2020go/util"
)

var day uint = 3
var fileName string = util.GenerateFileName(day)

func init() {
	if util.CheckDayAndPart(day, 1) {
		util.Results = append(util.Results, RunPart1)
	}

	if util.CheckDayAndPart(day, 2) {
		util.Results = append(util.Results, RunPart2)
	}
}

func RunPart1() util.AoCResult {
	input := util.ReadFile(fileName)

	start := time.Now()
	parsed := parse(input)
	parseTime := time.Since(start)

	start = time.Now()
	ans1 := part1(parsed)
	runTime := time.Since(start)

	return util.AoCResult{Day: day, Part: 1, ParseTime: parseTime, RunTime: runTime, Value: ans1}
}

func RunPart2() util.AoCResult {
	input := util.ReadFile(fileName)

	start := time.Now()
	parsed := parse(input)
	parseTime := time.Since(start)

	start = time.Now()
	ans2 := part2(parsed)
	runTime := time.Since(start)

	return util.AoCResult{Day: day, Part: 2, ParseTime: parseTime, RunTime: runTime, Value: ans2}
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
