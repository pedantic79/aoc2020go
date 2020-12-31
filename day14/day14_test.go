package day14

import "testing"

func TestRegression(t *testing.T) {
	right := []uint64{6386593869035, 4288986482164}
	left := []interface{}{RunPart1().Value, RunPart2().Value}

	for i, l := range left {
		if right[i] != l {
			t.Errorf("Expected (%v), Actual (%v)", right[i], l)
		}
	}

}
