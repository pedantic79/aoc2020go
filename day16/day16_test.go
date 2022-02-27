package day16

import "testing"

func TestRegression(t *testing.T) {
	right := []int{21071, 3429967441937}
	left := []interface{}{RunPart1().Value, RunPart2().Value}

	for i, l := range left {
		if right[i] != l {
			t.Errorf("Expected (%v), Actual (%v)", right[i], l)
		}
	}

}
