package day06

import (
	"strings"

	"github.com/pedantic79/aoc2020go/framework"
	"github.com/pedantic79/aoc2020go/util"
)

const day uint = 6

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

type CustomsGroup struct {
	answers map[byte]int
	size    int
}

func parse(input string) []CustomsGroup {
	customsGroup := []CustomsGroup{}
	groups := strings.Split(input, "\n\n")

	for _, group := range groups {
		answers := make(map[byte]int)
		persons := strings.Split(group, "\n")

		for _, person := range persons {
			for i := 0; i < len(person); i++ {
				answers[person[i]]++
			}
		}
		customsGroup = append(customsGroup, CustomsGroup{answers: answers, size: len(persons)})
	}

	return customsGroup
}

func part1(group []CustomsGroup) int {
	return util.SumFunc(group, func(s CustomsGroup) int { return len(s.answers) })
}

func part2(group []CustomsGroup) int {
	total := 0

	for _, count := range group {
		for _, v := range count.answers {
			if v == count.size {
				total += 1
			}
		}
	}

	return total
}
