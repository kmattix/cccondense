package main

import (
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
