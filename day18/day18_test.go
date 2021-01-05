package day18

import "testing"

func TestRegression(t *testing.T) {
	right := []int64{45283905029161, 216975281211165}
	left := []interface{}{RunPart1().Value, RunPart2().Value}

	for i, l := range left {
		if right[i] != l {
			t.Errorf("Expected (%v), Actual (%v)", right[i], l)
		}
	}

}
