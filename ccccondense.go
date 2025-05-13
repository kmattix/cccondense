package main

import (
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatalf("Expeced 1 arguments, got %v.", len(args))
	}
}
