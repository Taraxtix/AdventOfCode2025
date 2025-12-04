package main

import (
	"aoc2025"
)

type Grid struct {
	elements [][]uint8
	removed  uint64
}

func gridFromLines(lines []string) Grid {
	return Grid{aoc2025.Map(lines, func(line string) []uint8 {
		return aoc2025.Map([]rune(line), func(c rune) uint8 {
			switch c {
			case '@':
				return uint8(1)
			case '.':
				return uint8(0)
			}
			return 0 // UNREACHABLE
		})
	}), 0}
}

func (grid Grid) isAccessible(x, y int) bool {
	count := uint8(0)
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			i := x + dx
			j := y + dy
			if (dx == 0 && dy == 0) || i < 0 || j < 0 || i >= len(grid.elements) || j >= len(grid.elements[0]) {
				continue
			}
			count += grid.elements[j][i]
		}
	}
	return count < 4 && grid.elements[y][x] == 1
}

func (grid Grid) clone() Grid {
	newGrid := Grid{make([][]uint8, len(grid.elements)), 0}
	for y, line := range grid.elements {
		newGrid.elements[y] = make([]uint8, len(line))
		for x, v := range line {
			newGrid.elements[y][x] = v
		}
	}
	return newGrid
}

func (grid *Grid) removeAccessible() uint64 {
	newGrid := grid.clone()

	for y, line := range grid.elements {
		for x := range line {
			if grid.isAccessible(x, y) {
				newGrid.elements[y][x] = 0
				newGrid.removed++
			}
		}
	}

	copy(grid.elements, newGrid.elements)
	grid.removed += newGrid.removed
	return newGrid.removed
}

func main() {
	input := aoc2025.GetInput(4)
	lines := aoc2025.GetTrimmedLines(input)

	grid := gridFromLines(lines)

	grid.removeAccessible()
	println("Part1: ", grid.removed)

	for grid.removeAccessible() > 0 {
	}
	println("Part2: ", grid.removed)
}
