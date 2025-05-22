package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type srt struct {
	count     int
	timestamp string
	subtitle  string
}

func newSrt(blob string) *srt {
	lines := strings.Split(strings.TrimSpace(blob), "\n")
	if len(lines) < 3 {
		panic(fmt.Sprintf("Invalid SRT block (expected at least 3 lines): %q", blob))
	}

	count, err := strconv.Atoi(lines[0])
	if err != nil {
		panic(fmt.Sprintf("Invalid count in block: %q", lines[0]))
	}

	timestamp := strings.TrimSpace(lines[1])

	subtle := strings.Join(lines[2:], "\n")

	return &srt{
		count:     count,
		timestamp: timestamp,
		subtitle:  subtle,
	}
}

func parseSrt(file os.File) []srt {
	parsedSrts := []srt{}
	scanner := bufio.NewScanner(&file)

	unparsedSection := ""
	for scanner.Scan() {
		line := scanner.Text()
		unparsedSection += fmt.Sprintf("%s\n", line)
		if strings.TrimSpace(line) == "" {
			parsedSrts = append(parsedSrts, *newSrt(unparsedSection))
			unparsedSection = ""
		}
	}

	if strings.TrimSpace(unparsedSection) != "" {
		parsedSrts = append(parsedSrts, *newSrt(unparsedSection))
	}

	return parsedSrts
}

func condenseSrt(parsedSrts []srt) []srt {
	if len(parsedSrts) == 0 {
		return []srt{}
	}

	var result []srt
	i := 0

	for i < len(parsedSrts) {
		curr := parsedSrts[i]
		j := i + 1

		for j < len(parsedSrts) && parsedSrts[j].timestamp == curr.timestamp {
			curr.subtitle += "\n" + parsedSrts[j].subtitle
			j++
		}

		result = append(result, curr)
		i = j
	}

	for i := range result {
		result[i].count = i + 1
	}

	return result
}

func writeSrt(parsedSrts []srt, outputPath string) {

}
