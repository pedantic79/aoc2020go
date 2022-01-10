package day23

import (
	"fmt"
	"strings"

	"github.com/pedantic79/aoc2020go/util"
)

const day uint = 23

func init() {
	if util.CheckDayAndPart(day, 1) {
		util.Results = append(util.Results, RunPart1)
	}

	if util.CheckDayAndPart(day, 2) {
		util.Results = append(util.Results, RunPart2)
	}
}

func RunPart1() util.AoCResult {
	return util.Timer(day, 1, parse, part1)
}

func RunPart2() util.AoCResult {
	return util.Timer(day, 2, parse, part2)
}

func parse(input string) []int {
	nums := []int{}

	for _, c := range input {
		nums = append(nums, int(c-'0'))
	}

	return nums
}

func solve(ring *[]int, current int, iterations int) {
	len := len(*ring) - 1

	for i := 0; i < iterations; i++ {
		threeCups := [3]int{0, 0, 0}
		threeCups[0] = (*ring)[current]
		threeCups[1] = (*ring)[threeCups[0]]
		threeCups[2] = (*ring)[threeCups[1]]

		dest := current - 1
		if dest == 0 {
			dest = len
		}

		for threeCups[0] == dest || threeCups[1] == dest || threeCups[2] == dest {
			dest -= 1
			if dest == 0 {
				dest = len
			}
		}

		nextCurrent := (*ring)[threeCups[2]]
		destNext := (*ring)[dest]
		(*ring)[current] = nextCurrent
		(*ring)[dest] = threeCups[0]
		(*ring)[threeCups[2]] = destNext

		current = nextCurrent
	}
}

func part1(inputs []int) string {
	ring := make([]int, len(inputs)+1)

	for i, j := 0, 1; j < len(inputs); {
		ring[inputs[i]] = inputs[j]
		i++
		j++
	}
	ring[inputs[len(inputs)-1]] = inputs[0]
	solve(&ring, inputs[0], 100)

	output := strings.Builder{}
	for pos := ring[1]; pos != 1; pos = ring[pos] {
		output.WriteString(fmt.Sprint(pos))
	}

	return output.String()
}

const LEN2 = 1_000_000

func part2(inputs []int) string {
	ring := make([]int, LEN2+1)
	for i, j := 0, 1; j < len(inputs); {
		ring[inputs[i]] = inputs[j]
		i++
		j++
	}

	ring[inputs[len(inputs)-1]] = 10
	for i := 10; i < LEN2+1; i++ {
		ring[i] = i + 1
	}
	ring[LEN2] = inputs[0]

	solve(&ring, inputs[0], 10_000_000)
	a := uint64(ring[1])
	b := uint64(ring[a])

	return fmt.Sprint(a * b)
}
