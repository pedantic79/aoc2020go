package day08

import "testing"

func TestRegression(t *testing.T) {
	right := []int{1489, 1539}
	left := []interface{}{RunPart1().Value, RunPart2().Value}

	for i, l := range left {
		if right[i] != l {
			t.Errorf("Expected (%v), Actual (%v)", right[i], l)
		}
	}

}
