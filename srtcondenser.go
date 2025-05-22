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
	out := []srt{}
	scanner := bufio.NewScanner(&file)

	curr := ""
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			if curr != "" {
				out = append(out, *newSrt(curr))
				curr = ""
			}
		} else {
			if curr != "" {
				curr += "\n"
			}
			curr += line
		}
	}

	if curr != "" {
		out = append(out, *newSrt(curr))
	}

	return out
}
