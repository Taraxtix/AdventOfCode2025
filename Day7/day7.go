package main

import (
	"aoc2025"
	"strconv"
)

type State = int

const (
	Free = State(iota)
	Start
	Splitter
	HitSplitter
	Beamed // 4 or more == Beamed from (value-3) different timelines
)

func ParseGrid(lines []string) [][]State {
	grid := aoc2025.Map(make([][]State, len(lines)), func(line []State) []State { return make([]State, len(lines[0])) })
	for y, line := range lines {
		for x, c := range line {
			switch c {
			case '^':
				grid[y][x] = Splitter
			case 'S':
				grid[y][x] = Start
			}
		}
	}
	return grid
}

func StateToString(state State) string {
	switch state {
	case Free:
		return "."
	case Start:
		return "S"
	case Splitter:
		return "^"
	case HitSplitter:
		return "-"
	default:
		return strconv.Itoa(state - 3)
	}
}

func playLine(line []State, previousLine []State) []State {
	playedLine := aoc2025.Map(line, func(state State) State { return state }) // Deep copy

	for x, state := range playedLine {
		switch state {
		case Free:
			switch previousLine[x] {
			case Start:
				playedLine[x] = Beamed
			default:
				if previousLine[x] >= Beamed {
					playedLine[x] = previousLine[x]
				} else {
					playedLine[x] = state
				}
			}
			break
		case Splitter:
			if previousLine[x] >= Beamed {
				if x-1 >= 0 {
					if playedLine[x-1] == Free {
						playedLine[x-1] = previousLine[x]
					} else {
						playedLine[x-1] += previousLine[x] - 3
					}
				}
				if x+1 < len(playedLine) {
					playedLine[x+1] = previousLine[x]
				}
				playedLine[x] = HitSplitter
			}
			break
		default:
			if playedLine[x] >= Beamed && previousLine[x] >= Beamed {
				playedLine[x] += previousLine[x] - 3
			} else {
				playedLine[x] = state
			}
			break
		}
	}
	return playedLine
}

func PlayGrid(grid [][]State) [][]State {
	playedGrid := aoc2025.Map(grid, func(line []State) []State { return line }) // Deep copy

	for y, line := range playedGrid {
		if y == 0 {
			continue
		}
		playedGrid[y] = playLine(line, playedGrid[y-1])
	}

	return playedGrid
}

func DebugPrintGrid(grid [][]State) {
	for _, line := range grid {
		for _, state := range line {
			print(StateToString(state))
		}
		println()
	}
}

func main() {
	input := aoc2025.GetInput(7)
	trimmedLines := aoc2025.GetTrimmedLines(input)
	grid := ParseGrid(trimmedLines)
	playedGrid := PlayGrid(grid)

	//println("==================================")
	//DebugPrintGrid(grid)
	//println("==================================")
	//DebugPrintGrid(playedGrid)
	//println("==================================")

	nbSplit := aoc2025.Sum(aoc2025.Map(playedGrid, func(line []State) int {
		return len(aoc2025.Filter(line, func(state State) bool { return state == HitSplitter }))
	}))

	nbTimeline := aoc2025.Sum(aoc2025.Map(playedGrid[len(grid)-1], func(state State) int {
		if state >= Beamed {
			return state - 3
		}
		return 0
	}))

	println("Part1: ", nbSplit)
	println("Part2: ", nbTimeline)
}
