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

func TestParseSrtSingleLine(t *testing.T) {
	file, err := os.Open("test-data/test1.srt")
	check(err)
	defer file.Close()
	actual := ParseSrt(file)

	expected := []srt{
		{
			count:     1,
			timestamp: "00:00:03,330 --> 00:00:06,560",
			subtitle:  "Gentlemen, <i>Star Wars</i> Day",
		}, {
			count:     2,
			timestamp: "00:00:03,330 --> 00:00:06,560",
			subtitle:  "is rapidly approaching.",
		}, {
			count:     3,
			timestamp: "00:00:06,770 --> 00:00:08,360",
			subtitle:  "We should finalize our plans.",
		}, {
			count:     4,
			timestamp: "00:00:08,570 --> 00:00:10,540",
			subtitle:  "What? That's a real thing?",
		}, {
			count:     5,
			timestamp: "00:00:11,510 --> 00:00:13,670",
			subtitle:  "What is it, <i>Star Wars</i> Christmas?",
		},
	}

	if !slices.Equal(actual, expected) {
		t.Errorf("parseSrt(%v) = \n%v; want \n%v", file.Name(), actual, expected)
	}
}

func TestParseSrtMultiLine(t *testing.T) {
	file, err := os.Open("test-data/test2.srt")
	check(err)
	defer file.Close()
	actual := ParseSrt(file)

	expected := []srt{{
		count:     1,
		timestamp: "00:00:03,330 --> 00:00:06,560",
		subtitle:  "Gentlemen, <i>Star Wars</i> Day\nis rapidly\napproaching.",
	}, {
		count:     2,
		timestamp: "00:00:06,770 --> 00:00:08,360",
		subtitle:  "We should finalize our plans.",
	}, {
		count:     3,
		timestamp: "00:00:08,570 --> 00:00:10,540",
		subtitle:  "What? That's a real thing?",
	}, {
		count:     4,
		timestamp: "00:00:11,510 --> 00:00:13,670",
		subtitle:  "What is it, <i>Star Wars</i> Christmas?",
	},
	}

	if !slices.Equal(actual, expected) {
		t.Errorf("parseSrt(%v) = \n%v; want \n%v", file.Name(), actual, expected)
	}
}

func TestWriteSrt(t *testing.T) {
	file, err := os.Open("test-data/test.srt")
	check(err)
	defer file.Close()
	expected := ParseSrt(file)

	WriteSrt(expected, "test-data/writetest.srt")
	writtenFile, err := os.Open("test-data/writetest.srt")
	check(err)
	defer writtenFile.Close()
	actual := ParseSrt(writtenFile)

	if !slices.Equal(actual, expected) {
		t.Errorf("writeSrt(%v) = \n%v; want \n%v", file.Name(), actual, expected)
	}

	os.Remove(writtenFile.Name())
}

func TestCondenseSrt(t *testing.T) {
	data := []srt{
		{
			count:     1,
			timestamp: "00:00:03,330 --> 00:00:06,560",
			subtitle:  "Gentlemen, <i>Star Wars</i> Day",
		}, {
			count:     2,
			timestamp: "00:00:03,330 --> 00:00:06,560",
			subtitle:  "is rapidly",
		}, {
			count:     3,
			timestamp: "00:00:03,330 --> 00:00:06,560",
			subtitle:  "approaching.",
		}, {
			count:     4,
			timestamp: "00:00:06,770 --> 00:00:08,360",
			subtitle:  "We should finalize our plans.",
		}, {
			count:     5,
			timestamp: "00:00:08,570 --> 00:00:10,540",
			subtitle:  "What? That's a real thing?",
		},
		{
			count:     6,
			timestamp: "00:00:11,510 --> 00:00:13,670",
			subtitle:  "What is it, <i>Star Wars</i> Christmas?",
		},
	}

	actual := CondenseSrt(data)

	expected := []srt{
		{
			count:     1,
			timestamp: "00:00:03,330 --> 00:00:06,560",
			subtitle:  "Gentlemen, <i>Star Wars</i> Day\nis rapidly\napproaching.",
		}, {
			count:     2,
			timestamp: "00:00:06,770 --> 00:00:08,360",
			subtitle:  "We should finalize our plans.",
		}, {
			count:     3,
			timestamp: "00:00:08,570 --> 00:00:10,540",
			subtitle:  "What? That's a real thing?",
		}, {
			count:     4,
			timestamp: "00:00:11,510 --> 00:00:13,670",
			subtitle:  "What is it, <i>Star Wars</i> Christmas?",
		},
	}

	if !slices.Equal(actual, expected) {
		t.Errorf("condenseSrt(%v) = \n%v; want \n%v", data, actual, expected)
	}
}
