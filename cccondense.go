package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Printf("Expeced 2 arguments, got %v.\n", len(args))
		os.Exit(1)
	}

	inputPath := args[0]
	outputPath := args[1]

	file, err := os.Open(inputPath)
	check(err)
	defer file.Close()

	parsedSrts := ParseSrt(file)
	condensedSrts := CondenseSrt(parsedSrts)

	WriteSrt(condensedSrts, outputPath)
	fmt.Printf("Condesed SRT written to: %s\n", outputPath)
}
