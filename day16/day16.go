package day16

import (
	"strings"

	"github.com/pedantic79/aoc2020go/util"
	"github.com/pedantic79/aoc2020go/util/set"
)

const day uint = 16

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

type DoubleRange struct {
	dr [4]int64
}

func (dr DoubleRange) check(i int64) bool {
	return (dr.dr[0] <= i && i <= dr.dr[1]) || (dr.dr[2] <= i && i <= dr.dr[3])
}

type Info struct {
	rules  []DoubleRange
	yours  []int64
	nearby [][]int64
}

func (info Info) checkTicket(ticket []int64) (bool, int64) {
	for _, tick := range ticket {
		found := false
		for _, rule := range info.rules {
			if rule.check(tick) {
				found = true
				break
			}
		}

		if !found {
			return false, tick
		}
	}

	return true, 0
}

func (info Info) ticketPossibles(validTickets [][]int64) []set.Set[int] {
	possiblities := make([]set.Set[int], len(info.yours))

	for i := 0; i < len(possiblities); i++ {
		set := make(set.Set[int])

		for r, rule := range info.rules {
			valid := true
			for _, near := range validTickets {
				if !rule.check(near[i]) {
					valid = false
					break
				}
			}

			if valid {
				set.Add(r)
			}
		}

		possiblities[i] = set
	}

	return possiblities
}

func parseRule(line string) DoubleRange {
	var dr [4]int64
	colon := strings.Split(line, ": ")
	numGroups := strings.Split(colon[1], " or ")
	for i, num := range strings.Split(numGroups[0], "-") {
		dr[i] = util.ParseInt(num, 10, 64)
	}

	for i, num := range strings.Split(numGroups[1], "-") {
		dr[i+2] = util.ParseInt(num, 10, 64)
	}

	return DoubleRange{dr}
}

func parseTicket(line string) []int64 {
	ticket := []int64{}

	for _, num := range strings.Split(line, ",") {
		ticket = append(ticket, util.ParseInt(num, 10, 64))
	}

	return ticket
}

func parse(input string) Info {
	rules := []DoubleRange{}
	group := strings.Split(input, "\n\n")

	for _, ruleLine := range strings.Split(group[0], "\n") {
		rules = append(rules, parseRule(ruleLine))
	}

	yours := parseTicket(strings.Split(group[1], "\n")[1])
	nearby := [][]int64{}

	for _, ticketLine := range strings.Split(group[2], "\n")[1:] {
		nearby = append(nearby, parseTicket(ticketLine))
	}

	return Info{
		rules:  rules,
		yours:  yours,
		nearby: nearby,
	}
}

func part1(info Info) int64 {
	var total int64

	for _, near := range info.nearby {
		if valid, num := info.checkTicket(near); !valid {
			total += num
		}
	}

	return total
}

func findSingleton(poss []set.Set[int]) (int, int) {
	for ticketIndex, set := range poss {
		if len(set) == 1 {
			for ruleIndex := range set {
				return ticketIndex, ruleIndex
			}
		}
	}
	return -1, -1
}

func part2(info Info) int64 {
	validTickets := [][]int64{}
	for _, near := range info.nearby {
		if valid, _ := info.checkTicket(near); valid {
			validTickets = append(validTickets, near)
		}
	}

	ticketPossiblities := info.ticketPossibles(validTickets)
	rulesToField := make([]int, len(ticketPossiblities))

	for ticketIdx, ruleIdx := findSingleton(ticketPossiblities); ticketIdx >= 0; ticketIdx, ruleIdx = findSingleton(ticketPossiblities) {
		rulesToField[ruleIdx] = ticketIdx
		for _, s := range ticketPossiblities {
			delete(s, ruleIdx)
		}
	}

	total := int64(1)
	for _, v := range rulesToField[:6] {
		total *= info.yours[v]
	}

	return total

}
