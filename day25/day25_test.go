package day25

import "testing"

func TestRegression(t *testing.T) {
	right := []uint64{1478097}
	left := []interface{}{RunPart1().Value}

	for i, l := range left {
		if right[i] != l {
			t.Errorf("Expected (%v), Actual (%v)", right[i], l)
		}
	}

}
