package day22

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/pedantic79/aoc2020go/util"
	"github.com/pedantic79/aoc2020go/util/set"
)

var day uint = 22
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

func parse(input string) [2][]int {
	a := [2][]int{}
	index := 0

	for _, line := range strings.Split(input, "\n") {
		if line == "Player 2:" {
			index = 1
		}

		if n, err := strconv.Atoi(line); err == nil {
			a[index] = append(a[index], n)
		}
	}

	return a
}

func calculateScore(player []int) int {
	score := 0
	for i, n := range player {
		score += (len(player) - i) * n
	}

	return score
}

func part1(players [2][]int) int {
	for len(players[0]) > 0 && len(players[1]) > 0 {
		player0 := players[0][0]
		players[0] = players[0][1:]

		player1 := players[1][0]
		players[1] = players[1][1:]

		if player0 > player1 {
			players[0] = append(players[0], player0, player1)
		} else {
			players[1] = append(players[1], player1, player0)
		}
	}

	if len(players[0]) > 0 {
		return calculateScore(players[0])
	} else {
		return calculateScore(players[1])
	}
}

func solve2(player1, player2 *[]int) int {
	seen := make(set.Set[string])

	for len(*player1) > 0 && len(*player2) > 0 {
		s := fmt.Sprint(*player1)
		if seen.Contains(s) {
			return 1
		} else {
			seen.Add(s)
		}

		var p1, p2 int
		p1, *player1 = (*player1)[0], (*player1)[1:]
		p2, *player2 = (*player2)[0], (*player2)[1:]

		var result bool
		if p1 <= len(*player1) && p2 <= len(*player2) {
			player1Copy := make([]int, p1)
			player2Copy := make([]int, p2)
			copy(player1Copy, (*player1)[:p1])
			copy(player2Copy, (*player2)[:p2])

			result = solve2(&player1Copy, &player2Copy) == 1
		} else {
			result = p1 > p2
		}

		if result {
			*player1 = append(*player1, p1, p2)
		} else {
			*player2 = append(*player2, p2, p1)
		}
	}

	if len(*player1) == 0 {
		return 2
	} else {
		return 1
	}
}

func part2(players [2][]int) int {
	num := solve2(&players[0], &players[1])
	return calculateScore(players[num-1])
}
