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

func isInvalidsPart2(v uint64) bool {
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

func ParseRange(s string) []uint64 {
	bounds := strings.Split(s, "-")
	start, err := strconv.ParseUint(bounds[0], 10, 64)
	aoc2025.AssertSuccess(err, "Unable to parse first bound (`"+bounds[0]+"`)")
	end, err := strconv.ParseUint(bounds[1], 10, 64)
	aoc2025.AssertSuccess(err, "Unable to parse first bound (`"+bounds[1]+"`)")
	return aoc2025.MakeRangeInclusive(start, end)
}

func main() {
	input := aoc2025.GetInput(2)
	lines := aoc2025.GetTrimmedLines(input)
	aoc2025.AssertEqual(len(lines), 1)

	rangesStr := strings.Split(lines[0], ",")
	ranges := make([][]uint64, len(rangesStr))
	for i, rangeStr := range rangesStr {
		ranges[i] = ParseRange(rangeStr)
	}

	invalidsPart1 := make([]uint64, 0)
	for _, range_ := range ranges {
		invalidsPart1 = append(invalidsPart1, aoc2025.Filter(range_, isInvalidPart1)...)
	}

	invalidsPart2 := make([]uint64, 0)
	for _, range_ := range ranges {
		invalidsPart2 = append(invalidsPart2, aoc2025.Filter(range_, isInvalidsPart2)...)
	}

	println("Part1: " + strconv.FormatUint(aoc2025.Sum(invalidsPart1), 10))
	println("Part2: " + strconv.FormatUint(aoc2025.Sum(invalidsPart2), 10))
}
