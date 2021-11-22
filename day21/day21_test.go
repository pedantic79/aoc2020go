package day21

import "testing"

func TestRegression(t *testing.T) {
	right := []interface{}{2230, "qqskn,ccvnlbp,tcm,jnqcd,qjqb,xjqd,xhzr,cjxv"}
	left := []interface{}{RunPart1().Value, RunPart2().Value}

	for i, l := range left {
		if right[i] != l {
			t.Errorf("Expected (%v), Actual (%v)", right[i], l)
		}
	}

}

func TestSamplePart1(t *testing.T) {
	input := `mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)`

	parsed := parse(input)
	left := part1(parsed)
	right := 5

	if right != left {
		t.Errorf("Expected (%v), Actual (%v)", right, left)
	}
}

func TestSamplePart2(t *testing.T) {
	input := `mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)`

	parsed := parse(input)
	left := part2(parsed)
	right := "mxmxvkd,sqjhc,fvjkl"

	if right != left {
		t.Errorf("Expected (%v), Actual (%v)", right, left)
	}
}
