package main

import (
	"fmt"

	_ "github.com/pedantic79/aoc2020go/startup"
	"github.com/pedantic79/aoc2020go/util"

	_ "github.com/pedantic79/aoc2020go/day01"
	_ "github.com/pedantic79/aoc2020go/day02"
	_ "github.com/pedantic79/aoc2020go/day03"
	_ "github.com/pedantic79/aoc2020go/day04"
	_ "github.com/pedantic79/aoc2020go/day05"
	_ "github.com/pedantic79/aoc2020go/day06"
	_ "github.com/pedantic79/aoc2020go/day07"
	_ "github.com/pedantic79/aoc2020go/day08"
	_ "github.com/pedantic79/aoc2020go/day09"
	_ "github.com/pedantic79/aoc2020go/day10"
	_ "github.com/pedantic79/aoc2020go/day11"
	_ "github.com/pedantic79/aoc2020go/day12"
	_ "github.com/pedantic79/aoc2020go/day13"
	_ "github.com/pedantic79/aoc2020go/day14"
	_ "github.com/pedantic79/aoc2020go/day15"
	_ "github.com/pedantic79/aoc2020go/day16"
	_ "github.com/pedantic79/aoc2020go/day17"
	_ "github.com/pedantic79/aoc2020go/day18"
	_ "github.com/pedantic79/aoc2020go/day19"
	_ "github.com/pedantic79/aoc2020go/day20"
)

func main() {
	fmt.Printf("🎄 Advent of Code 2020 - Day %02d\n\n", *util.Day)

	for _, result := range util.Results {
		fmt.Printf("%v", result())
	}
}
