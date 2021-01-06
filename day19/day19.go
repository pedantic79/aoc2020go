package day19

import (
	"fmt"
	"strings"
	"time"

	"github.com/pedantic79/aoc2020go/util"
)

var day uint = 19
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

type Rule struct {
	char     byte
	children [][]int
}

func (r Rule) String() string {
	if r.char == 0 {
		return fmt.Sprint(r.children)
	} else {
		return fmt.Sprint(r.char)
	}
}

type Input struct {
	rules    map[int]Rule
	messages []string
}

func parseRules(input string) map[int]Rule {
	rules := make(map[int]Rule)

	for _, line := range strings.Split(input, "\n") {
		kv := strings.Split(line, ": ")
		var rule Rule
		if kv[1][0] == '"' {
			rule.char = kv[1][1]
		} else {
			children := [][]int{}
			for _, child := range strings.Split(kv[1], " | ") {
				childRule := []int{}
				for _, individualChildRule := range strings.Split(child, " ") {
					childRule = append(childRule, util.Atoi(individualChildRule))
				}
				children = append(children, childRule)
			}
			rule.children = children
		}

		rules[util.Atoi(kv[0])] = rule
	}

	return rules
}

func matchChar(char byte, s string) (string, bool) {
	if len(s) > 0 && s[0] == char {
		return s[1:], true
	}
	return s, false
}

func (inp Input) matchRule(ruleno int, s string) []string {
	if inp.rules[ruleno].char > 0 {
		if newS, flag := matchChar(inp.rules[ruleno].char, s); flag {
			return []string{newS}
		}
		return []string{}
	}

	res := []string{}
	for _, subRule := range inp.rules[ruleno].children {
		queue := []string{s}
		for _, num := range subRule {
			newQueue := []string{}
			for _, q0 := range queue {
				newQueue = append(newQueue, inp.matchRule(num, q0)...)
			}
			queue = newQueue
		}

		res = append(res, queue...)
	}

	return res
}

func parse(input string) Input {
	groups := strings.Split(input, "\n\n")

	rules := parseRules(groups[0])
	messages := []string{}

	for _, line := range strings.Split(groups[1], "\n") {
		if len(line) == 0 {
			continue
		}
		messages = append(messages, line)
	}

	return Input{rules: rules, messages: messages}
}

func part1(input Input) int {
	total := 0
	for _, message := range input.messages {
		for _, s := range input.matchRule(0, message) {
			if len(s) == 0 {
				total += 1
			}
		}
	}

	return total
}

func part2(input Input) int {
	input.rules[8] = Rule{
		char:     0,
		children: [][]int{{42}, {42, 8}},
	}

	input.rules[11] = Rule{
		char:     0,
		children: [][]int{{42, 31}, {42, 11, 31}},
	}

	total := 0
	for _, message := range input.messages {
		for _, s := range input.matchRule(0, message) {
			if len(s) == 0 {
				total += 1
			}
		}
	}

	return total
}
