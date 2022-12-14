package day02

import (
	"strings"

	"github.com/pedantic79/aoc2020go/framework"
	"github.com/pedantic79/aoc2020go/util"
)

const day uint = 2

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

type PasswordRule struct {
	bounds [2]int
	letter byte
	passwd string
}

func (pr PasswordRule) validatePart1() bool {
	count := 0

	for i := 0; i < len(pr.passwd); i++ {
		if pr.passwd[i] == pr.letter {
			count++
		}
	}

	return pr.bounds[0] <= count && count <= pr.bounds[1]
}

func (pr PasswordRule) validatePart2() bool {
	return (pr.passwd[pr.bounds[0]-1] == pr.letter) != (pr.passwd[pr.bounds[1]-1] == pr.letter)
}

func parse(input string) []PasswordRule {
	split := strings.Split(input, "\n")

	var output []PasswordRule
	for _, line := range split {
		fields := strings.Split(line, " ")
		dash := strings.Split(fields[0], "-")

		bounds := [2]int{util.Atoi(dash[0]), util.Atoi(dash[1])}
		letter := fields[1][0]

		output = append(output, PasswordRule{
			bounds: bounds,
			letter: letter,
			passwd: fields[2],
		})
	}

	return output
}

func part1(passwords []PasswordRule) int {
	total := 0

	for _, password := range passwords {
		if password.validatePart1() {
			total++
		}
	}

	return total
}

func part2(passwords []PasswordRule) int {
	total := 0

	for _, password := range passwords {
		if password.validatePart2() {
			total++
		}
	}

	return total
}
