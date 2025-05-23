package condenser

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Srt struct {
	count     int
	timestamp string
	subtitle  string
}

func NewSrt(blob string) *Srt {
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

	return &Srt{
		count:     count,
		timestamp: timestamp,
		subtitle:  subtle,
	}
}

func ParseSrt(file *os.File) []Srt {
	parsedSrts := []Srt{}
	scanner := bufio.NewScanner(file)

	unparsedSection := ""
	for scanner.Scan() {
		line := scanner.Text()
		unparsedSection += fmt.Sprintf("%s\n", line)
		if strings.TrimSpace(line) == "" {
			parsedSrts = append(parsedSrts, *NewSrt(unparsedSection))
			unparsedSection = ""
		}
	}

	if strings.TrimSpace(unparsedSection) != "" {
		parsedSrts = append(parsedSrts, *NewSrt(unparsedSection))
	}

	return parsedSrts
}

func CondenseSrt(parsedSrts []Srt) []Srt {
	if len(parsedSrts) == 0 {
		return []Srt{}
	}

	var result []Srt
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

func WriteSrt(parsedSrts []Srt, path string) {
	output := ""

	for _, srt := range parsedSrts {
		output += fmt.Sprintf("%d\n%s\n%s\n\n", srt.count, srt.timestamp, srt.subtitle)
	}

	os.WriteFile(path, []byte(output), 0644)
}
