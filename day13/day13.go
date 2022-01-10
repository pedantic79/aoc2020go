package day13

import (
	"strings"

	"github.com/pedantic79/aoc2020go/util"
)

const day uint = 13

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
