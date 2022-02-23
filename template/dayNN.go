package dayNN

import (
	"github.com/pedantic79/aoc2020go/framework"
)

const day uint = 0

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
