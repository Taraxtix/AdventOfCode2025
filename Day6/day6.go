package main

import (
	"aoc2025"
	"slices"
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

func parseValueColumnsAsCephalopod(lines []string) [][]uint64 {
	lineLength := aoc2025.Map(lines, func(line string) int { return len(line) })

	columnsChars := aoc2025.Map(
		make([]any, slices.Max(lineLength)),
		func(column any) []string {
			return make([]string, len(lines))
		},
	)
	for i, line := range lines {
		for j, value := range strings.Split(line, "") {
			if columnsChars[j] == nil {
				columnsChars[j] = make([]string, len(lines))
			}
			columnsChars[j][i] = value
		}
	}
	columnsStr := aoc2025.Map(
		columnsChars,
		func(chars []string) string {
			return strings.TrimSpace(strings.Join(chars, ""))
		},
	)
	sectors := len(aoc2025.Filter(columnsStr, func(s string) bool { return s == "" })) + 1
	idx := 0
	columns := make([][]uint64, sectors)
	for _, val := range columnsStr {
		if val == "" {
			idx++
			continue
		}
		columns[idx] = append(columns[idx], aoc2025.AssertedParseUint64(val))
	}

	return columns
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
	trimmedLines := aoc2025.GetTrimmedLines(input)
	lines := make([]string, 0)
	for line := range strings.Lines(input) {
		lines = append(lines, strings.Trim(line, "\n"))
	}

	humanValuesColumns := parseValueColumnsAsHuman(trimmedLines[0 : len(trimmedLines)-1])
	cephalopodValuesColumns := parseValueColumnsAsCephalopod(lines[0 : len(lines)-1])

	operators := aoc2025.Filter(
		aoc2025.Map(strings.Split(trimmedLines[len(trimmedLines)-1], " "), strings.TrimSpace),
		func(s string) bool { return s != "" },
	)

	println("Part1: ", getResult(humanValuesColumns, operators))
	println("Part2: ", getResult(cephalopodValuesColumns, operators))
}
