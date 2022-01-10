package day24

import (
	"fmt"
	"strings"

	"github.com/pedantic79/aoc2020go/util"
	"github.com/pedantic79/aoc2020go/util/set"
)

const day uint = 24

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

type HexCoord struct {
	x, y int
}

func (hc HexCoord) generateNeighbors() [6]HexCoord {
	neigh := [6][2]int{{1, 0}, {1, -1}, {0, -1}, {-1, 0}, {-1, 1}, {0, 1}}
	res := [6]HexCoord{}
	for i, n := range neigh {
		res[i] = HexCoord{hc.x + n[0], hc.y + n[1]}
	}

	return res
}

func newHexCoord(input string) HexCoord {
	x, y := 0, 0

	for i := 0; i < len(input); i++ {
		switch input[i] {
		case 'e':
			x += 1
		case 'w':
			x -= 1
		case 'n':
			i++
			switch input[i] {
			case 'e':
				y += 1
			case 'w':
				y += 1
				x -= 1
			default:
				panic(fmt.Sprintf("unknown character %v after 'n'", input[i]))
			}
		case 's':
			i++
			switch input[i] {
			case 'e':
				y -= 1
				x += 1
			case 'w':
				y -= 1
			default:
				panic(fmt.Sprintf("unknown character %v after 's'", input[i]))
			}
		default:
			panic(fmt.Sprintf("unknown character %v", input[i]))
		}
	}
	return HexCoord{x, y}
}

func parse(input string) set.Set[HexCoord] {
	hex := make(map[HexCoord]int)
	for _, line := range strings.Split(input, "\n") {
		hex[newHexCoord(line)]++
	}

	keys := make(set.Set[HexCoord])
	for coord, count := range hex {
		if count%2 == 1 {
			keys.Add(coord)

		}
	}

	return keys
}

func part1(hex set.Set[HexCoord]) int {
	return len(hex)
}

func tick(hex set.Set[HexCoord]) set.Set[HexCoord] {
	counts := make(map[HexCoord]int, len(hex)*6)

	for tile := range hex {
		for _, neighbors := range tile.generateNeighbors() {
			counts[neighbors]++
		}
	}

	keys := make(set.Set[HexCoord])
	for coord, count := range counts {
		_, contains := hex[coord]
		if contains && !(count == 0 || count > 2) || !contains && count == 2 {
			keys.Add(coord)
		}
	}

	return keys
}

func part2(hex set.Set[HexCoord]) int {
	for i := 0; i < 100; i++ {
		hex = tick(hex)
	}

	return len(hex)
}
