package main

import (
	"aoc2025"
	"strconv"
	"strings"
)

func combineRanges(ranges []aoc2025.Range) ([]aoc2025.Range, bool) {
	hasUpdated := false
	processedRanges := make([]aoc2025.Range, 0)
outer:
	for _, r := range ranges {
		for i := range processedRanges {
			if processedRanges[i].Contains(r.Start) {
				processedRanges[i].End = max(processedRanges[i].End, r.End)
				hasUpdated = true
				continue outer
			} else if processedRanges[i].Contains(r.End) {
				processedRanges[i].Start = min(processedRanges[i].Start, r.Start)
				hasUpdated = true
				continue outer
			} else if r.Contains(processedRanges[i].Start) || r.Contains(processedRanges[i].End) {
				processedRanges[i].Start = r.Start
				processedRanges[i].End = r.End
				hasUpdated = true
				continue outer
			}
		}
		processedRanges = append(processedRanges, r)
	}
	return processedRanges, hasUpdated
}

func main() {
	input := aoc2025.GetInput(5)
	wholeRanges, wholeIds, found := strings.Cut(input, "\n\n")
	aoc2025.Assert(found, "Unable to find range/ids separator")

	ranges := aoc2025.Map(aoc2025.GetTrimmedLines(wholeRanges), aoc2025.ParseRange)

	ids := aoc2025.Map(aoc2025.GetTrimmedLines(wholeIds), func(s string) uint64 {
		i, err := strconv.ParseUint(s, 10, 64)
		aoc2025.AssertSuccess(err, "Unable to parse id `"+s+"'")
		return i
	})

	FilteredIds := aoc2025.Filter(ids, func(id uint64) bool {
		for _, r := range ranges {
			if r.Contains(id) {
				return true
			}
		}
		return false
	})

	hasUpdated := true
	for hasUpdated {
		ranges, hasUpdated = combineRanges(ranges)
	}

	rangesLength := aoc2025.Map(ranges, func(r aoc2025.Range) uint64 { return r.End - r.Start + 1 })

	println("Part1: ", len(FilteredIds))
	println("Part2: ", aoc2025.Sum(rangesLength))
}
