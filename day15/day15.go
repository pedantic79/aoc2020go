package day15

import (
	"strings"
	"time"

	"github.com/pedantic79/aoc2020go/util"
)

var day uint = 15
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
	for _, n := range strings.Split(input, ",") {
		num := util.Atoi(n)
		nums = append(nums, num)
	}

	return nums
}

// func solve(nums []int, limit int) int {
// 	seen := make(map[int]int)

// 	last := nums[0]
// 	for i := 1; i < limit; i++ {
// 		log.Println(i, last)
// 		if i < len(nums) {
// 			seen[last] = i
// 			last = nums[i]
// 		} else if old, ok := seen[last]; ok {
// 			seen[last] = i
// 			last = i - old
// 		} else {
// 			seen[last] = i
// 			last = 0
// 		}
// 	}

// 	return last
// }

func solve(nums []int, limit int) int {
	seen := make([]int, limit)

	for i, val := range nums {
		seen[val] = i + 1
	}

	last := nums[len(nums)-1]
	for i := len(nums); i < limit; i++ {
		// log.Println(i, last)
		if seen[last] > 0 {
			val := seen[last]
			seen[last] = i
			last = i - val
		} else {
			seen[last] = i
			last = 0
		}
	}

	return last
}

func part1(nums []int) int {
	return solve(nums, 2020)
}

func part2(nums []int) int {
	return solve(nums, 30000000)
}
