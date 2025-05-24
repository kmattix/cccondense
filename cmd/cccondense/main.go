package main

import (
	"flag"

	"github.com/kmattix/cccondense/internal/cli"
	"github.com/kmattix/cccondense/internal/utils"
)

func main() {
	recursive := flag.Bool("r", false, "Condense all closed-caption files in the directory and its subdirectories recursively.")
	verbose := flag.Bool("v", false, "Display progress information.")
	write := flag.Bool("w", false, "Overwrite the original file(s) with the condensed version.")

	flag.Parse()

	flags := cli.Flags{
		Recursive: *recursive,
		Verbose:   *verbose,
		Write:     *write,
	}

	args := flag.Args()

	err := cli.Run(args, flags)
	utils.Check(err)
}
