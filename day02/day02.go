package day02

import (
	"strings"
	"time"

	"github.com/pedantic79/aoc2020go/util"
)

var day uint = 2
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
