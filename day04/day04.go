package day04

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/pedantic79/aoc2020go/framework"
	"github.com/pedantic79/aoc2020go/util"
)

const day uint = 4

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

func parse(input string) []map[string]string {
	passports := []map[string]string{}
	split := strings.Split(input, "\n\n")

	for _, group := range split {
		fields := strings.FieldsFunc(group, func(r rune) bool { return r == ' ' || r == '\n' })

		passport := make(map[string]string)
		for _, field := range fields {
			kv := strings.Split(field, ":")

			passport[kv[0]] = kv[1]
		}
		passports = append(passports, passport)
	}

	return passports
}

func validatePart1(passport map[string]string) bool {
	for _, key := range []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"} {
		if _, ok := passport[key]; !ok {
			return false
		}
	}

	return true
}

func validatePart2(passport map[string]string) bool {
	for _, key := range []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"} {
		value, ok := passport[key]
		if !ok {
			return false
		}

		switch key {
		case "byr":
			n := util.Atoi(value)
			if !(1920 <= n && n <= 2002) {
				return false
			}
		case "iyr":
			n := util.Atoi(value)
			if !(2010 <= n && n <= 2020) {
				return false
			}
		case "eyr":
			n := util.Atoi(value)
			if !(2020 <= n && n <= 2030) {
				return false
			}
		case "hgt":
			n, _ := strconv.Atoi(value[:len(value)-2])

			if value[len(value)-2:] == "cm" {
				if !(150 <= n && n <= 193) {
					return false
				}
			} else {
				if !(59 <= n && n <= 76) {
					return false
				}
			}
		case "hcl":
			if len(value) != 7 || value[0] != '#' {
				return false
			}

			if _, err := strconv.ParseUint(value[1:], 16, 64); err != nil {
				return false
			}
		case "ecl":
			if !(value == "amb" || value == "blu" || value == "brn" || value == "gry" || value == "grn" || value == "hzl" || value == "oth") {
				return false
			}
		case "pid":
			if len(value) != 9 {
				return false
			}

			for _, d := range value {
				if !(d < unicode.MaxASCII && unicode.IsDigit(d)) {
					return false
				}
			}
		}
	}

	return true
}

func part1(passports []map[string]string) int {
	total := 0

	for _, passport := range passports {
		if validatePart1(passport) {
			total += 1
		}
	}

	return total
}

func part2(passports []map[string]string) int {
	total := 0

	for _, passport := range passports {
		if validatePart2(passport) {
			total += 1
		}
	}

	return total
}
