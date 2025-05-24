package cli

import "testing"

func TestRun(t *testing.T) {
	/*
		Usage:
			cccondense [OPTION]... FILE [DEST]
			cccondense [OPTION]... DIRECTORY [DEST]
		Condenses closed-caption file(s).
		When DEST is not specified, the program saves a new file in the source directory with the name
		format FILE-condensed.EXTENSION

		OPTION
			-r, --recursive		Condense all closed-caption files in the directory and its subdirectories.
			-v, --verbose		Display progress information.
			-w, --write			Overwrite the original file with the condensed version.
	*/
}
