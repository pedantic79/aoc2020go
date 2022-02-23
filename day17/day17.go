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

type Coord interface {
	neighors() []Coord
}

type Coord3 [3]int

func NewCoord3(p Point) Coord3 {
	var coord Coord3
	coord[0] = p.x
	coord[1] = p.y

	return coord
}

func (c3 Coord3) neighors() []Coord {
	// 27 = 3**3
	stack := make([]Coord, 0, 27)
	tempS := make([]Coord, 0, 27)

	stack = append(stack, c3)
	for i := 0; i < len(c3); i++ {
		for len(stack) > 0 {
			// pop stack contents
			v := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			for _, n := range []int{-1, 0, 1} {
				newV := v.(Coord3)
				newV[i] += n
				tempS = append(tempS, Coord3(newV))
			}
		}
		// copy contents to stack, and clear tempS without deallocating
		stack = append(stack, tempS...)
		tempS = tempS[:0]
	}

	return stack
}

type Coord4 [4]int

func NewCoord4(p Point) Coord4 {
	var coord Coord4
	coord[0] = p.x
	coord[1] = p.y

	return coord
}

func (c4 Coord4) neighors() []Coord {
	// 81 = 3**4
	stack := make([]Coord, 0, 81)
	tempS := make([]Coord, 0, 81)

	stack = append(stack, c4)
	for i := 0; i < len(c4); i++ {
		for len(stack) > 0 {
			// pop stack contents
			v := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			for _, n := range []int{-1, 0, 1} {
				newV := v.(Coord4)
				newV[i] += n
				tempS = append(tempS, Coord4(newV))
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

type state set.Set[Coord]

func (s *state) add(v Coord) {
	set.Add(*s, v)
}

func (s *state) contains(v Coord) bool {
	return set.Contains(*s, v)
}

func (coords state) countNeighbors() map[Coord]int {
	count := make(map[Coord]int)

	for coord := range coords {
		for _, neigh := range coord.neighors() {
			count[neigh]++
		}
	}

	return count
}

func (coords state) tick() state {
	newState := make(state)
	neighCount := coords.countNeighbors()

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

func part1(points []Point) int {
	c3s := make(state)

	for _, point := range points {
		c3s.add(NewCoord3(point))
	}

	for i := 0; i < 6; i++ {
		c3s = c3s.tick()
	}

	return len(c3s)
}

func part2(points []Point) int {
	c4s := make(state)

	for _, point := range points {
		c4s.add(NewCoord4(point))
	}

	for i := 0; i < 6; i++ {
		c4s = c4s.tick()
	}

	return len(c4s)
}
