package main

import (
	"os"

	"github.com/kmattix/cccondense/internal/cli"
)

func main() {
	cli.Run(os.Args[1:])
}
