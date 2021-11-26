package day22

import "testing"

func TestRegression(t *testing.T) {
	right := []int{33098, 35055}
	left := []interface{}{RunPart1().Value, RunPart2().Value}

	for i, l := range left {
		if right[i] != l {
			t.Errorf("Expected (%v), Actual (%v)", right[i], l)
		}
	}

}
