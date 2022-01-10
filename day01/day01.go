package day01

import (
	"sort"
	"strings"

	"github.com/pedantic79/aoc2020go/util"
	"github.com/pedantic79/aoc2020go/util/set"
)

const day uint = 1

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
	split := strings.Split(input, "\n")

	var output []int
	for _, line := range split {
		output = append(output, util.Atoi(line))
	}

	return output
}

func part1(input []int) int {
	seen := make(set.Set[int])

	for _, i := range input {
		target := 2020 - i

		if seen.Contains(target) {
			return target * i
		}
		seen.Add(i)
	}

	return -1
}

func part2(input []int) int {
	sort.Ints(input)

	for i := 0; i < len(input); i++ {
		l := i + 1
		r := len(input) - 1

		for l < r {
			sum := input[i] + input[l] + input[r]

			if sum == 2020 {
				return input[i] * input[l] * input[r]
			} else if sum > 2020 {
				r--
			} else {
				l++
			}
		}
	}

	return -1
}
