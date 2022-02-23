package startup

// Package startup only exists to seperate initialization from tests

import (
	"flag"

	_ "github.com/pedantic79/aoc2020go/framework"
)

func init() {
	flag.Parse()
}
