package day06

import (
	"strings"
	"time"

	"github.com/pedantic79/aoc2020go/util"
)

var day uint = 6
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

type CustomsGroup struct {
	answers map[byte]int
	size    int
}

func parse(input string) []CustomsGroup {
	customsGroup := []CustomsGroup{}
	groups := strings.Split(input, "\n\n")

	for _, group := range groups {
		answers := make(map[byte]int)
		persons := strings.Split(group, "\n")

		for _, person := range persons {
			for i := 0; i < len(person); i++ {
				answers[person[i]]++
			}
		}
		customsGroup = append(customsGroup, CustomsGroup{answers: answers, size: len(persons)})
	}

	return customsGroup
}

func part1(group []CustomsGroup) int {
	total := 0

	for _, set := range group {
		total += len(set.answers)
	}

	return total
}

func part2(group []CustomsGroup) int {
	total := 0

	for _, count := range group {
		for _, v := range count.answers {
			if v == count.size {
				total += 1
			}
		}
	}

	return total
}
