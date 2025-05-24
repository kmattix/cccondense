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

func TestValidateExtensionNone(t *testing.T) {
	data := ""
	actual := validateExtension(data)
	expected := false
	if actual != expected {
		t.Errorf("validateExtension(%s) = %v; want %v", data, actual, expected)
	}
}

func TestValidateExtensionTrue(t *testing.T) {
	data := "/path/to/my/.secret/cc.srt"
	actual := validateExtension(data)
	expected := true
	if actual != expected {
		t.Errorf("validateExtension(%s) = %v; want %v", data, actual, expected)
	}
}

func TestValidateExtensionFalse(t *testing.T) {
	data := "/path/to/my/.secret/audio.wav"
	actual := validateExtension(data)
	expected := false
	if actual != expected {
		t.Errorf("validateExtension(%s) = %v; want %v", data, actual, expected)
	}
}

func TestHandleFile(t *testing.T) {

}

func TestHandleFileDest(t *testing.T) {

}

func TestHandleFileWrite(t *testing.T) {

}
func TestHandleFileWriteDest(t *testing.T) {

}

func TestHandleDirectory(t *testing.T) {

}

func TestHandleDirectoryWrite(t *testing.T) {

}

func TestHandleDirectoryDest(t *testing.T) {

}

func TestHandleDirectoryWriteDest(t *testing.T) {

}
