package day14

import (
	"strings"

	"github.com/pedantic79/aoc2020go/framework"
	"github.com/pedantic79/aoc2020go/util"
)

const day uint = 14

func init() {
	if framework.CheckDayAndPart(day, 1) {
		framework.Results = append(framework.Results, RunPart1)
	}

	if framework.CheckDayAndPart(day, 2) {
		framework.Results = append(framework.Results, RunPart2)
	}
}

func RunPart1() framework.AoCResult {
	return framework.Timer(day, 1, parse, part1)
}

func RunPart2() framework.AoCResult {
	return framework.Timer(day, 2, parse, part2)
}

type MemoryOp struct {
	location uint64
	value    uint64
}

type ProgramSegment struct {
	mask      [36]byte
	operation []MemoryOp
}

func parse(input string) []ProgramSegment {
	segments := []ProgramSegment{}

	for _, segment := range strings.Split(input, "mask = ") {
		// first segment is blank
		if len(segment) == 0 {
			continue
		}
		lines := strings.Split(segment, "\n")

		var mask [36]byte
		for i, b := range lines[0][:36] {
			j := 36 - i - 1
			switch b {
			case '0':
				mask[j] = 0
			case '1':
				mask[j] = 1
			case 'X':
				mask[j] = 2
			}
		}
		// copy(mask[:], lines[0][:36])

		operation := []MemoryOp{}

		for _, memOp := range lines[1:] {
			if len(memOp) == 0 {
				continue
			}
			memSplit := strings.Split(memOp, "] = ")
			l := util.ParseUint(memSplit[0][4:], 10, 64)
			r := util.ParseUint(memSplit[1], 10, 64)

			operation = append(operation, MemoryOp{
				location: l,
				value:    r,
			})
		}

		segments = append(segments, ProgramSegment{
			mask:      mask,
			operation: operation,
		})

	}

	// fmt.Println(segments)
	return segments
}

func part1(segments []ProgramSegment) uint64 {
	memory := make(map[uint64]uint64)

	for _, segment := range segments {
		for _, op := range segment.operation {
			n := op.value

			for i := 0; i < 36; i++ {
				var mask uint64

				switch segment.mask[i] {
				case 2:
					continue
				case 1:
					if (n>>i)&1 == 1 {
						continue
					}
					mask = 1 << i
				case 0:
					if (n>>i)&1 == 0 {
						continue
					}
					mask = 1 << i
				}

				n ^= mask
			}

			// log.Println(op.location, n)
			memory[op.location] = n
		}
	}

	var total uint64
	for _, v := range memory {
		total += v
	}
	return total
}

func mergeMask(num uint64, mask [36]byte) [36]byte {
	var address [36]byte

	for i := 0; i < 36; i++ {
		if num&(1<<i) > 0 || mask[i] == 1 {
			address[i] = 1
		}

		if mask[i] == 2 {
			address[i] = 2
		}
	}

	return address
}

func generateAddresses(value [36]byte) [][36]byte {
	output := [][36]byte{}
	stack := [][36]byte{}
	stack = append(stack, value)

outer:
	for len(stack) > 0 {
		value := stack[len(stack)-1] // peek
		stack = stack[:len(stack)-1] // pop

		for i := 0; i < 36; i++ {
			if value[i] == 2 {
				var modified [36]byte
				copy(modified[:], value[:])
				modified[i] = 0
				stack = append(stack, modified)

				modified[i] = 1
				stack = append(stack, modified)
				continue outer
			}
		}

		output = append(output, value)
	}

	return output
}

func mask2Num(value [36]byte) uint64 {
	var num uint64
	for i := 35; i >= 0; i-- {
		num = num*2 + uint64(value[i])
	}

	return num
}

func part2(segments []ProgramSegment) uint64 {
	memory := make(map[uint64]uint64)

	for _, segment := range segments {
		for _, op := range segment.operation {
			address := mergeMask(op.location, segment.mask)
			for _, add := range generateAddresses(address) {
				memory[mask2Num(add)] = op.value
			}
		}
	}

	var total uint64
	for _, v := range memory {
		total += v
	}
	return total
}
