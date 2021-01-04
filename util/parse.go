package util

import (
	"log"
	"strconv"
)

func Atoi(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		log.Panicf("parsing [%v] to int failed", s)
	}

	return v
}

func ParseInt(s string, base int, bitSize int) int64 {
	v, err := strconv.ParseInt(s, base, bitSize)

	if err != nil {
		log.Panicf("parsing [%v] to int64(base=%v, bitSize=%v) failed", s, base, bitSize)
	}

	return v
}

func ParseUint(s string, base int, bitSize int) uint64 {
	v, err := strconv.ParseUint(s, base, bitSize)

	if err != nil {
		log.Panicf("parsing [%v] to int64(base=%v, bitSize=%v) failed", s, base, bitSize)
	}

	return v
}
