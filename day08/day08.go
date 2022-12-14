package day08

import (
	"log"
	"strings"

	"github.com/pedantic79/aoc2020go/framework"
	"github.com/pedantic79/aoc2020go/util"
	"github.com/pedantic79/aoc2020go/util/set"
)

const day uint = 8

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

type Instruction struct {
	ins string
	num int
}

func parse(input string) []Instruction {
	ins := []Instruction{}

	for _, line := range strings.Split(input, "\n") {
		num := util.Atoi(line[4:])
		ins = append(ins, Instruction{ins: line[:3], num: num})
	}

	return ins
}

func run(instructions []Instruction) (int, bool) {
	insCache := make(set.Set[int])
	acc := 0

	for pc := 0; pc < len(instructions); pc++ {
		// log.Printf("pc: %v acc: %v ins: %v cache: %v", pc, acc, instructions[pc], insCache)
		if set.Contains(insCache, pc) {
			return acc, false
		}

		set.Add(insCache, pc)
		ins := instructions[pc]

		switch ins.ins {
		case "nop":
		case "acc":
			acc += ins.num
		case "jmp":
			pc += ins.num - 1
			continue
		default:
			log.Fatalf("invalid instruction %v", ins)
		}
	}

	return acc, true
}

func part1(instructions []Instruction) int {
	ans, _ := run(instructions)
	return ans
}

func part2(instructions []Instruction) int {
	for i := 0; i < len(instructions); i++ {
		ins := instructions[i]

		switch ins.ins {
		case "acc":
			continue
		case "nop":
			instructions[i] = Instruction{
				ins: "jmp",
				num: ins.num,
			}
		case "jmp":
			instructions[i] = Instruction{
				ins: "nop",
				num: ins.num,
			}
		}

		if acc, terminated := run(instructions); terminated {
			return acc
		}

		instructions[i] = ins
	}

	return -1
}
