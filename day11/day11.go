package day11

import (
	"reflect"
	"strings"
	"time"

	"github.com/pedantic79/aoc2020go/util"
)

var day uint = 11
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

func copySlice(slice [][]byte) [][]byte {
	output := make([][]byte, 0, len(slice))

	for _, row := range slice {
		newRow := append([]byte(nil), row...)
		output = append(output, newRow)
	}

	return output
}

func parse(input string) [][]byte {
	rows := [][]byte{}

	for _, row := range strings.Split(input, "\n") {
		r := []byte(row)
		rows = append(rows, r)
	}

	return rows
}

func countNeighbors(floor [][]byte, r int, c int) int {
	count := 0

	for rDelta := -1; rDelta <= 1; rDelta++ {
		for cDelta := -1; cDelta <= 1; cDelta++ {
			if rDelta == 0 && cDelta == 0 {
				continue
			}

			rNew := r + rDelta
			cNew := c + cDelta

			if rNew >= 0 && rNew < len(floor) && cNew >= 0 && cNew < len(floor[0]) && floor[rNew][cNew] == '#' {
				count++
			}
		}
	}

	return count
}

func countNeighbors2(floor [][]byte, r int, c int) int {
	h, w := len(floor), len(floor[0])
	count := 0

	for _, delta := range [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}} {
		a, b := delta[0], delta[1]

		for y, x := r+a, c+b; y >= 0 && y < h && x >= 0 && x < w; y, x = y+a, x+b {
			if floor[y][x] == '#' {
				count++
			}

			if floor[y][x] != '.' {
				break
			}
		}
	}

	return count
}

func format(floor [][]byte) string {
	var sb strings.Builder

	for _, row := range floor {
		for _, cell := range row {
			sb.WriteByte(cell)
		}
		sb.WriteByte('\n')
	}

	return sb.String()
}

func part1(floor [][]byte) int {
	newFloor := copySlice(floor)

	for {
		for r := 0; r < len(floor); r++ {
			for c := 0; c < len(floor[0]); c++ {
				if floor[r][c] == '.' {
					continue
				}
				newFloor[r][c] = floor[r][c]

				count := countNeighbors(floor, r, c)
				switch floor[r][c] {
				case '#':
					if count >= 4 {
						newFloor[r][c] = 'L'
					}
				case 'L':
					if count == 0 {
						newFloor[r][c] = '#'
					}
				}
			}
		}

		if reflect.DeepEqual(floor, newFloor) {
			break
		}
		floor, newFloor = newFloor, floor
	}

	count := 0
	for r := 0; r < len(floor); r++ {
		for c := 0; c < len(floor[0]); c++ {
			if floor[r][c] == '#' {
				count++
			}
		}
	}

	return count
}

func part2(floor [][]byte) int {
	newFloor := copySlice(floor)

	for {
		for r := 0; r < len(floor); r++ {
			for c := 0; c < len(floor[0]); c++ {
				if floor[r][c] == '.' {
					continue
				}
				newFloor[r][c] = floor[r][c]

				count := countNeighbors2(floor, r, c)
				switch floor[r][c] {
				case '#':
					if count >= 5 {
						newFloor[r][c] = 'L'
					}
				case 'L':
					if count == 0 {
						newFloor[r][c] = '#'
					}
				}
			}
		}

		if reflect.DeepEqual(floor, newFloor) {
			break
		}
		floor, newFloor = newFloor, floor
	}

	count := 0
	for r := 0; r < len(floor); r++ {
		for c := 0; c < len(floor[0]); c++ {
			if floor[r][c] == '#' {
				count++
			}
		}
	}

	return count
}
