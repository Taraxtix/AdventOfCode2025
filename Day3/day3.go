package main

import (
	"aoc2025"
	"math"
	"strconv"
	"strings"
)

func MaxVoltages(bank []uint64, nbActivate uint64) uint64 {
	values := make([]uint64, nbActivate)
	for i, v := range bank {
		if aoc2025.All(values, func(v uint64) bool { return v == 9 }) {
			break
		}
		for j := int(nbActivate - 1); j >= 0; j-- {
			if values[j] != 9 && i < len(bank)-j && v > values[j] {
				values[j] = v
				for k := j - 1; k >= 0; k-- {
					values[k] = 0
				}
				break
			}
		}
	}
	ret := uint64(0)
	for i, v := range values {
		ret += uint64(math.Pow10(i)) * v
	}
	return ret
}

func main() {
	input := aoc2025.GetInput(3)
	lines := aoc2025.GetTrimmedLines(input)
	banks := aoc2025.Map(lines, func(line string) []uint64 {
		return aoc2025.Map(strings.Split(line, ""), func(s string) uint64 {
			i, err := strconv.ParseUint(s, 10, 8)
			aoc2025.AssertSuccess(err, "Unable to parse digit `"+s+"'")
			return i
		})
	})
	maxVoltages2 := aoc2025.Map(banks, func(bank []uint64) uint64 { return MaxVoltages(bank, 2) })
	maxVoltages12 := aoc2025.Map(banks, func(bank []uint64) uint64 { return MaxVoltages(bank, 12) })

	println("Part1: ", aoc2025.Sum(maxVoltages2))
	println("Part2: ", aoc2025.Sum(maxVoltages12))
}
