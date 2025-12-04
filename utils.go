package aoc2025

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func MakeRange(start, end uint64) []uint64 {
	if start > end {
		return []uint64{}
	}
	ret := make([]uint64, end-start)
	for i := uint64(0); i < end-start; i++ {
		ret[i] = start + i
	}
	return ret
}

func MakeRangeInclusive(start, end uint64) []uint64 { return MakeRange(start, end+1) }

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

func All[T any](list []T, fn func(T) bool) bool {
	for _, v := range list {
		if !fn(v) {
			return false
		}
	}
	return true
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

func Map[Arg, Ret any](list []Arg, fn func(Arg) Ret) []Ret {
	ret := make([]Ret, len(list))
	for i, v := range list {
		ret[i] = fn(v)
	}
	return ret
}

func MapWithIndex[Arg, Ret any](list []Arg, fn func(Arg, int) Ret) []Ret {
	ret := make([]Ret, len(list))
	for i, v := range list {
		ret[i] = fn(v, i)
	}
	return ret
}

func Filter[T any](list []T, fn func(T) bool) []T {
	ret := make([]T, 0)
	for _, v := range list {
		if fn(v) {
			ret = append(ret, v)
		}
	}
	return ret
}

func FilterWithIndex[T any](list []T, fn func(T, int) bool) []T {
	ret := make([]T, 0)
	for i, v := range list {
		if fn(v, i) {
			ret = append(ret, v)
		}
	}
	return ret
}

func MapSome[Arg, Ret any](list []Arg, fn func(Arg) (Ret, bool)) []Ret {
	ret := make([]Ret, 0)
	for _, v := range list {
		if r, ok := fn(v); ok {
			ret = append(ret, r)
		}
	}
	return ret
}

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
	return MapSome(lines, func(line string) (string, bool) {
		line = strings.TrimSpace(line)
		return line, line != ""
	})
}
