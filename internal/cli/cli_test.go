package cli

import "testing"

func TestRun(t *testing.T) {
	/*
		Useage: cccondense [OPTION]... FILE... [DEST]
				cccondense [OPTION]... DIRECTORY [DEST]
			Condenses closed-caption file(s). When DEST is not specified it will save a new file to the
			source DIRECTORY in the format of FILE-condensed.EXTENSTION

		OPTION
			-r, --recursive			condenses all closed-caption FILEs in a directory and their
									contents recursively
			-v, --verbose			displays the progress of what is being done
			-w, --write				overwites on top of the source FILE that is being condensed in place
	*/
}
