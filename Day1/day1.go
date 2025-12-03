package main

import (
	"aoc2025"
	"log"
	"strconv"
)

type Direction = uint8

const (
	Left Direction = iota
	Right
)

type Rotation struct {
	direction Direction
	steps     uint64
}

func RotationFromString(s string) Rotation {
	aoc2025.Assert(len(s) >= 2, "Rotation string too short: `"+s+"`.")
	steps, err := strconv.ParseUint(s[1:], 10, 64)
	aoc2025.AssertSuccess(err, "Unable to parse rotation steps (`"+s[1:]+"`): ")

	switch s[0] {
	case 'L':
		return Rotation{Left, steps}
	case 'R':
		return Rotation{Right, steps}
	}

	log.Fatalln("Unable to parse rotation direction (`" + s[0:1] + "`)")
	return Rotation{}
}

type Dial struct {
	position   uint64
	zeroPassed uint64
	atZero     uint64
}

func (d *Dial) rotate(rotation Rotation) {
	switch rotation.direction {
	case Left:
		for rotation.steps > d.position {
			rotation.steps -= d.position + 1
			if d.position != 0 {
				d.zeroPassed++
			}
			d.position = 99
		}
		d.position -= rotation.steps
		if d.position == 0 {
			d.zeroPassed++
		}
		break
	case Right:
		for rotation.steps > 99-d.position {
			rotation.steps -= 100 - d.position
			d.position = 0
			d.zeroPassed++
		}
		d.position += rotation.steps
		break
	}
	if d.position == 0 {
		d.atZero++
	}
}

func (d *Dial) rotateAll(rotations []Rotation) {
	for _, rotation := range rotations {
		d.rotate(rotation)
	}
}

func main() {
	input := aoc2025.GetInput(1)
	lines := aoc2025.GetTrimmedLines(input)

	rotations := aoc2025.Map(lines, RotationFromString)

	dial := Dial{50, 0, 0}
	dial.rotateAll(rotations)
	println("Part1: " + strconv.Itoa(int(dial.atZero)))
	println("Part2: " + strconv.Itoa(int(dial.zeroPassed)))
}
