package day25

import (
	"strings"
	"time"

	"github.com/pedantic79/aoc2020go/util"
)

var day uint = 25
var fileName string = util.GenerateFileName(day)

func init() {
	if util.CheckDayAndPart(day, 1) {
		util.Results = append(util.Results, RunPart1)
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

func parse(input string) [2]uint64 {
	lines := strings.Split(input, "\n")
	a := util.ParseUint(lines[0], 10, 64)
	b := util.ParseUint(lines[1], 10, 64)

	return [2]uint64{a, b}
}

const SUBJECT uint64 = 7
const MOD = 20201227

func loopSize(target uint64) int {
	value, count := uint64(1), 0

	for ; target != value; count++ {
		value = value * SUBJECT % MOD
	}

	return count
}

func encryptionKey(key uint64, loopCount int) uint64 {
	value := uint64(1)

	for count := 0; count < loopCount; count++ {
		value = value * key % MOD
	}

	return value
}

func part1(input [2]uint64) uint64 {
	ls := loopSize(input[1])
	return encryptionKey(input[0], ls)

}
