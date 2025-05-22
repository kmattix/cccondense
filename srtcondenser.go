package main

import (
	"bufio"
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
	lines := strings.Split(blob, "\n")
	s := srt{
		timestamp: lines[1],
		subtitle:  lines[2],
	}
	count, err := strconv.Atoi(lines[0])
	if err == nil {
		s.count = count
	}
	return &s
}

func parseSrt(file os.File) []srt {
	parsedSlc := []srt{}
	scanner := bufio.NewScanner(&file)

	unparsedSection := ""
	for scanner.Scan() {
		scannedLine := scanner.Text()
		if strings.TrimSpace(scannedLine) == "" {
			if unparsedSection != "" {
				parsedSlc = append(parsedSlc, *newSrt(unparsedSection))
				unparsedSection = ""
			}
		} else {
			if unparsedSection != "" {
				unparsedSection += "\n"
			}
			unparsedSection += scannedLine
		}
	}

	if unparsedSection != "" {
		parsedSlc = append(parsedSlc, *newSrt(unparsedSection))
	}

	return parsedSlc
}
