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
	parsedSrts := []srt{}
	scanner := bufio.NewScanner(&file)

	unparsedSection := ""
	for scanner.Scan() {
		scannedLine := scanner.Text()
		if strings.TrimSpace(scannedLine) == "" {
			if unparsedSection != "" {
				parsedSrts = append(parsedSrts, *newSrt(unparsedSection))
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
		parsedSrts = append(parsedSrts, *newSrt(unparsedSection))
	}

	return parsedSrts
}

func condenseSrt(parsedSrts []srt) []srt {
	condensedSrts := []srt{}

	return condensedSrts
}
