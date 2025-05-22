package main

import (
	"os"
	"slices"
	"testing"
)

func TestNewSrt(t *testing.T) {
	data := `1
00:00:03,330 --> 00:00:06,560
Gentlemen, <i>Star Wars</i> Day

`
	actual := newSrt(data)

	expected := srt{
		count:     1,
		timestamp: "00:00:03,330 --> 00:00:06,560",
		subtitle:  "Gentlemen, <i>Star Wars</i> Day",
	}

	if *actual != expected {
		t.Errorf("newSrt(\n%s) = \n%v; want \n%v", data, actual, expected)
	}

}

func TestParseSrt(t *testing.T) {
	file, err := os.Open("test-data/test.srt")
	check(err)
	defer file.Close()
	actual := parseSrt(*file)

	expected := []srt{
		srt{
			count:     1,
			timestamp: "00:00:03,330 --> 00:00:06,560",
			subtitle:  "Gentlemen, <i>Star Wars</i> Day",
		},
		srt{
			count:     2,
			timestamp: "00:00:03,330 --> 00:00:06,560",
			subtitle:  "is rapidly approaching.",
		},
		srt{
			count:     3,
			timestamp: "00:00:06,770 --> 00:00:08,360",
			subtitle:  "We should finalize our plans.",
		},
		srt{
			count:     4,
			timestamp: "00:00:08,570 --> 00:00:10,540",
			subtitle:  "What? That's a real thing?",
		},
		srt{
			count:     5,
			timestamp: "00:00:11,510 --> 00:00:13,670",
			subtitle:  "What is it, <i>Star Wars</i> Christmas?",
		},
	}

	if !slices.Equal(actual, expected) {
		t.Errorf("parseSrt(%v) = \n%v; want \n%v", file.Name(), actual, expected)
	}
}
