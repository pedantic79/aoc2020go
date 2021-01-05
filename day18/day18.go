package day18

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/pedantic79/aoc2020go/util"
)

var day uint = 18
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

type token struct {
	isNum  bool
	symbol byte
	num    int64
}

func (t token) String() string {
	if t.isNum {
		return fmt.Sprintf("Num(%v)", t.num)
	}

	return fmt.Sprintf("Symbol(%c)", t.symbol)
}

func (t token) isOperator() bool {
	if !t.isNum {
		return t.symbol == '+' || t.symbol == '*'
	}

	return false
}

func (t token) isLeftBracket() bool {
	if !t.isNum {
		return t.symbol == '('
	}

	return false
}

func parse(input string) [][]token {
	tokens := [][]token{}

	for _, line := range strings.Split(input, "\n") {
		equation := []token{}
		for i := 0; i < len(line); i++ {
			r := line[i]

			if r == ' ' {
				continue
			} else if r >= '0' && r <= '9' {
				equation = append(equation, token{isNum: true, num: int64(r - '0'), symbol: r})
			} else {
				equation = append(equation, token{isNum: false, num: 0, symbol: r})
			}
		}

		tokens = append(tokens, equation)
	}
	return tokens
}

func shuntingYard(tokens []token, priority func(t token) int) []token {
	output := []token{}
	opStack := []token{}

	for _, tok := range tokens {
		if tok.isNum {
			output = append(output, tok)
		} else {
			switch tok.symbol {
			case '+':
				fallthrough
			case '*':
				for len(opStack) > 0 && opStack[len(opStack)-1].isOperator() && priority(opStack[len(opStack)-1]) >= priority(tok) {
					op := opStack[len(opStack)-1]
					opStack = opStack[:len(opStack)-1]
					output = append(output, op)
				}

				opStack = append(opStack, tok)
			case '(':
				opStack = append(opStack, tok)
			case ')':
				for len(opStack) > 0 && !opStack[len(opStack)-1].isLeftBracket() {
					op := opStack[len(opStack)-1]
					opStack = opStack[:len(opStack)-1]
					output = append(output, op)
				}

				opStack = opStack[:len(opStack)-1]
			default:
				log.Panicf("unknown token %v", tok)
			}

		}
	}

	for len(opStack) > 0 {
		op := opStack[len(opStack)-1]
		opStack = opStack[:len(opStack)-1]
		output = append(output, op)
	}

	return output
}

func evalRPN(tokens []token) int64 {
	stack := []int64{}
	for _, tok := range tokens {
		if tok.isNum {
			stack = append(stack, tok.num)
		} else {
			lastb := len(stack) - 1
			lasta := len(stack) - 2

			switch tok.symbol {
			case '+':
				stack[lasta] += stack[lastb]
				stack = stack[:len(stack)-1]
			case '*':
				stack[lasta] *= stack[lastb]
				stack = stack[:len(stack)-1]
			default:
				log.Panicf("invalid token %v", tok)
			}
		}

	}

	return stack[0]
}

func part1(tokens [][]token) int64 {
	total := int64(0)
	for _, equation := range tokens {
		total += evalRPN(shuntingYard(equation, func(t token) int { return 0 }))
	}
	return total
}

func part2(tokens [][]token) int64 {
	total := int64(0)
	for _, equation := range tokens {
		total += evalRPN(shuntingYard(equation, func(t token) int {
			if t.symbol == '*' {
				return 0
			}
			return 1
		}))
	}
	return total
}
