package day09

import (
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/pedantic79/aoc2020go/util"
)

var day uint = 9
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

func parse(input string) []int {
	nums := []int{}

	for _, line := range strings.Split(input, "\n") {
		num, _ := strconv.Atoi(line)
		nums = append(nums, num)
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
