package day21

import (
	"sort"
	"strings"
	"time"

	"github.com/pedantic79/aoc2020go/util"
	"github.com/pedantic79/aoc2020go/util/set"
)

var day uint = 21
var fileName string = util.GenerateFileName(day)

type Allergy = string
type Ingredient = string

type Food struct {
	Ingredients []Ingredient
	Allergens   []Allergy
}

type AllergySet set.Set[Allergy]
type IngredientSet set.Set[Ingredient]

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

func parse(input string) []Food {
	foods := []Food{}

	for _, line := range strings.Split(input, "\n") {
		l := strings.Split(line, " (contains ")
		l[1] = strings.TrimRight(l[1], ")")
		f := Food{
			Ingredients: strings.Split(l[0], " "),
			Allergens:   strings.Split(l[1], ", "),
		}

		foods = append(foods, f)
	}

	return foods
}

func contains(haystack []Ingredient, needle Ingredient) bool {
	for _, f := range haystack {
		if f == needle {
			return true
		}
	}

	return false
}

func solve(foods []Food) (map[Ingredient]AllergySet, map[Ingredient]int) {
	possibilities := make(map[Ingredient]AllergySet)
	for _, food := range foods {
		for _, ing := range food.Ingredients {
			if _, ok := possibilities[ing]; !ok {
				possibilities[ing] = make(AllergySet)
			}
		}
	}

	for _, food := range foods {
		for _, all := range food.Allergens {
			for k := range possibilities {
				entry := set.Set[string](possibilities[k])
				entry.Add(all)
			}
		}
	}

	frequency := make(map[Ingredient]int)
	for _, food := range foods {
		for _, ing := range food.Ingredients {
			frequency[ing]++
		}

		for _, allergen := range food.Allergens {
			for key, s := range possibilities {
				if !contains(food.Ingredients, key) {
					delete(s, allergen)
				}
			}

		}
	}

	return possibilities, frequency
}

func part1(foods []Food) int {
	possibilities, frequency := solve(foods)

	sum := 0
	for ing, allergens := range possibilities {
		if len(allergens) == 0 {
			sum += frequency[ing]
		}
	}
	return sum
}

func part2(foods []Food) string {
	possibilities, _ := solve(foods)
	allergenMap := make(map[Allergy]Ingredient)

	for {
		var found *string
		found = nil
		for ing, v := range possibilities {
			if len(v) == 1 {
				for a := range v {
					allergenMap[a] = ing
					found = &a
					break
				}
			}
			if found != nil {
				break
			}
		}

		if found != nil {
			for _, v := range possibilities {
				delete(v, *found)
			}
		} else {
			break
		}
	}

	keys := make([]string, 0, len(allergenMap))
	for k := range allergenMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	output := strings.Builder{}
	for i, k := range keys {
		if i > 0 {
			output.WriteRune(',')
		}
		output.WriteString(allergenMap[k])
	}

	return output.String()
}
