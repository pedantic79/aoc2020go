package day20

import (
	"fmt"
	"strings"

	"github.com/pedantic79/aoc2020go/framework"
	"github.com/pedantic79/aoc2020go/util"
	"github.com/pedantic79/aoc2020go/util/set"
)

const day uint = 20

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

func (t *Tile) Rotate() {
	l := len(t.data)
	for i := 0; i < l/2; i++ {
		for j := i; j < l-i-1; j++ {
			temp := t.data[i][j]
			t.data[i][j] = t.data[l-1-j][i]
			t.data[l-1-j][i] = t.data[l-1-i][l-1-j]
			t.data[l-1-i][l-1-j] = t.data[j][l-1-i]
			t.data[j][l-1-i] = temp
		}
	}
}

func (t Tile) Clone() Tile {
	ret := Tile{id: t.id}
	for r := 0; r < 10; r++ {
		copy(ret.data[r][:], t.data[r][:])
	}
	return ret
}

func reverse(numbers *[10]bool) {
	for i := 0; i < len(*numbers)/2; i++ {
		j := len(*numbers) - i - 1
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
}

func (t *Tile) Flip() {
	for r := range t.data {
		reverse(&t.data[r])
	}
}

func CalculateEdgeValue(edge []bool) (int, int) {
	forward := 0
	backward := 0

	for i := 0; i < len(edge); i++ {
		forward = forward*2 + util.Bool2Int[int](edge[i])
		backward = backward*2 + util.Bool2Int[int](edge[len(edge)-1-i])
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

type Grid struct {
	grid    [20][20]*Tile
	visited set.Set[int]
}

func (g *Grid) search(r, c int, cache []Tile) {

}

func part2(tiles TileMap) int {
	cache := []Tile{}
	for _, tile := range tiles {
		for j := 0; j < 2; j++ {
			for i := 0; i < 4; i++ {
				tile.Rotate()
				cache = append(cache, tile.Clone())
			}
			tile.Flip()
		}
	}

	grid := Grid{visited: make(set.Set[int])}
	grid.search(0, 0, cache)

	return -1
}
