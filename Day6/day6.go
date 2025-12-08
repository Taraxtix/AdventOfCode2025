package main

import (
	"aoc2025"
	"strings"
)

func parseValueColumnsAsHuman(lines []string) [][]uint64 {
	valuesLines := aoc2025.Map(
		lines,
		func(line string) []uint64 {
			return aoc2025.Map(
				aoc2025.Filter(
					strings.Split(line, " "),
					func(s string) bool { return s != "" },
				),
				func(s string) uint64 { return aoc2025.AssertedParseUint64(s) },
			)
		},
	)
	columns := make([][]uint64, len(valuesLines[0]))
	for _, line := range valuesLines {
		for i, value := range line {
			columns[i] = append(columns[i], value)
		}
	}
	return columns
}

func parseValueColumnsAsCephalopod(line []string) [][]uint64 {
	return nil
}

func baseForOp(s string) uint64 {
	switch s {
	case "+":
		return 0
	case "*":
		return 1
	}

	aoc2025.Assert(false, "Unknown operator `"+s+"`.")
	return 0
}

func getResult(values [][]uint64, operators []string) uint64 {
	return aoc2025.Sum(
		aoc2025.MapWithIndex(values, func(column []uint64, columnIndex int) uint64 {
			return aoc2025.Reduce(column, baseForOp(operators[columnIndex]), func(acc uint64, value uint64) uint64 {
				switch operators[columnIndex] {
				case "+":
					return acc + value
				case "*":
					return acc * value
				}
				aoc2025.Assert(false, "Unknown operator `"+operators[columnIndex]+"`.")
				return 0
			})
		}),
	)
}

func main() {
	input := aoc2025.GetInput(6)
	lines := aoc2025.GetTrimmedLines(input)

	humanValuesColumns := parseValueColumnsAsHuman(lines[0 : len(lines)-1])
	cephalopodValuesColumns := parseValueColumnsAsCephalopod(lines[0 : len(lines)-1])

	operators := aoc2025.Filter(
		aoc2025.Map(strings.Split(lines[len(lines)-1], " "), strings.TrimSpace),
		func(s string) bool { return s != "" },
	)

	println("Part1: ", getResult(humanValuesColumns, operators))
	println("Part2: ", getResult(cephalopodValuesColumns, operators))
}
