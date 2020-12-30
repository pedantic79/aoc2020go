package day05

import (
	"sort"
	"strings"
	"time"

	"github.com/pedantic79/aoc2020go/util"
)

var day uint = 5
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

func convert(seats []string) []int {
	ids := []int{}

	for _, seat := range seats {
		id := 0
		for i := 0; i < len(seat); i++ {
			if seat[i] == 'R' || seat[i] == 'B' {
				id = id*2 + 1
			} else {
				id = id * 2
			}
		}

		ids = append(ids, id)
	}

	sort.Ints(ids)
	return ids
}

func part1(seats []string) int {
	ids := convert(seats)

	return ids[len(ids)-1]
}

func part2(seats []string) int {
	ids := convert(seats)

	last := ids[0]
	for i := 1; i < len(ids); i++ {
		if ids[i]-last != 1 {
			return last + 1
		}
		last = ids[i]
	}

	return -1
}
