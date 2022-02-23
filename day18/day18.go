package day18

import (
	"fmt"
	"log"
	"strings"

	"github.com/pedantic79/aoc2020go/framework"
)

const day uint = 18

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
