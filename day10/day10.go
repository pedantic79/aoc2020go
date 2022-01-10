package day10

import (
	"log"
	"sort"
	"strings"

	"github.com/pedantic79/aoc2020go/util"
)

const day uint = 10

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
	nums := []int{0}
	for _, line := range strings.Split(input, "\n") {
		nums = append(nums, util.Atoi(line))
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
