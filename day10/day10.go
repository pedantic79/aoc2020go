package day10

import (
	"log"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/pedantic79/aoc2020go/util"
)

var day uint = 10
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
	nums := []int{0}
	for _, line := range strings.Split(input, "\n") {
		num, _ := strconv.Atoi(line)
		nums = append(nums, num)
	}

	sort.Ints(nums)
	nums = append(nums, nums[len(nums)-1]+3)

	return nums
}

func part1(nums []int) int {
	diff := [2]int{}

	for i := 1; i < len(nums); i++ {
		switch nums[i] - nums[i-1] {
		case 1:
			diff[0]++
		case 3:
			diff[1]++
		default:
			log.Fatalf("difference between %v and %v is invalid", nums[i], nums[i-1])
		}
	}

	return diff[0] * diff[1]
}

func part2(nums []int) uint64 {
	dp := make([]uint64, len(nums))
	dp[0] = 1

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums) && nums[j]-nums[i] <= 3; j++ {
			dp[j] += dp[i]
		}
	}

	return dp[len(dp)-1]
}
