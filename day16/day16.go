package day16

import (
	"strconv"
	"strings"
	"time"

	"github.com/pedantic79/aoc2020go/util"
)

var day uint = 16
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

func (info Info) ticketPossibles(validTickets [][]int64) []map[int]struct{} {
	possiblities := make([]map[int]struct{}, len(info.yours))

	for i := 0; i < len(possiblities); i++ {
		set := make(map[int]struct{})

		for r, rule := range info.rules {
			valid := true
			for _, near := range validTickets {
				if !rule.check(near[i]) {
					valid = false
					break
				}
			}

			if valid {
				set[r] = util.Empty
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
		n, _ := strconv.ParseInt(num, 10, 64)
		dr[i] = n
	}

	for i, num := range strings.Split(numGroups[1], "-") {
		n, _ := strconv.ParseInt(num, 10, 64)
		dr[i+2] = n
	}

	return DoubleRange{dr}
}

func parseTicket(line string) []int64 {
	ticket := []int64{}

	for _, num := range strings.Split(line, ",") {
		n, _ := strconv.ParseInt(num, 10, 64)
		ticket = append(ticket, n)
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

func findSingleton(poss []map[int]struct{}) (int, int) {
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
