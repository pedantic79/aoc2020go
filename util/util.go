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

func IntPow[I constraints.Integer](n, m I) I {
    if m == 0 {
        return 1
    }
    result := n
    for i := I(2); i <= m; i++ {
        result *= n
    }
    return result
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

func Sum[N constraints.Integer](s []N) N {
	var total N

	for _, n := range s {
		total += n
	}

	return total
}

func SumFunc[T constraints.Integer, N any](s []N, fn func(N) T) T {
	var total T

	for _, n := range s {
		total += fn(n)
	}

	return total
}

func Map[T, U any](s []T, fn func(T) U) []U {
	ret := make([]U, 0, len(s))

	for _, v := range s {
		ret = append(ret, fn(v))
	}

	return ret
}

func Bool2Int[I constraints.Integer](b bool) I {
	if b {
		return 1
	}
	return 0
}
