package day09

import (
	"sort"
	"strings"

	"github.com/pedantic79/aoc2020go/framework"
	"github.com/pedantic79/aoc2020go/util"
)

const day uint = 9

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

	for _, line := range strings.Split(input, "\n") {
		nums = append(nums, util.Atoi(line))
	}

	return nums
}

const PREAMBLE = 25

func add_contains(set []int, target int) bool {
	for _, x := range set {
		for _, y := range set {
			if x+y == target {
				return true
			}
		}
	}

	return false
}

func part1(nums []int) int {
	for i := PREAMBLE; i < len(nums); i++ {
		target := nums[i]

		if !add_contains(nums[(i-25):i], target) {
			return target
		}
	}

	return -1
}

func part2(nums []int) int {
	target := part1(nums)
	prefix := []int{}

	sum := int(0)
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		prefix = append(prefix, sum)
	}

	for i := 0; i < len(prefix); i++ {
		for j := i + 1; j < len(prefix) && prefix[j]-prefix[i] <= target; j++ {
			if prefix[j]-prefix[i] == target {
				ans := nums[i : j+1]
				sort.Slice(ans, func(m, n int) bool { return ans[m] < ans[n] })
				return ans[0] + ans[len(ans)-1]
			}
		}
	}

	return -1
}
