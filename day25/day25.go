package day25

import (
	"strings"

	"github.com/pedantic79/aoc2020go/util"
)

const day uint = 25

func init() {
	if util.CheckDayAndPart(day, 1) {
		util.Results = append(util.Results, RunPart1)
	}

	// if util.CheckDayAndPart(day, 2) {
	// 	util.Results = append(util.Results, RunPart2)
	// }
}

func RunPart1() util.AoCResult {
	return util.Timer(day, 1, parse, part1)
}

// func RunPart2() util.AoCResult {
// 	return util.Timer(day, 2, parse, part2)
// }

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
