package day05

import (
	"sort"
	"strings"

	"github.com/pedantic79/aoc2020go/util"
)

const day uint = 5

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

func parse(input string) []string {
	return strings.Split(input, "\n")
}

func convert(seats []string) []int {
	ids := []int{}

	for _, seat := range seats {
		id := 0
		for i := 0; i < len(seat); i++ {
			if seat[i] == 'R' || seat[i] == 'B' {
				id = id*2 + 1
			} else {
				id = id * 2
			}
		}

		ids = append(ids, id)
	}

	sort.Ints(ids)
	return ids
}

func part1(seats []string) int {
	ids := convert(seats)

	return ids[len(ids)-1]
}

func part2(seats []string) int {
	ids := convert(seats)

	last := ids[0]
	for i := 1; i < len(ids); i++ {
		if ids[i]-last != 1 {
			return last + 1
		}
		last = ids[i]
	}

	return -1
}
