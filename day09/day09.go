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

func parse(input string) []int64 {
	nums := []int64{}

	for _, line := range strings.Split(input, "\n") {
		num, _ := strconv.ParseInt(line, 10, 64)
		nums = append(nums, num)
	}

	return nums
}

const PREAMBLE = 25

func add_contains(set map[int64]struct{}, target int64) bool {
	for k := range set {
		if _, ok := set[target-k]; ok {
			return true
		}
	}

	return false
}

func part1(nums []int64) int64 {
	seen := make(map[int64]struct{})
	for i := 0; i < PREAMBLE; i++ {
		seen[nums[i]] = util.Empty
	}

	for i := PREAMBLE; i < len(nums); i++ {
		target := nums[i]

		if !add_contains(seen, target) {
			return target
		}

		delete(seen, nums[i-PREAMBLE])
		seen[target] = util.Empty
	}

	return -1
}

func part2(nums []int64) int64 {
	target := part1(nums)
	prefix := []int64{}

	sum := int64(0)
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
