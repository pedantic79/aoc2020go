package day13

import "testing"

func TestRegression(t *testing.T) {
	right := []int64{1915, 294354277694107}
	left := []interface{}{RunPart1().Value, RunPart2().Value}

	for i, l := range left {
		if right[i] != l {
			t.Errorf("Expected (%v), Actual (%v)", right[i], l)
		}
	}

}
