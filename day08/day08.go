package day08

import (
	"log"
	"strings"
	"time"

	"github.com/pedantic79/aoc2020go/util"
)

var day uint = 8
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
	insCache := make(util.Set)
	acc := 0

	for pc := 0; pc < len(instructions); pc++ {
		// log.Printf("pc: %v acc: %v ins: %v cache: %v", pc, acc, instructions[pc], insCache)
		if insCache.Contains(pc) {
			return acc, false
		}

		insCache.Add(pc)
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
