package main

import (
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatalf("Expeced 1 arguments, got %v.", len(args))
	}
	path := args[0]

	file, err := os.Open(path)
	check(err)
	defer file.Close()

}
