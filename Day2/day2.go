package main

import (
	"aoc2025"
	"strconv"
	"strings"
)

func isInvalidPart1(v uint64) bool {
	str := strconv.FormatUint(v, 10)
	return str[:len(str)/2] == str[len(str)/2:]
}

func isInvalidPart2(v uint64) bool {
	str := strconv.FormatUint(v, 10)
	uSeqLengths := aoc2025.MakeRangeInclusive(1, uint64(len(str)/2))
	seqLengths := aoc2025.Map(uSeqLengths, func(seqLength uint64) int { return int(seqLength) })
	seqLengths = aoc2025.Filter(seqLengths, func(seqLength int) bool { return len(str)%seqLength == 0 })

seqLength:
	for _, seqLength := range seqLengths {
		for k := 1; k < len(str)/seqLength; k++ {
			start := (k - 1) * seqLength
			mid := k * seqLength
			end := (k + 1) * seqLength
			if str[start:mid] != str[mid:end] {
				continue seqLength
			}
		}
		return true
	}
	return false
}

func main() {
	input := aoc2025.GetInput(2)
	lines := aoc2025.GetTrimmedLines(input)
	aoc2025.AssertEqual(len(lines), 1)

	ranges := aoc2025.Map(strings.Split(lines[0], ","), aoc2025.ParseRange)

	invalidsPart1 := aoc2025.Flatten(
		aoc2025.Map(
			ranges,
			func(r aoc2025.Range) []uint64 { return aoc2025.Filter(r.AsList(), isInvalidPart1) },
		),
	)
	invalidsPart2 := aoc2025.Flatten(
		aoc2025.Map(
			ranges,
			func(r aoc2025.Range) []uint64 { return aoc2025.Filter(r.AsList(), isInvalidPart2) },
		),
	)

	println("Part1: " + strconv.FormatUint(aoc2025.Sum(invalidsPart1), 10))
	println("Part2: " + strconv.FormatUint(aoc2025.Sum(invalidsPart2), 10))
}
