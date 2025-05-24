package cli

import (
	"errors"
)

type Flags struct {
	Recursive bool
	Verbose   bool
	Write     bool
}

func Run(args []string, flags Flags) error {
	return errors.New("not implemented")
}
