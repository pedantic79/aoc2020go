package day17

import (
	"strings"

	"github.com/pedantic79/aoc2020go/framework"
	"github.com/pedantic79/aoc2020go/util/set"
)

const day uint = 17

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

type Point struct {
	x, y int
}

type Neighborable interface {
	comparable
	neighborSize() int
	dim() int
	add(int, int) interface{}
}

type Coord3 [3]int

func NewCoord3(p Point) Coord3 {
	var coord Coord3
	coord[0] = p.x
	coord[1] = p.y

	return coord
}

func (c3 Coord3) neighborSize() int {
	return 27
}

func (c3 Coord3) dim() int {
	return 3
}

func (c3 Coord3) add(index, n int) interface{} {
	c3[index] += n
	return c3
}

type Coord4 [4]int

func NewCoord4(p Point) Coord4 {
	var coord Coord4
	coord[0] = p.x
	coord[1] = p.y

	return coord
}

func (c4 Coord4) neighborSize() int {
	return 81
}

func (c4 Coord4) dim() int {
	return 4
}

func (c4 Coord4) add(index, n int) interface{} {
	c4[index] += n
	return c4
}

func neighbors[N Neighborable](c N) []N {
	stack := make([]N, 0, c.neighborSize())
	tempS := make([]N, 0, c.neighborSize())

	stack = append(stack, c)
	for i := 0; i < c.dim(); i++ {
		for len(stack) > 0 {
			v := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			for _, n := range []int{-1, 0, 1} {
				tempS = append(tempS, v.add(i, n).(N))
			}
		}

		// copy contents to stack, and clear tempS without deallocating
		stack = append(stack, tempS...)
		tempS = tempS[:0]
	}

	return stack
}

func parse(input string) []Point {
	nums := []Point{}

	for y, row := range strings.Split(input, "\n") {
		for x, cell := range row {
			if cell == '#' {
				nums = append(nums, Point{x, y})
			}
		}
	}

	return nums
}

func countNeighbors[T Neighborable](coords set.Set[T]) map[T]int {
	count := make(map[T]int)

	for coord := range coords {
		for _, neigh := range neighbors(coord) {
			count[neigh]++
		}
	}

	return count
}

func tick[T Neighborable](coords set.Set[T]) set.Set[T] {
	newState := set.New[T]()
	neighCount := countNeighbors(coords)

	for coord, count := range neighCount {
		if set.Contains(coords, coord) {
			// we're including ourselves
			if count == 3 || count == 4 {
				set.Add(newState, coord)
			}
		} else if count == 3 {
			set.Add(newState, coord)
		}
	}

	return newState
}

func solve[T Neighborable](points []Point, new func(Point) T) int {
	cs := set.New[T]()

	for _, point := range points {
		set.Add(cs, new(point))
	}

	for i := 0; i < 6; i++ {
		cs = tick(cs)
	}

	return len(cs)
}

func part1(points []Point) int {
	return solve(points, NewCoord3)
}

func part2(points []Point) int {
	return solve(points, NewCoord4)
}
