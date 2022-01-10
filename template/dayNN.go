package dayNN

import (
	"github.com/pedantic79/aoc2020go/util"
)

const day uint = 0

func init() {
	if util.CheckDayAndPart(day, 1) {
		util.Results = append(util.Results, RunPart1)
	}

	if util.CheckDayAndPart(day, 2) {
		util.Results = append(util.Results, RunPart2)
	}
}

func RunPart1() util.AoCResult {
	return util.Timer(day, 1, parse, part1)
}

func RunPart2() util.AoCResult {
	return util.Timer(day, 2, parse, part2)
}

func parse(input string) []int {
	nums := []int{}

	return nums
}

func part1(nums []int) int {
	return -1
}

func part2(nums []int) int {
	return -1
}
