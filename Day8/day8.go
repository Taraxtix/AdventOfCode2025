package main

import (
	"aoc2025"
	"slices"
	"sort"
	"strings"
)

type JunctionBox struct {
	x       uint64
	y       uint64
	z       uint64
	circuit int
}

type Distance struct {
	distanceSquared uint64
	from            int
	to              int
}

func getDistances(junctionBoxes []JunctionBox) []Distance {
	distances := make([]Distance, 0)
	for from := 0; from < len(junctionBoxes); from++ {
		for to := from + 1; to < len(junctionBoxes); to++ {
			a := junctionBoxes[from]
			b := junctionBoxes[to]
			distances = append(distances, Distance{
				(a.x-b.x)*(a.x-b.x) + (a.y-b.y)*(a.y-b.y) + (a.z-b.z)*(a.z-b.z),
				from,
				to,
			})
		}
	}
	return distances
}

func connect(junctionBoxes []JunctionBox, distances []Distance) (int, int) {
	lastConnectionFrom := 0
	lastConnectionTo := 0
	for _, d := range distances {
		if junctionBoxes[d.from].circuit == junctionBoxes[d.to].circuit {
			continue
		}
		lastConnectionFrom = d.from
		lastConnectionTo = d.to
		if junctionBoxes[d.to].circuit < junctionBoxes[d.from].circuit {
			t := d.to
			d.to = d.from
			d.from = t
		}
		for i := 0; i < len(junctionBoxes); i++ {
			if i != d.to && junctionBoxes[i].circuit == junctionBoxes[d.to].circuit {
				junctionBoxes[i].circuit = junctionBoxes[d.from].circuit
			}
		}
		junctionBoxes[d.to].circuit = junctionBoxes[d.from].circuit
	}
	return lastConnectionFrom, lastConnectionTo
}

func main() {
	toConnectFirst := 1000
	input := aoc2025.GetInput(8)
	trimmedLines := aoc2025.GetTrimmedLines(input)

	junctionBoxes := aoc2025.MapWithIndex(trimmedLines, func(line string, idx int) JunctionBox {
		coordinates := strings.Split(line, ",")
		return JunctionBox{
			aoc2025.AssertedParseUint64(coordinates[0]),
			aoc2025.AssertedParseUint64(coordinates[1]),
			aoc2025.AssertedParseUint64(coordinates[2]),
			idx,
		}
	})

	distances := getDistances(junctionBoxes)
	sort.Slice(distances, func(i, j int) bool { return distances[i].distanceSquared < distances[j].distanceSquared })

	connect(junctionBoxes, distances[:toConnectFirst])

	circuitLengths := make([]int, len(junctionBoxes))
	for i := range circuitLengths {
		circuitLengths[i] = len(aoc2025.Filter(junctionBoxes, func(j JunctionBox) bool { return j.circuit == i }))
	}
	sort.Ints(circuitLengths)
	slices.Reverse(circuitLengths)

	lastConnectionFrom, lastConnectionTo := connect(junctionBoxes, distances[toConnectFirst:])

	println("Part1: ", aoc2025.Reduce(circuitLengths[:3], 1, func(acc, val int) int { return acc * val }))
	println("Part2: ", junctionBoxes[lastConnectionFrom].x*junctionBoxes[lastConnectionTo].x)
}
