package day07

import (
	"strconv"
	"strings"
	"time"

	"github.com/pedantic79/aoc2020go/util"
)

var day uint = 7
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

type BagCount struct {
	color  string
	number int
}

func parse(input string) map[string][]BagCount {
	bags := make(map[string][]BagCount)

	for _, line := range strings.Split(input, "\n") {
		rule := strings.Split(line[:len(line)-1], " bags contain ")
		key := rule[0]

		if rule[1] == "no other bags" {
			bags[key] = []BagCount{}
			continue
		}

		for _, contents := range strings.Split(rule[1], ", ") {
			contents = strings.TrimSuffix(contents, " bags")
			contents = strings.TrimSuffix(contents, " bag")

			count, _ := strconv.Atoi(contents[:1])
			bags[key] = append(bags[key], BagCount{color: contents[2:], number: count})
		}
	}

	return bags
}

func containsGold(bags map[string][]BagCount, color string) bool {
	if color == "shiny gold" {
		return true
	}

	for _, subBag := range bags[color] {
		if containsGold(bags, subBag.color) {
			return true
		}
	}

	return false
}

func part1(bags map[string][]BagCount) int {
	total := 0

	for color := range bags {
		if containsGold(bags, color) {
			total++
		}
	}

	return total - 1
}

func bagsContained(bags map[string][]BagCount, color string) int {
	total := 1

	for _, subBag := range bags[color] {
		total += subBag.number * bagsContained(bags, subBag.color)
	}

	return total
}

func part2(bags map[string][]BagCount) int {
	return bagsContained(bags, "shiny gold") - 1
}
