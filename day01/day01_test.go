package day01

import "testing"

func TestRegression(t *testing.T) {
	right := []int{1018944, 8446464}
	left := []interface{}{RunPart1().Value, RunPart2().Value}

	for i, l := range left {
		if right[i] != l {
			t.Errorf("Expected (%v), Actual (%v)", right[i], l)
		}
	}

}
