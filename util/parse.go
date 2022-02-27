package util

import (
	"fmt"
	"log"
	"strconv"

	"golang.org/x/exp/constraints"
)

func Atoi(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		log.Panicf("parsing [%v] to int failed", s)
	}

	return v
}

func ParseInteger[I constraints.Integer](s string) I {
	var i I
	fmt.Sscan(s, &i)
	return i
}
