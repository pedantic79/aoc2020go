package day15

import (
	"strings"

	"github.com/pedantic79/aoc2020go/framework"
	"github.com/pedantic79/aoc2020go/util"
)

const day uint = 15

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
