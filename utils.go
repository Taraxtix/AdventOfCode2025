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
	if err != nil {
		log.Fatalln("Could not open file `" + path + "` verify that it exists in the same directory as the executable:" + err.Error())
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}()

	fileStat, err := file.Stat()
	if err != nil {
		log.Fatalln("Could not get file stats" + err.Error())
	}

	bytes := make([]byte, fileStat.Size())
	n, err := file.Read(bytes)
	if err != nil {
		log.Fatalln("Could not read file contents" + err.Error())
	}
	if n < len(bytes) {
		log.Fatalln("Could not read entire file")
	}

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
