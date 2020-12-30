package day10

import "testing"

func TestRegression(t *testing.T) {
	right := []uint64{2310, 64793042714624}
	left := []interface{}{uint64(RunPart1().Value.(int)), RunPart2().Value}

	for i, l := range left {
		if right[i] != l {
			t.Errorf("Expected (%v), Actual (%v)", right[i], l)
		}
	}

}
