package day20

import (
	"fmt"
	"strings"
	"time"

	"github.com/pedantic79/aoc2020go/util"
)

var day uint = 20
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

type Tile struct {
	data [10][10]bool
	id   int
}

func (t Tile) String() string {
	sb := strings.Builder{}
	for r := 0; r < len(t.data); r++ {
		for c := 0; c < len(t.data[r]); c++ {
			if t.data[r][c] {
				sb.WriteRune('#')
			} else {
				sb.WriteRune('.')
			}
		}
		sb.WriteRune('\n')
	}

	return sb.String()
}

func CalculateEdgeValue(edge []bool) (int, int) {
	forward := 0
	backward := 0

	for i := 0; i < len(edge); i++ {
		forward = forward*2 + util.Bool2Int(edge[i])
		backward = backward*2 + util.Bool2Int(edge[len(edge)-1-i])
	}

	return forward, backward
}

func (t Tile) edges() [8]int {
	t0, t1 := CalculateEdgeValue(t.data[0][:])
	b0, b1 := CalculateEdgeValue(t.data[9][:])

	left := [10]bool{}
	right := [10]bool{}
	for r := 0; r < len(t.data); r++ {
		left[r] = t.data[r][0]
		right[r] = t.data[r][9]
	}

	l0, l1 := CalculateEdgeValue(left[:])
	r0, r1 := CalculateEdgeValue(right[:])

	return [8]int{t0, r0, b0, l0, t1, r1, b1, l1}
}

type TileMap map[int]Tile

func parse(input string) TileMap {
	tiles := make(map[int]Tile)

	for _, group := range strings.Split(input, "\n\n") {
		lines := strings.Split(group, "\n")
		n := util.Atoi(lines[0][5:9])

		t := Tile{id: n}
		for r, row := range lines[1:] {
			for i := 0; i < len(row); i++ {
				if row[i] == '#' {
					t.data[r][i] = true
				}
			}
		}
		tiles[n] = t
	}

	return tiles
}

type TileCache map[int][]*Tile

func (tc TileCache) String() string {
	data := make(map[int][]int, len(tc))

	for k, v := range tc {
		entry := make([]int, len(v))
		for i := 0; i < len(v); i++ {
			entry[i] = v[i].id
		}
		data[k] = entry
	}

	return fmt.Sprint(data)
}

func calculateEdges(tiles TileMap) (TileCache, [4]int) {
	edge2Tile := make(map[int][]*Tile)
	for _, t := range tiles {
		t := t
		for _, id := range t.edges() {
			edge2Tile[id] = append(edge2Tile[id], &t)
		}
	}

	tileEdges := make(map[int]int)
	for n, t := range tiles {
		for _, id := range t.edges() {
			tileEdges[n] += len(edge2Tile[id])
		}
	}

	i := 0
	var res [4]int
	for tileID, count := range tileEdges {
		if count == 12 { // 2 sides with 2 entries, 2 sides with 4 entries
			res[i] = tileID
			i++
		}
	}

	return edge2Tile, res
}

func part1(tiles TileMap) int {
	prod := 1
	_, corners := calculateEdges(tiles)
	for _, n := range corners {
		prod *= n
	}

	return prod
}

func part2(tiles TileMap) int {
	cache, corners := calculateEdges(tiles)
	corner := tiles[corners[0]]
	edges := corner.edges()

	for _, edge := range edges[:4] {
		fmt.Printf("%010b -> len(%v)\n", edge, len(cache[edge]))
	}
	// Rotate corner until properly oriented
	// fill in corner
	// fill rest of row
	// fill in new row

	return -1
}
