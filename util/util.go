package util

import (
	"constraints"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"strings"
	"time"
)

var Day = flag.Uint("d", 0, "Day to run.")
var Part = flag.Uint("p", 3, "Part to run.")
var Results []func() AoCResult

func CheckDayAndPart(day, part uint) bool {
	return (*Day == 0 || *Day == day) && *Part&part > 0
}

func GenerateFileName(day uint) string {
	dir, _ := os.Getwd()

	if day == 0 {
		return fmt.Sprintf("%v/dayNN.txt", dir)
	}

	// This handles test cases which run in the directory rather than top-level
	if strings.HasSuffix(dir, fmt.Sprintf("day%02d", day)) {
		return fmt.Sprintf("%v/day%d.txt", dir, day)
	}

	return fmt.Sprintf("%v/day%02d/day%d.txt", dir, day, day)
}

func ReadFile(name string) string {
	content, err := ioutil.ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}

	// Convert []byte to string and print to screen
	text := string(content)
	return strings.TrimSuffix(text, "\n")
}

type AoCResult struct {
	Day       uint
	Part      uint
	ParseTime time.Duration
	RunTime   time.Duration
	Value     interface{}
}

func (r AoCResult) String() string {
	return fmt.Sprintf("Day %02d - Part %d: %v\n\tParser: %v\n\tRunner: %v\n\n", r.Day, r.Part, r.Value, r.ParseTime, r.RunTime)
}

func Timer[T any, R any](day uint, part uint, parse func(string) T, run func(T) R) AoCResult {
	input := ReadFile(GenerateFileName(day))

	start := time.Now()
	parsed := parse(input)
	parseTime := time.Since(start)

	start = time.Now()
	ans1 := run(parsed)
	runTime := time.Since(start)

	return AoCResult{Day: day, Part: part, ParseTime: parseTime, RunTime: runTime, Value: ans1}
}

func IntAbs[I constraints.Integer](i I) I {
	if i < 0 {
		return -i
	}

	return i
}

func ModInv(a, m int64) int64 {
	return new(big.Int).ModInverse(big.NewInt(a), big.NewInt(m)).Int64()
}

func ChineseRemainderTheorem(offsets, modulos []int64) int64 {
	product := int64(1)
	for _, m := range modulos {
		product *= m
	}

	total := int64(0)
	for i, off := range offsets {
		m := product / modulos[i]
		y := ModInv(m, modulos[i])

		total += off * m * y
	}

	return total % product
}

func Bool2Int[I constraints.Integer](b bool) I {
	if b {
		return 1
	}
	return 0
}
