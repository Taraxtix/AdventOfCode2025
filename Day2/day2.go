package main

import (
	"aoc2025"
	"strconv"
	"strings"
)

type Range struct {
	start uint64
	end   uint64
}

func (r Range) collectInvalidsPart1() []uint64 {
	invalids := make([]uint64, 0)
	for i := r.start; i <= r.end; i++ {
		str := strconv.FormatUint(i, 10)
		if len(str)%2 != 0 { // Odd number of digits can't be invalid
			continue
		}
		if str[:len(str)/2] == str[len(str)/2:] {
			invalids = append(invalids, i)
		}
	}
	return invalids
}

func (r Range) collectInvalidsPart2() []uint64 {
	invalids := make([]uint64, 0)
Value:
	for val := r.start; val <= r.end; val++ {
		str := strconv.FormatUint(val, 10)
	SequenceLength:
		for seqLength := 1; seqLength <= len(str)/2; seqLength++ {
			if len(str)%seqLength != 0 { // Cannot found a repeating sequence of this length
				continue
			}
			for k := 1; k < len(str)/seqLength; k++ {
				start := (k - 1) * seqLength
				mid := k * seqLength
				end := (k + 1) * seqLength
				if str[start:mid] != str[mid:end] {
					continue SequenceLength
				}
			}
			invalids = append(invalids, val)
			continue Value
		}
	}
	return invalids
}

func ParseRange(s string) Range {
	bounds := strings.Split(s, "-")
	start, err := strconv.ParseUint(bounds[0], 10, 64)
	aoc2025.AssertSuccess(err, "Unable to parse first bound (`"+bounds[0]+"`)")
	end, err := strconv.ParseUint(bounds[1], 10, 64)
	aoc2025.AssertSuccess(err, "Unable to parse first bound (`"+bounds[1]+"`)")
	return Range{start, end}
}

func main() {
	input := aoc2025.GetInput(2)
	lines := aoc2025.GetTrimmedLines(input)
	aoc2025.AssertEqual(len(lines), 1)

	rangesStr := strings.Split(lines[0], ",")
	ranges := make([]Range, len(rangesStr))
	for i, rangeStr := range rangesStr {
		ranges[i] = ParseRange(rangeStr)
	}

	invalidsPart1 := make([]uint64, 0)
	for _, range_ := range ranges {
		invalidsPart1 = append(invalidsPart1, range_.collectInvalidsPart1()...)
	}

	invalidsPart2 := make([]uint64, 0)
	for _, range_ := range ranges {
		invalidsPart2 = append(invalidsPart2, range_.collectInvalidsPart2()...)
	}

	println("Part1: " + strconv.FormatUint(aoc2025.Sum(invalidsPart1), 10))
	println("Part2: " + strconv.FormatUint(aoc2025.Sum(invalidsPart2), 10))
}
