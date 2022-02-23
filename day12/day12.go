package day12

import (
	"log"
	"strings"

	"github.com/pedantic79/aoc2020go/framework"
	"github.com/pedantic79/aoc2020go/util"
)

const day uint = 12

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

type Directions struct {
	dir byte
	num int
}

type Ship struct {
	dir   byte
	coord Point
}

type Point struct {
	x int
	y int
}

func (s *Ship) rotateLeft(amount int) {
	for ; amount > 0; amount -= 90 {
		switch s.dir {
		case 'E':
			s.dir = 'N'
		case 'N':
			s.dir = 'W'
		case 'W':
			s.dir = 'S'
		case 'S':
			s.dir = 'E'
		default:
			log.Fatalf("%v is not valid direction", s.dir)
		}
	}
}

func (s *Ship) rotateRight(amount int) {
	for ; amount > 0; amount -= 90 {
		switch s.dir {
		case 'E':
			s.dir = 'S'
		case 'S':
			s.dir = 'W'
		case 'W':
			s.dir = 'N'
		case 'N':
			s.dir = 'E'
		default:
			log.Fatalf("%v is not valid direction", s.dir)
		}
	}
}

func (s *Ship) forward(amount int) {
	switch s.dir {
	case 'E':
		s.coord.x += amount
	case 'W':
		s.coord.x -= amount
	case 'N':
		s.coord.y += amount
	case 'S':
		s.coord.y -= amount
	default:
		log.Fatalf("%v is not valid direction", s.dir)
	}
}

func (p *Point) rotateLeft(amount int) {
	switch amount {
	case 0:
	case 90:
		p.x, p.y = -p.y, p.x
	case 180:
		p.x, p.y = -p.x, -p.y
	case 270:
		p.x, p.y = p.y, -p.x
	default:
		log.Fatalf("%v is not a supported rotation amount", amount)
	}
}

func (p *Point) rotateRight(amount int) {
	switch amount {
	case 0:
	case 90:
		p.x, p.y = p.y, -p.x
	case 180:
		p.x, p.y = -p.x, -p.y
	case 270:
		p.x, p.y = -p.y, p.x
	default:
		log.Fatalf("%v is not a supported rotation amount", amount)
	}
}

func (p *Point) forward(waypoint Point, amount int) {
	p.x += waypoint.x * amount
	p.y += waypoint.y * amount
}

func parse(input string) []Directions {
	directions := []Directions{}

	for _, line := range strings.Split(input, "\n") {
		num := util.Atoi(line[1:])

		directions = append(directions, Directions{
			dir: line[0],
			num: num,
		})
	}

	return directions
}

func part1(directions []Directions) int {
	ship := Ship{dir: 'E', coord: Point{x: 0, y: 0}}

	for _, direction := range directions {
		switch direction.dir {
		case 'N':
			ship.coord.y += direction.num
		case 'S':
			ship.coord.y -= direction.num
		case 'E':
			ship.coord.x += direction.num
		case 'W':
			ship.coord.x -= direction.num
		case 'L':
			ship.rotateLeft(direction.num)
		case 'R':
			ship.rotateRight(direction.num)
		case 'F':
			ship.forward(direction.num)
		}
	}

	return util.IntAbs(ship.coord.x) + util.IntAbs(ship.coord.y)
}

func part2(directions []Directions) int {
	waypoint := Point{x: 10, y: 1}
	ship := Point{x: 0, y: 0}

	for _, direction := range directions {
		switch direction.dir {
		case 'N':
			waypoint.y += direction.num
		case 'S':
			waypoint.y -= direction.num
		case 'E':
			waypoint.x += direction.num
		case 'W':
			waypoint.x -= direction.num
		case 'L':
			waypoint.rotateLeft(direction.num)
		case 'R':
			waypoint.rotateRight(direction.num)
		case 'F':
			ship.forward(waypoint, direction.num)
		}
	}

	return util.IntAbs(ship.x) + util.IntAbs(ship.y)
}
