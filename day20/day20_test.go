package day20

import "testing"

func TestRegression(t *testing.T) {
	right := [2]int{104831106565027, -1}
	left := [2]interface{}{RunPart1().Value, RunPart2().Value}

	for i, l := range left {
		if right[i] != l {
			t.Errorf("Expected (%v), Actual (%v)", right[i], l)
		}
	}

}

func TestEdgeId(t *testing.T) {
	right := [2]int{713, 589}
	a, b := CalculateEdgeValue([]bool{true, false, true, true, false, false, true, false, false, true})

	if right[0] != a {
		t.Errorf("Expected (%v), Actual (%v)", right[0], a)
	}

	if right[1] != b {
		t.Errorf("Expected (%v), Actual (%v)", right[1], b)
	}

}
