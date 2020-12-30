package util

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

var Empty struct{}
var Day = flag.Uint("d", 0, "Day to run.")
var Part = flag.Uint("p", 3, "Part to run.")
var Results []func() AoCResult

func CheckDayAndPart(day, part uint) bool {
	return (*Day == 0 || *Day == day) && *Part&part > 0
}

func GenerateFileName(day uint) string {
	dir, _ := os.Getwd()

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
