package aoc2025

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func GetInput(day uint8) string {
	path := fmt.Sprintf("Day%d/input.txt", day)
	return ReadFileToString(path)
}

func GetTestInput(day uint8) string {
	path := fmt.Sprintf("Day%d/test.txt", day)
	return ReadFileToString(path)
}

func ReadFileToString(path string) string {
	file, err := os.Open(path)
	AssertSuccess(err, "Could not open file `"+path+"` verify that it exists in the same directory as the executable")
	defer func() { AssertSuccess(file.Close(), "Could not close file `"+path+"`") }()

	fileStat, err := file.Stat()
	AssertSuccess(err, "Could not get file stats")

	bytes := make([]byte, fileStat.Size())
	n, err := file.Read(bytes)
	AssertSuccess(err, "Could not read file contents")

	Assert(n == len(bytes), "Could not read entire file")
	return string(bytes)
}

func GetTrimmedLines(s string) []string {
	lines := strings.Split(s, "\n")
	suppressed := 0
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			suppressed++
			continue
		}
		lines[i-suppressed] = line
	}
	return lines[:len(lines)-suppressed]
}

func Assert(condition bool, message string) {
	if !condition {
		log.Fatalln("ASSERTION FAILED:\n\t" + message)
	}
}

func AssertEqual(actual, expected interface{}) {
	Assert(actual == expected, fmt.Sprintf("Expected `%v`, got `%v`", expected, actual))
}

func AssertSuccess(err error, message string) {
	if err != nil {
		Assert(false, message+": "+err.Error())
	}
}

type Addable interface {
	~int | ~int64 | ~float64 | ~string | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func Sum[T Addable](list []T) T {
	var sum T
	for _, v := range list {
		sum = sum + v
	}
	return sum
}
