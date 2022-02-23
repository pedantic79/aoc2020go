package util

import (
	"math/big"

	"golang.org/x/exp/constraints"
)

func IntAbs[I constraints.Integer](i I) I {
	if i < 0 {
		return -i
	}

	return i
}

func ModInv(a, m int64) int64 {
	return new(big.Int).ModInverse(big.NewInt(a), big.NewInt(m)).Int64()
}

func ChineseRemainderTheorem(offsets, modulos []int64) int64 {
	product := int64(1)
	for _, m := range modulos {
		product *= m
	}

	total := int64(0)
	for i, off := range offsets {
		m := product / modulos[i]
		y := ModInv(m, modulos[i])

		total += off * m * y
	}

	return total % product
}

func Bool2Int[I constraints.Integer](b bool) I {
	if b {
		return 1
	}
	return 0
}
