package main

import (
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
