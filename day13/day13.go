package day13

import (
	"strings"

	"github.com/pedantic79/aoc2020go/framework"
	"github.com/pedantic79/aoc2020go/util"
)

const day uint = 13

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

type Data struct {
	start   int64
	busIDs  []int64
	offsets []int64
}

func parse(input string) Data {
	lines := strings.Split(input, "\n")

	start := util.ParseInt(lines[0], 10, 64)

	ids := []int64{}
	offsets := []int64{}

	for idx, id := range strings.Split(lines[1], ",") {
		if id == "x" {
			continue
		}

		n := util.ParseInt(id, 10, 64)

		ids = append(ids, n)
		offsets = append(offsets, n-int64(idx))
	}

	return Data{start: start, busIDs: ids, offsets: offsets}
}

func part1(data Data) int64 {
	for i := data.start; ; i++ {
		for _, id := range data.busIDs {
			if i%id == 0 {
				return (i - data.start) * id
			}
		}
	}
}

func part2(data Data) int64 {
	return util.ChineseRemainderTheorem(data.offsets, data.busIDs)
}
