package day06

import "testing"

func TestRegression(t *testing.T) {
	right := []int{6443, 3232}
	left := []interface{}{RunPart1().Value, RunPart2().Value}

	for i, l := range left {
		if right[i] != l {
			t.Errorf("Expected (%v), Actual (%v)", right[i], l)
		}
	}

}
