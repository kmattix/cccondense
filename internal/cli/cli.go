package cli

import (
	"errors"
	"strings"
)

type Flags struct {
	Recursive bool
	Verbose   bool
	Write     bool
}

const fileSuffix = "_condensed"

func Run(args []string, flags Flags) error {
	return errors.New("not implemented")
}

func validateExtension(path string) bool {
	extensions := [...]string{"srt"}
	tokens := strings.Split(path, ".")
	extension := tokens[len(tokens)-1]
	for _, e := range extensions {
		if extension == e {
			return true
		}
	}
	return false
}

func handleFile(input string, output string, write bool) error {
	//if output is empty this writes to current directory with fileSuffix
	return errors.New("not implemented")
}

func handleDirectory(input string, output string, write bool) error {
	return errors.New("not implemented")
}
